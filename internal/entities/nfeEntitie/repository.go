package nfeentitie

type IdNumberRepository interface {
	GetCnf() (*string, error)
	GetAcessKey(lastNumber *int64, cnf *string, info *NfeInfo) (*string, error)
	GetcDv(acessKey *string) (*string, error)
	GetFullAcessKey(acessKey *string, Dv *string) (*string, error)
}
