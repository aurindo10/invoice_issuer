package nfeservice_test

import (
	"testing"

	"github.com/aurindo10/invoice_issuer/internal/db"
	nfeentitie "github.com/aurindo10/invoice_issuer/internal/entities/nfeEntitie"
	clientrepository "github.com/aurindo10/invoice_issuer/internal/repositories/clientRepository"
	nfeidrepository "github.com/aurindo10/invoice_issuer/internal/repositories/nfeIdRepository"
	nfeservice "github.com/aurindo10/invoice_issuer/internal/service/nfeService"
	nfeusecase "github.com/aurindo10/invoice_issuer/internal/usecase/nfeUseCase"
	"github.com/google/uuid"
)

func TestCreateXmlAndValidateService(t *testing.T) {
	db := db.NewDbConnection()
	repo := nfeidrepository.NewIdRepository(db)
	generateNfeIduseCase := nfeusecase.NewGenerateID(repo)
	opticaRepo := clientrepository.NewOpticaRepository(db)
	clientUseCase := nfeusecase.NewClientUseCase(opticaRepo)
	clientInfo, error := clientUseCase.Execute(uuid.New())
	if error != nil {
		t.Errorf("expected no error, got this error %v", error)
	}
	companyUseCase := nfeusecase.NewCompanyUseCase(opticaRepo)
	companyInfo, error := companyUseCase.Execute(uuid.New())
	if error != nil {
		t.Errorf("expected no error, got this error %v", error)
	}
	params := nfeentitie.NfeInfo{
		Cuf:    "35",             // Estado (por exemplo, São Paulo)
		Cnpj:   "12345678000195", // CNPJ
		Mod:    "55",             // Modelo (55 por padrão)
		Serie:  "001",            // Série (1 ou 2)
		TpEmis: "1",
	}
	nfeId, error := generateNfeIduseCase.Execute(&params)
	if error != nil {
		t.Errorf("expected no error, got this error %v", error)
	}
	service := nfeservice.NewCreateAndValidateNFe()
	clientParams := nfeservice.CreateAndValidateNFeServiceParams{
		Id:          nfeId,
		ClientInfo:  clientInfo,
		CompanyInfo: companyInfo,
	}
	error = service.CreateAndValidateNFeService(&clientParams)
	if error != nil {
		t.Errorf("expected no error, got this error %v", error)
	}
}
