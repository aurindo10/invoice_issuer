package server

import (
	"net/http"

	"github.com/aurindo10/invoice_issuer/internal/http/handlers"
	"github.com/aurindo10/invoice_issuer/internal/http/middlewares"
)

func AddRoutes(c *http.ServeMux) {
	c.Handle("POST /get-token", middlewares.InternalAcessMiddleware(handlers.GetTokenHandler()))
}
