package tokenentitie

import companyentitie "github.com/aurindo10/invoice_issuer/internal/entities/CompanyEntitie"

type Repository interface {
	CreteToken(p *companyentitie.Company) (*Token, error)
}
