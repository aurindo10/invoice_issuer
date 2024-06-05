package middlewares

import (
	"net/http"
	"strings"

	"github.com/aurindo10/invoice_issuer/pkg/utils"
)

func InternalAcessMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer")
		secretKey := utils.GetEnv("SECRETKEY", "KJKJSDNSAJHLKASLKDJASHDKJ")
		if secretKey != tokenString {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
