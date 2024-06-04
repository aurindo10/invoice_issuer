package tokenusecase

import (
	companyentitie "github.com/aurindo10/invoice_issuer/internal/entities/companyEntitie"
	tokenentitie "github.com/aurindo10/invoice_issuer/internal/entities/tokenEntitie"
)

type CreateTokenUseCase struct {
	repository tokenentitie.Repository
}

func (p *CreateTokenUseCase) Execute(c *companyentitie.Company) (*tokenentitie.Token, error) {
	res, error := p.repository.CreteToken(c)
	if error != nil {
		return nil, error
	}
	return res, nil
}
func NewCreateTokenUseCase(repository tokenentitie.Repository) *CreateTokenUseCase {
	return &CreateTokenUseCase{
		repository: repository,
	}
}
