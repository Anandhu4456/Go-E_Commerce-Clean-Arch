package pkcs7

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

var encodeIndent = 0

type asn1Object interface {
	EncodeTo(writer *bytes.Buffer) error
	TagBytes() []byte
}

type asn1Structured struct {
	tagBytes []byte
	content  []asn1Object
}

func (s asn1Structured) TagBytes() []byte {
	return s.tagBytes
}

func (s asn1Structured) EncodeTo(out *bytes.Buffer) error {
	//fmt.Printf("%s--> tag: % X\n", strings.Repeat("| ", encodeIndent), s.tagBytes)
	encodeIndent++
	inner := new(bytes.Buffer)
	for _, obj := range s.content {
		err := obj.EncodeTo(inner)
		if err != nil {
			return err
		}
	}
	encodeIndent--
	out.Write(s.tagBytes)
	encodeLength(out, inner.Len())
	out.Write(inner.Bytes())
	return nil
}

type asn1Primitive struct {
	tagBytes []byte
	length   int
	content  []byte
}

func (s asn1Primitive) TagBytes() []byte {
	return s.tagBytes
}

func (p asn1Primitive) EncodeTo(out *bytes.Buffer) error {
	_, err := out.Write(p.tagBytes)
	if err != nil {
		return err
	}
	if err = encodeLength(out, p.length); err != nil {
		return err
	}
	//fmt.Printf("%s--> tag: % X length: %d\n", strings.Repeat("| ", encodeIndent), p.tagBytes, p.length)
	//fmt.Printf("%s--> content length: %d\n", strings.Repeat("| ", encodeIndent), len(p.content))
	out.Write(p.content)

	return nil
}

func ber2der(data []byte) ([]byte, error) {
	out := new(bytes.Buffer)

	obj, err := readObject(bytes.NewReader(data))
	if err != nil && err != io.EOF {
		return nil, err
	}
	if obj == nil {
		return nil, fmt.Errorf("error to parse BER")
	}
	obj.EncodeTo(out)

	return out.Bytes(), nil
}

// encodes lengths that are longer than 127 into string of bytes
func marshalLongLength(out *bytes.Buffer, i int) (err error) {
	n := lengthLength(i)

	for ; n > 0; n-- {
		err = out.WriteByte(byte(i >> uint((n-1)*8)))
		if err != nil {
			return
		}
	}

	return nil
}

// computes the byte length of an encoded length value
func lengthLength(i int) (numBytes int) {
	numBytes = 1
	for i > 255 {
		numBytes++
		i >>= 8
	}
	return
}

// encodes the length in DER format
// If the length fits in 7 bits, the value is encoded directly.
//
// Otherwise, the number of bytes to encode the length is first determined.
// This number is likely to be 4 or less for a 32bit length. This number is
// added to 0x80. The length is encoded in big endian encoding follow after
//
// Examples:
//  length | byte 1 | bytes n
//  0      | 0x00   | -
//  120    | 0x78   | -
//  200    | 0x81   | 0xC8
//  500    | 0x82   | 0x01 0xF4
//
func encodeLength(out *bytes.Buffer, length int) (err error) {
	if length >= 128 {
		l := lengthLength(length)
		err = out.WriteByte(0x80 | byte(l))
		if err != nil {
			return
		}
		err = marshalLongLength(out, length)
		if err != nil {
			return
		}
	} else {
		err = out.WriteByte(byte(length))
		if err != nil {
			return
		}
	}
	return
}

func readObject(r *bytes.Reader) (obj asn1Object, err error) {
	var tagB byte
	if tagB, err = r.ReadByte(); err != nil {
		return
	}

	primitive := tagB&0x20 == 0

	var l byte
	if l, err = r.ReadByte(); err != nil {
		return nil, fmt.Errorf("end of ber data reached")
	}
	length := (int)(l & 0x7F)
	if l > 0x80 {
		numberOfBytes := length
		length = 0
		if numberOfBytes > 4 { // int is only guaranteed to be 32bit
			return nil, errors.New("ber2der: BER tag length too long")
		}
		for i := 0; i < numberOfBytes; i++ {
			var sl byte
			if sl, err = r.ReadByte(); err != nil {
				return nil, fmt.Errorf("length is more than available data")
			}
			if i == 0 {
				if numberOfBytes == 4 && (int)(sl) > 0x7F {
					return nil, errors.New("ber2der: BER tag length is negative")
				}
				if 0x0 == (int)(sl) {
					return nil, errors.New("ber2der: BER tag length has leading zero")
				}
			}
			length = length*256 + (int)(sl)
		}
	}

	if primitive {
		p := asn1Primitive{
			tagBytes: []byte{tagB},
			length:   length,
			content:  make([]byte, length),
		}
		if length > 0 {
			_, err = r.Read(p.content)
		}
		if err != nil {
			return nil, fmt.Errorf("Invalid BER format")
		}
		obj = p
		return
	} else {
		subObjects := make([]asn1Object, 0)
		var sobj asn1Object
		if length > 0 {
			content := make([]byte, length)
			n, err := r.Read(content)
			if err != nil {
				return nil, err
			}
			if n != length {
				return nil, fmt.Errorf("length is more than available data")
			}
			r = bytes.NewReader(content)
		}
		for {
			sobj, err = readObject(r)
			if err != nil && err != io.EOF {
				return
			}
			if err == io.EOF && length >= 0 {
				err = nil
				break
			}
			if sobj == nil {
				break
			}
			if po, ok := sobj.(asn1Primitive); ok {
				if po.length == 0 && po.tagBytes[0] == 0 && length == 0 {
					break
				}
			}
			subObjects = append(subObjects, sobj)

			if err != nil {
				break
			}
		}
		obj = asn1Structured{
			tagBytes: []byte{tagB},
			content:  subObjects,
		}
		return
	}
	return nil, io.EOF
}
