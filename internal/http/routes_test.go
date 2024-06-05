package server_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	companyentitie "github.com/aurindo10/invoice_issuer/internal/entities/companyEntitie"
	"github.com/aurindo10/invoice_issuer/internal/http/handlers"
	"github.com/aurindo10/invoice_issuer/internal/http/middlewares"
	"github.com/aurindo10/invoice_issuer/pkg/utils"
)

func TestGetTokenEndPoint(t *testing.T) {
	ts := httptest.NewServer(middlewares.InternalAcessMiddleware(handlers.GetTokenHandler()))

	company := companyentitie.Company{
		CNPJ:           "90982389081239021",
		RAZAO_SOCIAL:   "Sol Engenharia",
		Owner:          "KJSLKDJSKDDKSDJSDIKDJ",
		FoundationDate: time.Now(),
	}
	body, error := json.Marshal(&company)
	if error != nil {
		t.Fatal(error)
	}
	request, error := http.NewRequest("GET", ts.URL, bytes.NewBuffer(body))
	if error != nil {
		t.Fatal(error)
	}
	token := utils.GetEnv("SECRETKEY", "djskhjdkahskjdhakjshdjask")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer"+token)
	response, err := ts.Client().Do(request)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %v", response.Status)
	}
}
