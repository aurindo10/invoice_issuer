package companyentitie

import "time"

type Company struct {
	CNPJ           string
	RAZAO_SOCIAL   string
	Owner          string
	FoundationDate time.Time
}
