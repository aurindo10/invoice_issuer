package nfeservice_test

import (
	"testing"

	"github.com/aurindo10/invoice_issuer/internal/db"
	nfeentitie "github.com/aurindo10/invoice_issuer/internal/entities/nfeEntitie"
	nfeidrepository "github.com/aurindo10/invoice_issuer/internal/repositories/nfeIdRepository"
	nfeservice "github.com/aurindo10/invoice_issuer/internal/service/nfeService"
	nfeusecase "github.com/aurindo10/invoice_issuer/internal/usecase/nfeUseCase"
)

func TestCreateXmlAndValidateService(t *testing.T) {
	db := db.NewDbConnection()
	repo := nfeidrepository.NewIdRepository(db)
	useCase := nfeusecase.NewGenerateID(repo)
	params := nfeentitie.NfeInfo{
		Cuf:    "35",             // Estado (por exemplo, São Paulo)
		Cnpj:   "12345678000195", // CNPJ
		Mod:    "55",             // Modelo (55 por padrão)
		Serie:  "001",            // Série (1 ou 2)
		TpEmis: "1",
	}
	res, error := useCase.Execute(&params)
	if error != nil {
		t.Errorf("expected no error, got this error %v", error)
	}
	service := nfeservice.NewCreateAndValidateNFe()
	error = service.CreateAndValidateNFeService(*res)
	if error != nil {
		t.Errorf("expected no error, got this error %v", error)
	}
}
