package usecase

import (
	"github.com/aurindo10/invoice_issuer/internal/entities/certificate"
)

type RegisterCertificate struct {
	Integration certificate.Integration
	Repository  certificate.Repository
}

func (c *RegisterCertificate) Execute(p *certificate.Certificate) (*certificate.Certificate, error) {
	repo, error := c.Repository.RegisterCertificate(p)
	if error != nil {
		return nil, error
	}
	if error := c.Integration.RegisterCertificate(p); error != nil {
		return nil, error
	}
	return repo, nil
}

func NewRegisterCertificate(
	integration certificate.Integration,
	repository certificate.Repository,
) *RegisterCertificate {
	return &RegisterCertificate{
		Integration: integration,
		Repository:  repository,
	}
}
