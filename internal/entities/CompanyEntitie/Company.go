package companyentitie

import (
	"context"
	"reflect"
	"time"
)

type Company struct {
	CNPJ           string
	RAZAO_SOCIAL   string
	Owner          string
	FoundationDate time.Time
}
type CompanyParams struct {
	CNPJ           *string    `json:"cnpj"`
	RAZAO_SOCIAL   *string    `json:"razao_social"`
	Owner          *string    `json:"owner"`
	FoundationDate *time.Time `json:"foundation_date"`
}

func (p CompanyParams) Valid(ctx context.Context) (problems map[string]string) {
	problems = make(map[string]string)

	v := reflect.ValueOf(p)
	t := reflect.TypeOf(p)

	if v.Kind() != reflect.Struct {
		problems["error"] = "O tipo de dados não é uma estrutura"
		return problems
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Name

		if field.IsNil() {
			problems[fieldName] = "Campo está nulo"
		}
	}

	return problems
}
