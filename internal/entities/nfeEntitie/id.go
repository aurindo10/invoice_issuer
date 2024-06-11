package nfeentitie

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
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
	db *sql.DB
}

func (c *IdRepository) GetCnf() (*string, error) {
	max := int64(100000000) // Limite superior para 8 dígitos
	n := rand.Int63n(max)
	cnf := strconv.FormatInt(n, 10)
	return &cnf, nil
}
func (c *IdRepository) GetLastNumbernNF(cnpj *string) (*int64, error) {
	var number int64 = 0
	rows, err := c.db.Query("SELECT nfeNumber FROM nfe WHERE orgId = ? ORDER BY nfeNumber DESC LIMIT 1", cnpj)
	if err != nil {
		fmt.Println("Erro ao executar query:", err)
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&number); err != nil {
			return nil, err
		}
	}
	number = number + 1
	return &number, nil
}
func (c *IdRepository) GetAcessKey(lastNumber *int64, cnf *string, info *NfeInfo) (*string, error) {
	var builder strings.Builder
	now := time.Now()
	year := now.Year() % 100
	month := int(now.Month())

	builder.WriteString(info.Cuf)                  // Código da UF
	builder.WriteString(fmt.Sprintf("%02d", year)) // Ano
	builder.WriteString(fmt.Sprintf("%02d", month))
	builder.WriteString(info.Cnpj)                        // CNPJ
	builder.WriteString(info.Mod)                         // Modelo
	builder.WriteString(info.Serie)                       // Série
	builder.WriteString(fmt.Sprintf("%09d", *lastNumber)) // Número da NF
	builder.WriteString(info.TpEmis)                      // Tipo de Emissão
	builder.WriteString(*cnf)                             // Código Numérico
	accessKey := builder.String()
	return &accessKey, nil
}
func (c *IdRepository) GetcDv(accessKey *string) (*string, error) {
	peso := []int{2, 3, 4, 5, 6, 7, 8, 9}
	soma := 0
	pos := 0

	for i := len(*accessKey) - 1; i >= 0; i-- {
		num := int((*accessKey)[i] - '0')
		soma += num * peso[pos]
		pos++
		if pos == len(peso) {
			pos = 0
		}
	}

	resto := soma % 11
	var dv int
	if resto == 0 || resto == 1 {
		dv = 0
	} else {
		dv = 11 - resto
	}

	dvStr := fmt.Sprintf("%d", dv)
	return &dvStr, nil
}
func NewIdRepository(db *sql.DB) *IdRepository {
	return &IdRepository{
		db: db,
	}
}
