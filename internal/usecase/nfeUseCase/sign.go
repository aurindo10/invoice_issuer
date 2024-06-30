package nfeusecase

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
	"strings"

	nfeentitie "github.com/aurindo10/invoice_issuer/internal/entities/nfeEntitie"
	"github.com/beevik/etree"
)

func runOpenSSLCommand(args ...string) (string, error) {
	cmd := exec.Command("openssl", args...)
	var out strings.Builder
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run OpenSSL command: %v, output: %s", err, out.String())
	}
	return out.String(), nil
}

// Extrai o certificado e a chave privada do arquivo PFX usando OpenSSL
func extractPEMFromPFX(pfxPath, password, certOutPath, keyOutPath string) error {
	// Extrai o certificado
	_, err := runOpenSSLCommand("pkcs12", "-legacy", "-in", pfxPath, "-out", certOutPath, "-clcerts", "-nokeys", "-passin", "pass:"+password, "-nodes")
	if err != nil {
		return fmt.Errorf("failed to extract certificate: %v", err)
	}

	// Extrai a chave privada
	_, err = runOpenSSLCommand("pkcs12", "-legacy", "-in", pfxPath, "-out", keyOutPath, "-nocerts", "-passin", "pass:"+password, "-nodes")
	if err != nil {
		return fmt.Errorf("failed to extract private key: %v", err)
	}

	return nil
}

func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	pemData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block: no block found")
	}
	if block.Type == "RSA PRIVATE KEY" {
		privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse PKCS1 private key: %v", err)
		}
		return privateKey, nil
	} else if block.Type == "PRIVATE KEY" {
		privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse PKCS8 private key: %v", err)
		}
		if rsaKey, ok := privateKey.(*rsa.PrivateKey); ok {
			return rsaKey, nil
		} else {
			return nil, fmt.Errorf("not an RSA private key")
		}
	}

	return nil, fmt.Errorf("failed to decode PEM block: expected RSA PRIVATE KEY or PRIVATE KEY, got %s", block.Type)
}

func signXML(privateKey *rsa.PrivateKey, data []byte) (string, error) {
	h := sha1.New()
	h.Write(data)
	hashed := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA1, hashed)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

func canonicalizeXML(input []byte) ([]byte, error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(input); err != nil {
		return nil, err
	}
	doc.WriteSettings = etree.WriteSettings{
		CanonicalEndTags: true,
		CanonicalText:    true,
		CanonicalAttrVal: true,
	}
	return doc.WriteToBytes()
}

func SignXML(pfxPath string, password string, xmlContent []byte, id string) ([]byte, *nfeentitie.Signature, error) {
	certOutPath := "client.pem"
	keyOutPath := "key.pem"
	err := extractPEMFromPFX(pfxPath, password, certOutPath, keyOutPath)
	if err != nil {
		return nil, nil, err
	}

	privateKey, err := LoadPrivateKey(keyOutPath)
	if err != nil {
		return nil, nil, err
	}

	certPEM, err := os.ReadFile(certOutPath)
	if err != nil {
		return nil, nil, err
	}

	block, _ := pem.Decode(certPEM)
	if block == nil || block.Type != "CERTIFICATE" {
		return nil, nil, fmt.Errorf("failed to decode PEM block containing certificate")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, nil, err
	}

	// Canonicalizar o XML
	canonicalXML, err := canonicalizeXML(xmlContent)
	if err != nil {
		return nil, nil, err
	}

	// Calcular o valor do digest
	digest := sha1.New()
	digest.Write(canonicalXML)
	digestValue := base64.StdEncoding.EncodeToString(digest.Sum(nil))

	// Assinar o XML canonicalizado
	signatureValue, err := signXML(privateKey, canonicalXML)
	if err != nil {
		return nil, nil, err
	}

	certBase64 := base64.StdEncoding.EncodeToString(cert.Raw)
	sig := nfeentitie.Signature{
		XMLName: xml.Name{Local: "Signature"},
		SignedInfo: nfeentitie.SignedInfo{
			CanonicalizationMethod: nfeentitie.CanonicalizationMethod{
				Algorithm: "http://www.w3.org/TR/2001/REC-xml-c14n-20010315",
			},
			SignatureMethod: nfeentitie.SignatureMethod{
				Algorithm: "http://www.w3.org/2000/09/xmldsig#rsa-sha1",
			},
			Reference: nfeentitie.Reference{
				URI: "#" + id,
				Transforms: nfeentitie.Transforms{
					Transform: []nfeentitie.Transform{
						{Algorithm: "http://www.w3.org/TR/2001/REC-xml-c14n-20010315"},
						{Algorithm: "http://www.w3.org/2000/09/xmldsig#enveloped-signature"},
					},
				},
				DigestMethod: nfeentitie.DigestMethod{
					Algorithm: "http://www.w3.org/2000/09/xmldsig#sha1",
				},
				DigestValue: digestValue,
			},
		},
		SignatureValue: signatureValue,
		KeyInfo: nfeentitie.KeyInfo{
			X509Data: nfeentitie.X509Data{
				X509Certificate: certBase64,
			},
		},
	}

	// Serializar a assinatura para XML
	signatureXML, err := xml.Marshal(sig)
	if err != nil {
		return nil, nil, err
	}

	// Anexar a assinatura ao XML original
	finalXML := append(xmlContent, signatureXML...)

	return finalXML, &sig, nil
}
