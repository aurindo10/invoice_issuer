package nfeusecase

import (
	nfeentitie "github.com/aurindo10/invoice_issuer/internal/entities/nfeEntitie"
	"github.com/google/uuid"
)

type CompanyUseCase struct {
	repository nfeentitie.OpticaRepository
}

func (c *CompanyUseCase) Execute(id uuid.UUID) (*nfeentitie.Emit, error) {
	info, error := c.repository.GetCompanyInfo(id)
	if error != nil {
		return nil, error
	}
	return info, nil
}

func NewCompanyUseCase(repository nfeentitie.OpticaRepository) *CompanyUseCase {
	return &CompanyUseCase{
		repository: repository,
	}
}
