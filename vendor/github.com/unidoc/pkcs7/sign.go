package pkcs7

import (
	"bytes"
	"crypto"
	"crypto/dsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"errors"
	"fmt"
	"math/big"
	"time"
)

// SignedData is an opaque data structure for creating signed data payloads
type SignedData struct {
	sd                  signedData
	certs               []*x509.Certificate
	data, messageDigest []byte
	digestOid           asn1.ObjectIdentifier
	encryptionOid       asn1.ObjectIdentifier
}

// NewSignedData takes data and initializes a PKCS7 SignedData struct that is
// ready to be signed via AddSigner. The digest algorithm is set to SHA1 by default
// and can be changed by calling SetDigestAlgorithm.
func NewSignedData(data []byte) (*SignedData, error) {
	signedData, err := newSignedData(OIDData, data)
	if err != nil {
		return nil, err
	}
	return signedData, nil
}

// One of the practical use cases of NewSignedDataWithContentType is:
// Pasted from RFC 3161
//
// 2.4.2. Response Format
// TimeStampToken ::= ContentInfo
//      -- contentType is id-signedData ([CMS])
//      -- content is SignedData ([CMS])
//
// The fields of type EncapsulatedContentInfo of the SignedData
// construct have the following meanings:
//
// eContentType is an object identifier that uniquely specifies the
// content type.  For a time-stamp token it is defined as:
//
// id-ct-TSTInfo  OBJECT IDENTIFIER ::= { iso(1) member-body(2)
// us(840) rsadsi(113549) pkcs(1) pkcs-9(9) smime(16) ct(1) 4}
//
// eContent is the content itself, carried as an octet string.
// The eContent SHALL be the DER-encoded value of TSTInfo.

// NewSignedDataWithContentType takes content type and data and initializes a PKCS7 SignedData struct that is
// ready to be signed via AddSigner. The digest algorithm is set to SHA1 by default
// and can be changed by calling SetDigestAlgorithm.
func NewSignedDataWithContentType(contentType asn1.ObjectIdentifier, data []byte) (*SignedData, error) {
	signedData, err := newSignedData(contentType, data)
	if err != nil {
		return nil, err
	}
	return signedData, nil
}

func newSignedData(contentType asn1.ObjectIdentifier, data []byte) (*SignedData, error) {
	content, err := asn1.Marshal(data)
	if err != nil {
		return nil, err
	}
	ci := contentInfo{
		ContentType: contentType,
		Content:     asn1.RawValue{Class: 2, Tag: 0, Bytes: content, IsCompound: true},
	}
	sd := signedData{
		ContentInfo: ci,
		Version:     1,
	}
	return &SignedData{sd: sd, data: data, digestOid: OIDDigestAlgorithmSHA1}, nil
}

// SignerInfoConfig are optional values to include when adding a signer
type SignerInfoConfig struct {
	ExtraSignedAttributes   []Attribute
	ExtraUnsignedAttributes []Attribute
}

type signedData struct {
	Version                    int                        `asn1:"default:1"`
	DigestAlgorithmIdentifiers []pkix.AlgorithmIdentifier `asn1:"set"`
	ContentInfo                contentInfo
	Certificates               rawCertificates        `asn1:"optional,tag:0"`
	CRLs                       []pkix.CertificateList `asn1:"optional,tag:1"`
	SignerInfos                []signerInfo           `asn1:"set"`
}

type signerInfo struct {
	Version                   int `asn1:"default:1"`
	IssuerAndSerialNumber     issuerAndSerial
	DigestAlgorithm           pkix.AlgorithmIdentifier
	AuthenticatedAttributes   []attribute `asn1:"optional,tag:0"`
	DigestEncryptionAlgorithm pkix.AlgorithmIdentifier
	EncryptedDigest           []byte
	UnauthenticatedAttributes []attribute `asn1:"optional,tag:1"`
}

type attribute struct {
	Type  asn1.ObjectIdentifier
	Value asn1.RawValue `asn1:"set"`
}

func marshalAttributes(attrs []attribute) ([]byte, error) {
	encodedAttributes, err := asn1.MarshalWithParams(attrs, "set")
	if err != nil {
		return nil, err
	}

	return encodedAttributes, nil
}

type rawCertificates struct {
	Raw asn1.RawContent
}

type issuerAndSerial struct {
	IssuerName   asn1.RawValue
	SerialNumber *big.Int
}

// SetDigestAlgorithm sets the digest algorithm to be used in the signing process.
//
// This should be called before adding signers
func (sd *SignedData) SetDigestAlgorithm(d asn1.ObjectIdentifier) {
	sd.digestOid = d
}

// SetEncryptionAlgorithm sets the encryption algorithm to be used in the signing process.
//
// This should be called before adding signers
func (sd *SignedData) SetEncryptionAlgorithm(d asn1.ObjectIdentifier) {
	sd.encryptionOid = d
}

// AddSigner is a wrapper around AddSignerChain() that adds a signer without any parent.
func (sd *SignedData) AddSigner(ee *x509.Certificate, pkey crypto.PrivateKey, config SignerInfoConfig) error {
	var parents []*x509.Certificate
	return sd.addSignerChain(ee, pkey, parents, config, true, true)
}

// One of the practical use cases of AddSignerNoChain is:
// Pasted from RFC 3161

// 2.4.1. Request Format
// If the certReq field is present and set to true, the TSA's public key
// certificate that is referenced by the ESSCertID identifier inside a
// SigningCertificate attribute in the response MUST be provided by the
// TSA in the certificates field from the SignedData structure in that
// response.  That field may also contain other certificates.

// If the certReq field is missing or if the certReq field is present
// and set to false then the certificates field from the SignedData
// structure MUST not be present in the response.

// AddSignerNoChain is a wrapper around AddSignerChain() that adds a signer without any parent.
// Use this method, if no certificate needs to be placed in SignedData certificates
func (sd *SignedData) AddSignerNoChain(ee *x509.Certificate, pkey crypto.PrivateKey, config SignerInfoConfig) error {
	var parents []*x509.Certificate
	return sd.addSignerChain(ee, pkey, parents, config, false, true)
}

// AddSignerChain signs attributes about the content and adds certificates
// and signers infos to the Signed Data. The certificate and private
// of the end-entity signer are used to issue the signature, and any
// parent of that end-entity that need to be added to the list of
// certifications can be specified in the parents slice.
//
// The signature algorithm used to hash the data is the one of the end-entity
// certificate.
//
// Following RFC 2315, 9.2 SignerInfo type, the distinguished name of
// the issuer of the end-entity signer is stored in the issuerAndSerialNumber
// section of the SignedData.SignerInfo, alongside the serial number of
// the end-entity.
func (sd *SignedData) AddSignerChain(ee *x509.Certificate, pkey crypto.PrivateKey, chain []*x509.Certificate, config SignerInfoConfig) error {
	return sd.addSignerChain(ee, pkey, chain, config, true, true)
}

// AddSignerChainPAdES signs attributes about the content and adds certificates
// and signers infos to the Signed Data. The certificate and private
// of the end-entity signer are used to issue the signature, and any
// parent of that end-entity that need to be added to the list of
// certifications can be specified in the parents slice.
//
// It is compatible with PAdES specifications.
//
// The signature algorithm used to hash the data is the one of the end-entity
// certificate.
//
// Following RFC 2315, 9.2 SignerInfo type, the distinguished name of
// the issuer of the end-entity signer is stored in the issuerAndSerialNumber
// section of the SignedData.SignerInfo, alongside the serial number of
// the end-entity.
func (sd *SignedData) AddSignerChainPAdES(ee *x509.Certificate, pkey crypto.PrivateKey, chain []*x509.Certificate, config SignerInfoConfig) error {
	return sd.addSignerChain(ee, pkey, chain, config, true, false)
}

func (sd *SignedData) addSignerChain(ee *x509.Certificate, pkey crypto.PrivateKey, chain []*x509.Certificate, config SignerInfoConfig, includeCertificates bool, enableSigningTime bool) error {
	sd.sd.DigestAlgorithmIdentifiers = append(sd.sd.DigestAlgorithmIdentifiers, pkix.AlgorithmIdentifier{Algorithm: sd.digestOid})
	hash, err := getHashForOID(sd.digestOid)
	if err != nil {
		return err
	}
	h := hash.New()
	h.Write(sd.data)
	sd.messageDigest = h.Sum(nil)
	encryptionOid, err := getOIDForEncryptionAlgorithm(pkey, sd.digestOid)
	if err != nil {
		return err
	}
	attrs := &attributes{}
	attrs.Add(OIDAttributeContentType, sd.sd.ContentInfo.ContentType)
	attrs.Add(OIDAttributeMessageDigest, sd.messageDigest)
	if enableSigningTime {
		attrs.Add(OIDAttributeSigningTime, time.Now())
	}

	// Add id-aa-signing-certificate-v2.
	if b, err := populateSigningCertificateV2Ext(ee); err == nil {
		attrs.Add(OIDAttributeSigningCertificateV2, asn1.RawValue{FullBytes: b})
	}

	for _, attr := range config.ExtraSignedAttributes {
		attrs.Add(attr.Type, attr.Value)
	}
	finalAttrs, err := attrs.ForMarshalling()
	if err != nil {
		return err
	}
	unsignedAttrs := &attributes{}
	for _, attr := range config.ExtraUnsignedAttributes {
		unsignedAttrs.Add(attr.Type, attr.Value)
	}
	finalUnsignedAttrs, err := unsignedAttrs.ForMarshalling()
	if err != nil {
		return err
	}
	signature, err := signAttributes(finalAttrs, pkey, hash)
	if err != nil {
		return err
	}
	var ias issuerAndSerial
	// No parent, the issue is the end-entity cert itself
	ias.IssuerName = asn1.RawValue{FullBytes: ee.RawIssuer}
	ias.SerialNumber = ee.SerialNumber
	if len(chain) == 0 {
		// no parent, the issue is the end-entity cert itself
		ias.IssuerName = asn1.RawValue{FullBytes: ee.RawIssuer}
	} else {
		err = verifyPartialChain(ee, chain)
		if err != nil {
			return err
		}
		// the first parent is the issuer
		ias.IssuerName = asn1.RawValue{FullBytes: chain[0].RawSubject}
	}
	signer := signerInfo{
		AuthenticatedAttributes:   finalAttrs,
		UnauthenticatedAttributes: finalUnsignedAttrs,
		DigestAlgorithm:           pkix.AlgorithmIdentifier{Algorithm: sd.digestOid},
		DigestEncryptionAlgorithm: pkix.AlgorithmIdentifier{Algorithm: encryptionOid},
		IssuerAndSerialNumber:     ias,
		EncryptedDigest:           signature,
		Version:                   1,
	}
	// create signature of signed attributes
	sd.sd.SignerInfos = append(sd.sd.SignerInfos, signer)

	if includeCertificates {
		sd.certs = append(sd.certs, ee)
		sd.certs = append(sd.certs, chain...)
	}

	if len(chain) > 0 {
		sd.certs = append(sd.certs, chain...)
	}

	return nil
}

// SignWithoutAttr issues a signature on the content of the pkcs7 SignedData.
// Unlike AddSigner/AddSignerChain, it calculates the digest on the data alone
// and does not include any signed attributes like timestamp and so on.
//
// This function is needed to sign old Android APKs, something you probably
// shouldn't do unless you're maintaining backward compatibility for old
// applications.
func (sd *SignedData) SignWithoutAttr(ee *x509.Certificate, pkey crypto.PrivateKey, config SignerInfoConfig) error {
	var signature []byte
	sd.sd.DigestAlgorithmIdentifiers = append(sd.sd.DigestAlgorithmIdentifiers, pkix.AlgorithmIdentifier{Algorithm: sd.digestOid})
	hash, err := getHashForOID(sd.digestOid)
	if err != nil {
		return err
	}
	h := hash.New()
	h.Write(sd.data)
	sd.messageDigest = h.Sum(nil)
	switch pkey.(type) {
	case *dsa.PrivateKey:
		// dsa doesn't implement crypto.Signer so we make a special case
		// https://github.com/golang/go/issues/27889
		r, s, err := dsa.Sign(rand.Reader, pkey.(*dsa.PrivateKey), sd.messageDigest)
		if err != nil {
			return err
		}
		signature, err = asn1.Marshal(dsaSignature{r, s})
		if err != nil {
			return err
		}
	default:
		key, ok := pkey.(crypto.Signer)
		if !ok {
			return errors.New("pkcs7: private key does not implement crypto.Signer")
		}
		signature, err = key.Sign(rand.Reader, sd.messageDigest, hash)
		if err != nil {
			return err
		}
	}
	var ias issuerAndSerial
	ias.SerialNumber = ee.SerialNumber
	// no parent, the issue is the end-entity cert itself
	ias.IssuerName = asn1.RawValue{FullBytes: ee.RawIssuer}
	if sd.encryptionOid == nil {
		// if the encryption algorithm wasn't set by SetEncryptionAlgorithm,
		// infer it from the digest algorithm
		sd.encryptionOid, err = getOIDForEncryptionAlgorithm(pkey, sd.digestOid)
	}
	if err != nil {
		return err
	}
	signer := signerInfo{
		DigestAlgorithm:           pkix.AlgorithmIdentifier{Algorithm: sd.digestOid},
		DigestEncryptionAlgorithm: pkix.AlgorithmIdentifier{Algorithm: sd.encryptionOid},
		IssuerAndSerialNumber:     ias,
		EncryptedDigest:           signature,
		Version:                   1,
	}
	// create signature of signed attributes
	sd.certs = append(sd.certs, ee)
	sd.sd.SignerInfos = append(sd.sd.SignerInfos, signer)
	return nil
}

func (si *signerInfo) SetUnauthenticatedAttributes(extraUnsignedAttrs []Attribute) error {
	unsignedAttrs := &attributes{}
	for _, attr := range extraUnsignedAttrs {
		unsignedAttrs.Add(attr.Type, attr.Value)
	}
	finalUnsignedAttrs, err := unsignedAttrs.ForMarshalling()
	if err != nil {
		return err
	}

	si.UnauthenticatedAttributes = finalUnsignedAttrs

	return nil
}

// TimestampTokenRequestCallback callback of timestamp token request.
type TimestampTokenRequestCallback func(digest []byte) ([]byte, error)

// RequestSignerTimestampToken add request of timestamp token with `signerID`
// the request of timestamp token is called within `callback` function.
func (sd *SignedData) RequestSignerTimestampToken(signerID int, callback TimestampTokenRequestCallback) error {
	if len(sd.sd.SignerInfos) < (signerID + 1) {
		return fmt.Errorf("no signer information found for ID %d", signerID)
	}

	if callback == nil {
		return fmt.Errorf("no callback defined")
	}

	tst, err := callback(sd.sd.SignerInfos[signerID].EncryptedDigest)
	if err != nil {
		return err
	}
	return sd.AddTimestampTokenToSigner(signerID, tst)
}

// AddTimestampTokenToSigner inserts `tst` TimestampToken which described in RFC3161 into
// unauthenticated attribute of `signerID` which obtaioned from identity service.
func (sd *SignedData) AddTimestampTokenToSigner(signerID int, tst []byte) (err error) {
	if len(sd.sd.SignerInfos) < (signerID + 1) {
		return fmt.Errorf("no signer information found for ID %d", signerID)
	}

	// Add the timestamp token to the unauthenticated attributes.
	attrs := &attributes{}
	for _, attr := range sd.sd.SignerInfos[signerID].UnauthenticatedAttributes {
		attrs.Add(attr.Type, attr.Value)
	}

	attrs.Add(OIDAttributeTimeStampToken, asn1.RawValue{FullBytes: tst})
	sd.sd.SignerInfos[signerID].UnauthenticatedAttributes, err = attrs.ForMarshalling()
	if err != nil {
		return err
	}
	return nil
}

// AddCertificate adds the certificate to the payload. Useful for parent certificates
func (sd *SignedData) AddCertificate(cert *x509.Certificate) {
	sd.certs = append(sd.certs, cert)
}

// Detach removes content from the signed data struct to make it a detached signature.
// This must be called right before Finish()
func (sd *SignedData) Detach() {
	sd.sd.ContentInfo = contentInfo{ContentType: OIDData}
}

// GetSignedData returns the private Signed Data
func (sd *SignedData) GetSignedData() *signedData {
	return &sd.sd
}

// Finish marshals the content and its signers
func (sd *SignedData) Finish() ([]byte, error) {
	sd.sd.Certificates = marshalCertificates(sd.certs)
	inner, err := asn1.Marshal(sd.sd)
	if err != nil {
		return nil, err
	}
	outer := contentInfo{
		ContentType: OIDSignedData,
		Content:     asn1.RawValue{Class: 2, Tag: 0, Bytes: inner, IsCompound: true},
	}
	return asn1.Marshal(outer)
}

// RemoveAuthenticatedAttributes removes authenticated attributes from signedData
// similar to OpenSSL's PKCS7_NOATTR or -noattr flags
func (sd *SignedData) RemoveAuthenticatedAttributes() {
	for i := range sd.sd.SignerInfos {
		sd.sd.SignerInfos[i].AuthenticatedAttributes = nil
	}
}

// RemoveUnauthenticatedAttributes removes unauthenticated attributes from signedData
func (sd *SignedData) RemoveUnauthenticatedAttributes() {
	for i := range sd.sd.SignerInfos {
		sd.sd.SignerInfos[i].UnauthenticatedAttributes = nil
	}
}

// verifyPartialChain checks that a given cert is issued by the first parent in the list,
// then continue down the path. It doesn't require the last parent to be a root CA,
// or to be trusted in any truststore. It simply verifies that the chain provided, albeit
// partial, makes sense.
func verifyPartialChain(cert *x509.Certificate, parents []*x509.Certificate) error {
	if len(parents) == 0 {
		return fmt.Errorf("pkcs7: zero parents provided to verify the signature of certificate %q", cert.Subject.CommonName)
	}
	err := cert.CheckSignatureFrom(parents[0])
	if err != nil {
		return fmt.Errorf("pkcs7: certificate signature from parent is invalid: %v", err)
	}
	if len(parents) == 1 {
		// there is no more parent to check, return
		return nil
	}
	return verifyPartialChain(parents[0], parents[1:])
}

func cert2issuerAndSerial(cert *x509.Certificate) (issuerAndSerial, error) {
	var ias issuerAndSerial
	// The issuer RDNSequence has to match exactly the sequence in the certificate
	// We cannot use cert.Issuer.ToRDNSequence() here since it mangles the sequence
	ias.IssuerName = asn1.RawValue{FullBytes: cert.RawIssuer}
	ias.SerialNumber = cert.SerialNumber

	return ias, nil
}

// signs the DER encoded form of the attributes with the private key
func signAttributes(attrs []attribute, pkey crypto.PrivateKey, digestAlg crypto.Hash) ([]byte, error) {
	attrBytes, err := marshalAttributes(attrs)
	if err != nil {
		return nil, err
	}
	h := digestAlg.New()
	h.Write(attrBytes)
	hash := h.Sum(nil)

	// dsa doesn't implement crypto.Signer so we make a special case
	// https://github.com/golang/go/issues/27889
	switch pkey.(type) {
	case *dsa.PrivateKey:
		r, s, err := dsa.Sign(rand.Reader, pkey.(*dsa.PrivateKey), hash)
		if err != nil {
			return nil, err
		}
		return asn1.Marshal(dsaSignature{r, s})
	}

	key, ok := pkey.(crypto.Signer)
	if !ok {
		return nil, errors.New("pkcs7: private key does not implement crypto.Signer")
	}
	return key.Sign(rand.Reader, hash, digestAlg)
}

type dsaSignature struct {
	R, S *big.Int
}

// concats and wraps the certificates in the RawValue structure
func marshalCertificates(certs []*x509.Certificate) rawCertificates {
	var buf bytes.Buffer
	for _, cert := range certs {
		buf.Write(cert.Raw)
	}
	rawCerts, _ := marshalCertificateBytes(buf.Bytes())
	return rawCerts
}

// Even though, the tag & length are stripped out during marshalling the
// RawContent, we have to encode it into the RawContent. If its missing,
// then `asn1.Marshal()` will strip out the certificate wrapper instead.
func marshalCertificateBytes(certs []byte) (rawCertificates, error) {
	var val = asn1.RawValue{Bytes: certs, Class: 2, Tag: 0, IsCompound: true}
	b, err := asn1.Marshal(val)
	if err != nil {
		return rawCertificates{}, err
	}
	return rawCertificates{Raw: b}, nil
}

// DegenerateCertificate creates a signed data structure containing only the
// provided certificate or certificate chain.
func DegenerateCertificate(cert []byte) ([]byte, error) {
	rawCert, err := marshalCertificateBytes(cert)
	if err != nil {
		return nil, err
	}
	emptyContent := contentInfo{ContentType: OIDData}
	sd := signedData{
		Version:      1,
		ContentInfo:  emptyContent,
		Certificates: rawCert,
		CRLs:         []pkix.CertificateList{},
	}
	content, err := asn1.Marshal(sd)
	if err != nil {
		return nil, err
	}
	signedContent := contentInfo{
		ContentType: OIDSignedData,
		Content:     asn1.RawValue{Class: 2, Tag: 0, Bytes: content, IsCompound: true},
	}
	return asn1.Marshal(signedContent)
}

func populateSigningCertificateV2Ext(certificate *x509.Certificate) ([]byte, error) {
	h := sha256.New()
	h.Write(certificate.Raw)

	signingCertificateV2 := signingCertificateV2{
		Certs: []essCertIDv2{
			{
				HashAlgorithm: pkix.AlgorithmIdentifier{
					Algorithm:  asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 2, 1},
					Parameters: asn1.NullRawValue,
				},
				CertHash: h.Sum(nil),
				IssuerSerial: issuerAndSerial{
					IssuerName:   asn1.RawValue{FullBytes: certificate.RawIssuer},
					SerialNumber: certificate.SerialNumber,
				},
			},
		},
	}
	signingCertV2Bytes, err := asn1.Marshal(signingCertificateV2)
	if err != nil {
		return nil, err
	}
	return signingCertV2Bytes, nil
}
