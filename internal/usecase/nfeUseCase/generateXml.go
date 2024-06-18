package nfeusecase

import (
	"encoding/xml"
	"fmt"
)

func GenerateBytesFromXml(p any) (*[]byte, error) {
	encoded, err := xml.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Println("Erro ao gerar o XML:", err)
		return nil, err
	}

	return &encoded, nil
}
