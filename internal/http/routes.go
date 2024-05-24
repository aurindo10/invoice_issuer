package server

import (
	"net/http"

	"github.com/aurindo10/invoice_issuer/internal/http/handlers"
)

func AddRoutes(c *http.ServeMux) {
	c.HandleFunc("POST /register-cetificate", handlers.RegisterCertificate)
}
