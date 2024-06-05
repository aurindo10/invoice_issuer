package tokenservice_test

import (
	"testing"
	"time"

	companyentitie "github.com/aurindo10/invoice_issuer/internal/entities/companyEntitie"
	tokenservice "github.com/aurindo10/invoice_issuer/internal/service/tokenService"
)

func TestCreateTokenService(t *testing.T) {
	company := companyentitie.Company{
		CNPJ:           "90982389081239021",
		RAZAO_SOCIAL:   "Sol Engenharia",
		Owner:          "KJSLKDJSKDDKSDJSDIKDJ",
		FoundationDate: time.Now(),
	}
	service := tokenservice.NewCreateTokenService(company)
	res, error := service.CreateTokenService()
	if error != nil {
		t.Errorf("expected no error, got %v", error)
	}
	if res.Token == "" {
		t.Errorf("expected no error, got nothing on token %v", res.Token)
	}
}
