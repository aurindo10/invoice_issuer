package integration

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/aurindo10/invoice_issuer/internal/entities/certificate"
	"github.com/aurindo10/invoice_issuer/pkg/utils"
)

type CertificateRepository struct{}

func (c *CertificateRepository) RegisterCertificate(p *certificate.Certificate) (*certificate.Certificate, error) {
	acessToken := utils.GetEnv("PLUGNOTASTOKEN", "NONE")
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, err := writer.CreateFormFile("arquivo", p.Cnpj)
	if err != nil {
		return nil, err
	}
	part.Write(p.Certificate)

	err = writer.WriteField("senha", p.Password)
	if err != nil {
		return nil, err
	}

	err = writer.WriteField("email", p.Email)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.sandbox.plugnotas.com.br/", &buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("x-api-key", acessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to register certificate: %s", resp.Status)
	}
	return p, nil
}

func NewCertificateIntegration() *CertificateRepository {
	return &CertificateRepository{}
}
