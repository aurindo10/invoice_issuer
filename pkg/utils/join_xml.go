package utils

import (
	"fmt"
)

func JoinXml(final *[]byte) (*[]byte, error) {
	xml := fmt.Sprintf(
		`<enviNFe xmlns="http://www.portalfiscal.inf.br/nfe" versao="4.00"><idLote>1234</idLote><indSinc>1</indSinc><NFe xmlns="http://www.portalfiscal.inf.br/nfe">%s</NFe></enviNFe>`,
		string(*final),
	)
	res := []byte(xml)
	return &res, nil
}
