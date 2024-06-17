package nfeservice

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
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
	nfeId, err := generateNfeIduseCase.Execute(nfeInfo)
	nfe.InfNFe.Id = *nfeId
	if err != nil {
		return err
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

const defaultTimeout = 15 * time.Second
const defaultUserAgent = "GoNFe/0.1"

func (c *CreateAndValidateNFe) SendNFeToReceitaFederal(xmlData []byte) error {
	// Carregar o certificado e a chave privada
	cert, err := tls.LoadX509KeyPair("client.pem", "key.pem")
	if err != nil {
		return fmt.Errorf("failed to load certificate: %v", err)
	}

	// Configuração do pool de CAs
	caCertPool, err := x509.SystemCertPool()
	if err != nil {
		return fmt.Errorf("failed to load system cert pool: %v", err)
	}
	if caCertPool == nil {
		caCertPool = x509.NewCertPool()
	}

	// Configuração do TLS com GetClientCertificate
	tlsConfig := tls.Config{
		Certificates:  []tls.Certificate{cert},
		RootCAs:       caCertPool,
		Renegotiation: tls.RenegotiateOnceAsClient,
		GetClientCertificate: func(info *tls.CertificateRequestInfo) (*tls.Certificate, error) {
			return &cert, nil
		},
		InsecureSkipVerify: true,
	}

	// Configuração do cliente HTTP
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tlsConfig,
		},
		Timeout: defaultTimeout,
	}

	// URL do serviço de homologação da Receita Federal (exemplo usando SVRS)
	url := "https://nfe-homologacao.svrs.rs.gov.br/ws/NfeAutorizacao/NFeAutorizacao4.asmx"

	// Corpo da requisição SOAP
	soapEnvelope := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap12:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap12="http://www.w3.org/2003/05/soap-envelope"  >
<soap12:Header>
<nfeCabecMsg xmlns="http://www.portalfiscal.inf.br/nfe/wsdl/NFeAutorizacao4">
<versaoDados>4.00</versaoDados>
<cUF>21</cUF> <!-- Ajuste o código da UF conforme necessário -->
</nfeCabecMsg>
</soap12:Header>
<soap12:Body>
<nfeDadosMsg xmlns="http://www.portalfiscal.inf.br/nfe/wsdl/NFeAutorizacao4">
<![CDATA[%s]]>
</nfeDadosMsg>
</soap12:Body>
</soap12:Envelope>`, xmlData)

	// Criação da requisição HTTP
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(soapEnvelope)))
	if err != nil {
		return err
	}

	// Configuração dos cabeçalhos da requisição
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", "http://www.portalfiscal.inf.br/nfe/wsdl/NFeAutorizacao4/nfeAutorizacaoLote")

	// Envio da requisição
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Tratamento da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}
	println(string(body))
	fmt.Println(string(body))
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("receita federal returned status: %v", resp.Status)
	}

	// Exemplo de tratamento da resposta
	fmt.Printf("Receita Federal response: %s\n", string(body))

	// Processamento adicional da resposta conforme necessário

	return nil
}

func NewCreateAndValidateNFe() *CreateAndValidateNFe {
	return &CreateAndValidateNFe{}
}
