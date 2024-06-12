package nfeusecase

import (
	nfeentitie "github.com/aurindo10/invoice_issuer/internal/entities/nfeEntitie"
	"github.com/google/uuid"
)

type ClientEntitie struct {
	repository nfeentitie.OpticaRepository
}

func (c *ClientEntitie) Execute(id uuid.UUID) (*nfeentitie.Dest, error) {
	info, error := c.repository.GetClientInfo(id)
	if error != nil {
		return nil, error
	}
	return info, nil
}

func NewClientUseCase(repository nfeentitie.OpticaRepository) *ClientEntitie {
	return &ClientEntitie{
		repository: repository,
	}
}
