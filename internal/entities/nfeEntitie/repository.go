package nfeentitie

import (
	"github.com/google/uuid"
)

type IdNumberRepository interface {
	GetLastNumbernNF(cnpj *string) (*int64, error)
	GetCnf() (*string, error)
	GetAcessKey(lastNumber *int64, cnf *string, info *NfeInfo) (*string, error)
	GetcDv(acessKey *string) (*string, error)
	GetFullAcessKey(acessKey *string, Dv *string) (*string, error)
}

type OpticaRepository interface {
	GetClientInfo(id uuid.UUID) (*Dest, error)
	GetCompanyInfo(id uuid.UUID) (*Emit, error)
}
