package nfeservice_test

import (
	"testing"

	nfeservice "github.com/aurindo10/invoice_issuer/internal/service/nfeService"
)

func TestCreateXmlAndValidateService(t *testing.T) {
	service := nfeservice.NewCreateAndValidateNFe()
	error := service.CreateAndValidateNFeService()
	if error != nil {
		t.Errorf("expected no error, got this error %v", error)
	}
}
