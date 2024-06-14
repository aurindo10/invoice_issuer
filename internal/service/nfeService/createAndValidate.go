package nfeservice

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
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

	// Enviar para a Receita Federal
	err = c.SendNFeToReceitaFederal(*xmlB)
	if err != nil {
		return err
	}

	return nil
}

func (c *CreateAndValidateNFe) SendNFeToReceitaFederal(xmlData []byte) error {
	// Configuração do cliente HTTP
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	// URL do serviço de homologação da Receita Federal (exemplo usando SVRS)
	url := "https://hom.sefazvirtual.fazenda.gov.br/NFeAutorizacao4/NFeAutorizacao4.asmx"

	// Corpo da requisição SOAP
	soapEnvelope := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:nfe="http://www.portalfiscal.inf.br/nfe/wsdl/NFeRecepcao2">
   <soapenv:Header/>
   <soapenv:Body>
      <nfe:nfeDadosMsg>
         <![CDATA[` + string(xmlData) + `]]>
      </nfe:nfeDadosMsg>
   </soapenv:Body>
</soapenv:Envelope>`

	// Criação da requisição HTTP
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(soapEnvelope)))
	if err != nil {
		return err
	}

	// Configuração dos cabeçalhos da requisição
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", "http://www.portalfiscal.inf.br/nfe/wsdl/NFeRecepcao2/nfeRecepcaoLote2")

	// Envio da requisição
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Tratamento da resposta
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("receita federal returned status: %v", resp.Status)
	}

	// Leitura da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Exemplo de tratamento da resposta
	fmt.Printf("Receita Federal response: %s\n", string(body))

	// Processamento adicional da resposta conforme necessário

	return nil
}

func NewCreateAndValidateNFe() *CreateAndValidateNFe {
	return &CreateAndValidateNFe{}
}
