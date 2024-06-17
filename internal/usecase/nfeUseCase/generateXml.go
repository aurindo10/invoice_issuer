package nfeusecase

import (
	"encoding/xml"
	"fmt"

	nfeentitie "github.com/aurindo10/invoice_issuer/internal/entities/nfeEntitie"
)

type XmlNfe struct {
	nfeentitie.NFe
}

func (c *XmlNfe) Generate() (*[]byte, error) {
	encoded, err := xml.MarshalIndent(c, "", "  ")
	if err != nil {
		fmt.Println("Erro ao gerar o XML:", err)
		return nil, err
	}

	// Imprimir o XML gerado
	fmt.Println("XML Gerado:")
	fmt.Println(string(encoded))

	return &encoded, nil
}

func NewXmlNfe(c nfeentitie.NFe) *XmlNfe {
	return &XmlNfe{
		NFe: c,
	}
}
