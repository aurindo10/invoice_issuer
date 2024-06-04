package tokenservice

import (
	companyentitie "github.com/aurindo10/invoice_issuer/internal/entities/CompanyEntitie"
	tokenentitie "github.com/aurindo10/invoice_issuer/internal/entities/tokenEntitie"
	tokenrepository "github.com/aurindo10/invoice_issuer/internal/repositories/tokenRepository"
	tokenusecase "github.com/aurindo10/invoice_issuer/internal/usecase/tokenUsecase"
)

type CreateTokenService struct {
	company companyentitie.Company
}

func (c *CreateTokenService) CreateTokenService() (*tokenentitie.Token, error) {

	repo := tokenrepository.NewTokenRepository()
	token := tokenusecase.NewCreateTokenUseCase(repo)
	res, error := token.Execute(&c.company)
	if error != nil {
		return nil, error
	}
	return res, nil
}

func NewCreateTokenService(company companyentitie.Company) *CreateTokenService {
	return &CreateTokenService{
		company: company,
	}
}
