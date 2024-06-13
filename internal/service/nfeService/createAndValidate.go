package nfeservice

import (
	"encoding/xml"
	"time"

	nfeentitie "github.com/aurindo10/invoice_issuer/internal/entities/nfeEntitie"
	nfeidrepository "github.com/aurindo10/invoice_issuer/internal/repositories/nfeIdRepository"
	nfeusecase "github.com/aurindo10/invoice_issuer/internal/usecase/nfeUseCase"
)

type CreateAndValidateNFe struct {
}
type CreateAndValidateNFeServiceParams struct {
	Id          *string              `json:"id"`
	ClientInfo  *nfeentitie.Dest     `json:"client_info"`
	CompanyInfo *nfeentitie.Emit     `json:"company_info"`
	Ide         *nfeentitie.Ide      `json:"ide"`
	Pagamento   *[]nfeentitie.DetPag `json:"pagamento"`
	Produtos    *[]nfeentitie.Det    `json:"produtos"`
	Total       *nfeentitie.Total    `json:"total"`
	Frete       *nfeentitie.Transp   `json:"frete"`
	Cobra       *nfeentitie.Cobr     `json:"cobra"`
	InfAdc      *nfeentitie.InfAdic  `json:"inf_adc"`
}

func (c *CreateAndValidateNFe) CreateAndValidateNFeService(p *CreateAndValidateNFeServiceParams) error {
	now := time.Now().Format(time.RFC3339)
	p.Ide.DhEmi = now
	nfe := nfeentitie.NFe{
		XMLName: xml.Name{Space: "http://www.portalfiscal.inf.br/nfe", Local: "NFe"},
		InfNFe: nfeentitie.InfNFe{
			Versao: "4.00",
			Pag: nfeentitie.Pag{
				XMLName: xml.Name{Local: "pag"},
				DetPag:  *p.Pagamento,
			},
			Dest:    *p.ClientInfo,
			Ide:     *p.Ide,
			Emit:    *p.CompanyInfo,
			Det:     *p.Produtos,
			Total:   *p.Total,
			Transp:  *p.Frete,
			Cobr:    *p.Cobra,
			InfAdic: *p.InfAdc,
		},
	}
	nfeInfo := &nfeentitie.NfeInfo{
		Cuf:        p.Ide.CUF,
		Cnpj:       p.CompanyInfo.CNPJ,
		Mod:        p.Ide.Mod,
		Serie:      "001",
		LastNumber: 1,
		TpEmis:     p.Ide.TpEmis,
	}
	repo := nfeidrepository.NewIdRepository()
	generateNfeIduseCase := nfeusecase.NewGenerateID(repo)
	nfeId, error := generateNfeIduseCase.Execute(nfeInfo)
	nfe.InfNFe.Id = *nfeId
	if error != nil {
		return error
	}
	crateXmlUseCase := nfeusecase.NewXmlNfe(nfe)
	xmlData, err := crateXmlUseCase.Generate()
	if err != nil {
		return err
	}

	_, sig, err := nfeusecase.SignXML("./S3D_8_240606145203.pfx", "12345678", *xmlData)
	if err != nil {
		return err
	}

	nfeAssined := nfe
	nfeAssined.Signature = sig

	crateXmlBytes := nfeusecase.NewXmlNfe(nfeAssined)
	xmlB, err := crateXmlBytes.Generate()
	if err != nil {
		return err
	}

	validate := nfeusecase.NewValidateXml(xmlB)
	_, err = validate.Validate()
	if err != nil {
		return err
	}

	return nil
}

func NewCreateAndValidateNFe() *CreateAndValidateNFe {
	return &CreateAndValidateNFe{}
}
