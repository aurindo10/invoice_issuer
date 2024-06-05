package handlers

import (
	"encoding/json"
	"net/http"

	companyentitie "github.com/aurindo10/invoice_issuer/internal/entities/companyEntitie"
	tokenservice "github.com/aurindo10/invoice_issuer/internal/service/tokenService"
	"github.com/aurindo10/invoice_issuer/pkg/utils"
)

func GetTokenHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, problems, err := utils.DecodeValid[companyentitie.CompanyParams](r)
		if err != nil {
			if len(problems) > 0 {
				utils.Encode(w, r, 400, problems)

				return
			}
			utils.Encode(w, r, 400, err.Error())
			return
		}
		campany := companyentitie.Company{
			CNPJ:           *body.CNPJ,
			RAZAO_SOCIAL:   *body.CNPJ,
			Owner:          *body.Owner,
			FoundationDate: *body.FoundationDate,
		}
		token, error := tokenservice.NewCreateTokenService(campany).CreateTokenService()
		if error != nil {
			http.Error(w, "Bad Request", http.StatusInternalServerError)
			return
		}

		response := map[string]string{
			"token": token.Token,
		}
		res, error := json.Marshal(response)
		if error != nil {
			http.Error(w, "Bad Request", http.StatusInternalServerError)
			return
		}
		_, error = w.Write(res)
		if error != nil {
			http.Error(w, "Bad Request", http.StatusInternalServerError)
			return
		}
	})
}
