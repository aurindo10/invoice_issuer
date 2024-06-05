package server_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aurindo10/invoice_issuer/internal/http/handlers"
	"github.com/aurindo10/invoice_issuer/internal/http/middlewares"
	"github.com/aurindo10/invoice_issuer/pkg/utils"
)

func TestGetTokenEndPoint(t *testing.T) {
	ts := httptest.NewServer(middlewares.InternalAcessMiddleware(handlers.GetTokenHandler()))

	company := map[string]string{
		"cnpj":            "90982389081239021",
		"razao_social":    "Sol Engenharia",
		"owner":           "KJSLKDJSKDDKSDJSDIKDJ",
		"foundation_date": "2024-06-04T12:34:56.789Z",
	}
	body, error := json.Marshal(&company)
	if error != nil {
		t.Fatal(error)
	}
	request, error := http.NewRequest("POST", ts.URL, bytes.NewBuffer(body))
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
	var acessToken map[string]string
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("Houve um erro ao transformar o body em bytes")
	}

	json.Unmarshal(bodyBytes, &acessToken)
	if acessToken["token"] == "" {
		t.Fatalf("Houve um erro ao obter token")
	}
	if len(acessToken["token"]) < 10 {
		t.Fatalf("Houve um erro ao obter token")
	}
}
func TestGetTokenEndPointWrongToken(t *testing.T) {
	ts := httptest.NewServer(middlewares.InternalAcessMiddleware(handlers.GetTokenHandler()))

	company := map[string]string{
		"cnpj":            "90982389081239021",
		"razao_social":    "Sol Engenharia",
		"owner":           "KJSLKDJSKDDKSDJSDIKDJ",
		"foundation_date": "2024-06-04T12:34:56.789Z",
	}
	body, error := json.Marshal(&company)
	if error != nil {
		t.Fatal(error)
	}
	request, error := http.NewRequest("POST", ts.URL, bytes.NewBuffer(body))
	if error != nil {
		t.Fatal(error)
	}
	token := "daskdjsakldjalkdjal"
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer"+token)
	response, err := ts.Client().Do(request)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusForbidden {
		t.Fatalf("expected status 500 OK, got %v", response.Status)
	}
}
func TestMissingParams(t *testing.T) {
	ts := httptest.NewServer(middlewares.InternalAcessMiddleware(handlers.GetTokenHandler()))
	company := map[string]string{
		"cnpj":         "90982389081239021",
		"razao_social": "Sol Engenharia",
		"owner":        "KJSLKDJSKDDKSDJSDIKDJ",
		// "foundation_date": "2024-06-04T12:34:56.789Z",
	}
	body, error := json.Marshal(&company)
	if error != nil {
		t.Fatal(error)
	}
	request, error := http.NewRequest("POST", ts.URL, bytes.NewBuffer(body))
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
	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status 400 OK, got %v", response.Status)
	}
}
