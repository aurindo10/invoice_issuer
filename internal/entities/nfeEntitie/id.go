package nfeentitie

import (
	"math/rand"
	"strconv"
)

type NfeInfo struct {
	Cuf    string //state
	Aamm   string //year and month
	Cnpj   string
	Mod    string //55 by default
	Serie  string // 1 or 2
	NNf    string // numero seqencial da nota
	TpEmis string // 1 by default
	// cNF    string // randon number
	// cdv    string // calculated number
}

type GenerateID struct {
	repository IdNumberRepository
}

type IdNumberRepository interface {
	GetLastNumbernNF(cnpj *string) (*int64, error)
	GetCnf() (*string, error)
	GetAcessKey(lastNumber *int64, cnf *string, info *NfeInfo) (*string, error)
	GetcDv(acessKey *string) (*string, error)
	GetFullAcessKey(acessKey *string, Dv *string) (*string, error)
}

func (c *GenerateID) Execute(p *NfeInfo) (*string, error) {
	lastNumber, error := c.repository.GetLastNumbernNF(&p.Cnpj)
	if error != nil {
		return nil, error
	}
	cnF, error := c.repository.GetCnf()
	if error != nil {
		return nil, error
	}
	acessKey, error := c.repository.GetAcessKey(lastNumber, cnF, p)
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

func NewGenerateID(repository IdNumberRepository) *GenerateID {
	return &GenerateID{
		repository: repository,
	}
}

type IdRepository struct {
}

func (c *IdRepository) GetCnf() (*string, error) {
	max := int64(100000000) // Limite superior para 8 dígitos
	n := rand.Int63n(max)
	cnf := strconv.FormatInt(n, 10)
	return &cnf, nil
}
func (c *IdRepository) GetLastNumbernNF(cnpj *string) (*int64, error)
