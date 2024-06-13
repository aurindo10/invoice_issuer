package nfeusecase_test

import (
	"testing"

	nfeentitie "github.com/aurindo10/invoice_issuer/internal/entities/nfeEntitie"
	nfeidrepository "github.com/aurindo10/invoice_issuer/internal/repositories/nfeIdRepository"
	nfeusecase "github.com/aurindo10/invoice_issuer/internal/usecase/nfeUseCase"
)

func TestGenerateIdSucess(t *testing.T) {
	repo := nfeidrepository.NewIdRepository()
	useCase := nfeusecase.NewGenerateID(repo)
	params := nfeentitie.NfeInfo{
		Cuf:        "35",             // Estado (por exemplo, São Paulo)
		Cnpj:       "12345678000195", // CNPJ
		Mod:        "55",             // Modelo (55 por padrão)
		Serie:      "1",              // Série (1 ou 2)
		TpEmis:     "1",
		LastNumber: 1,
	}
	res, error := useCase.Execute(&params)
	if error != nil {
		t.Fatalf("Houve algum erro: %v", error)
	}
	println(*res)
}
