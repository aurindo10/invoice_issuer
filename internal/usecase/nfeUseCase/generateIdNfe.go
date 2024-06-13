package nfeusecase

import nfeentitie "github.com/aurindo10/invoice_issuer/internal/entities/nfeEntitie"

type GenerateID struct {
	repository nfeentitie.IdNumberRepository
}

func (c *GenerateID) Execute(p *nfeentitie.NfeInfo) (*string, error) {
	cnF, error := c.repository.GetCnf()
	if error != nil {
		return nil, error
	}
	acessKey, error := c.repository.GetAcessKey(&p.LastNumber, cnF, p)
	if error != nil {
		return nil, error
	}
	cdv, error := c.repository.GetcDv(acessKey)
	if error != nil {
		return nil, error
	}

	fullAcessKey, error := c.repository.GetFullAcessKey(acessKey, cdv)
	if error != nil {
		return nil, error
	}
	return fullAcessKey, nil
}

func NewGenerateID(repository nfeentitie.IdNumberRepository) *GenerateID {
	return &GenerateID{
		repository: repository,
	}
}
