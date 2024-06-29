package nfeidrepository

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	nfeentitie "github.com/aurindo10/invoice_issuer/internal/entities/nfeEntitie"
)

type IdRepository struct {
}

func (c *IdRepository) GetCnf() (*string, error) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(99999999) + 1
	cnf := fmt.Sprintf("%08d", n)
	return &cnf, nil
}

func (c *IdRepository) GetAcessKey(lastNumber *int64, cnf *string, info *nfeentitie.NfeInfo) (*string, error) {
	now := time.Now()
	year := now.Year() % 100
	month := int(now.Month())

	// Ensure CNPJ is numeric and 14 digits
	cnpj := strings.ReplaceAll(info.Cnpj, ".", "")
	cnpj = strings.ReplaceAll(cnpj, "/", "")
	cnpj = strings.ReplaceAll(cnpj, "-", "")
	if len(cnpj) != 14 {
		return nil, fmt.Errorf("CNPJ inválido")
	}

	// Ensure lastNumber is not zero
	if *lastNumber == 0 {
		return nil, fmt.Errorf("número da NF não pode ser zero")
	}

	accessKey := fmt.Sprintf("%02s%02d%02d%s%02s%03s%09d%01s%08s",
		info.Cuf,
		year,
		month,
		cnpj,
		info.Mod,
		info.Serie,
		*lastNumber,
		info.TpEmis,
		*cnf)

	if len(accessKey) != 43 {
		return nil, fmt.Errorf("chave de acesso gerada com tamanho incorreto: %d", len(accessKey))
	}

	return &accessKey, nil
}

func (c *IdRepository) GetcDv(accessKey *string) (*string, error) {
	if len(*accessKey) != 43 {
		return nil, fmt.Errorf("base da chave de acesso deve ter 43 dígitos")
	}

	weights := []int{2, 3, 4, 5, 6, 7, 8, 9}
	sum := 0
	for i := len(*accessKey) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string((*accessKey)[i]))
		if err != nil {
			return nil, fmt.Errorf("caractere inválido na chave de acesso: %c", (*accessKey)[i])
		}
		sum += digit * weights[(len(*accessKey)-1-i)%len(weights)]
	}

	remainder := sum % 11
	dv := 11 - remainder
	if dv >= 10 {
		dv = 0
	}
	res := strconv.Itoa(dv)
	return &res, nil
}

func (c *IdRepository) GetFullAcessKey(acessKey *string, Dv *string) (*string, error) {
	fullAcessKey := fmt.Sprintf("NFe%s%s", *acessKey, *Dv)
	if len(fullAcessKey) != 47 {
		return nil, fmt.Errorf("chave de acesso completa com tamanho incorreto: %d", len(fullAcessKey))
	}
	return &fullAcessKey, nil
}

func NewIdRepository() *IdRepository {
	return &IdRepository{}
}
