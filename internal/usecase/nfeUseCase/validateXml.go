package nfeusecase

import (
	"fmt"
	"log"

	xsdvalidate "github.com/terminalstatic/go-xsd-validate"
)

type ValidateXml struct {
	xml *[]byte
}

func (c *ValidateXml) Validate() (*[]byte, error) {
	xsdvalidate.Init()
	xsdPath := "./PL_009o_NT2024_001_v100/nfe_v4.00.xsd"
	xsdhandler, err := xsdvalidate.NewXsdHandlerUrl(xsdPath, xsdvalidate.ParsErrDefault)
	if err != nil {
		log.Printf("failed to open file: %s", err)
		return nil, err
	}
	defer xsdhandler.Free()
	err = xsdhandler.ValidateMem(*c.xml, xsdvalidate.ValidErrDefault)
	if err != nil {
		switch e := err.(type) {
		case xsdvalidate.ValidationError:
			fmt.Printf("Error in line: %d\n", e.Errors[0].Line)
			fmt.Println(e.Errors[0].Message)
			return nil, fmt.Errorf(e.Errors[0].Message)
		default:
			fmt.Println(err)
		}
	}
	return c.xml, nil
}

func NewValidateXml(xml *[]byte) *ValidateXml {
	return &ValidateXml{
		xml: xml,
	}
}
