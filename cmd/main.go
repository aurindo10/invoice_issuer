package main

import (
	"context"

	"github.com/aurindo10/invoice_issuer/internal/app"
)

func main() {
	if error := app.Run(context.Background()); error != nil {
		panic(error.Error())
	}
}
