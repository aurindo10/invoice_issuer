package tokenentitie

import companyentitie "github.com/aurindo10/invoice_issuer/internal/entities/companyEntitie"

type Repository interface {
	CreteToken(p *companyentitie.Company) (*Token, error)
}
