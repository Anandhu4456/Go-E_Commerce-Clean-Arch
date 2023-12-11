package pkcs7

import (
	"crypto/x509/pkix"
	"encoding/asn1"
)

type generalNames struct {
	Name asn1.RawValue `asn1:"optional,tag:4"`
}

type essCertIDv2 struct {
	HashAlgorithm pkix.AlgorithmIdentifier `asn1:"optional"` // default sha256
	CertHash      []byte
	IssuerSerial  issuerAndSerial `asn1:"optional"`
}

type signingCertificateV2 struct {
	Certs []essCertIDv2
}
