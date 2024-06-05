package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	companyentitie "github.com/aurindo10/invoice_issuer/internal/entities/companyEntitie"
)

func GetTokenHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request companyentitie.Company
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(bodyBytes, &request)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
	})
}
