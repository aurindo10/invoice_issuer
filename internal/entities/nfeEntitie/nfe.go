package nfeentitie

import "encoding/xml"

// NFe é a estrutura principal da Nota Fiscal Eletrônica
type NFe struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/nfe NFe"`
	InfNFe  InfNFe   `xml:"infNFe"`
	// InfNFeSupl InfNFeSupl `xml:"infNFeSupl,omitempty"`
	Signature *Signature `xml:"Signature,omitempty"`
}
type InfNFeSupl struct {
	QrCode string `xml:"qrCode,omitempty"`
	URL    string `xml:"urlChave,omitempty"`
}

type Pag struct {
	XMLName xml.Name `xml:"pag"`
	DetPag  []DetPag `xml:"detPag"`
}

type DetPag struct {
	XMLName xml.Name `xml:"detPag"`
	IndPag  string   `xml:"indPag"`
	Tpag    string   `xml:"tPag"`
	VPag    string   `xml:"vPag"`
}

// InfNFe contém as informações da NF-e
type InfNFe struct {
	XMLName xml.Name `xml:"infNFe"`
	Versao  string   `xml:"versao,attr"` // B02-10: Versão do layout da NF-e (Obrigatório)
	Id      string   `xml:"Id,attr"`     // B03-10: Identificador único da NF-e (Obrigatório)
	Ide     Ide      `xml:"ide"`         // B04-10: Identificação da NF-e (Obrigatório)
	Emit    Emit     `xml:"emit"`        // C01-10: Emitente da NF-e (Obrigatório)
	Dest    Dest     `xml:"dest"`        // E01-10: Destinatário da NF-e (Obrigatório)
	Det     []Det    `xml:"det"`         // H01-10: Detalhamento dos produtos e serviços (Obrigatório)
	Total   Total    `xml:"total"`       // W01-10: Totais da NF-e (Obrigatório)
	Transp  Transp   `xml:"transp"`      // X01-10: Informações sobre transporte (Obrigatório)
	Cobr    Cobr     `xml:"cobr"`        // Y01-10: Informações sobre cobrança (Opcional)
	Pag     Pag      `xml:"pag"`         // Y07-10: Pagamento (Obrigatório)
	InfAdic InfAdic  `xml:"infAdic"`     // Z01-10: Informações adicionais (Opcional)
}

// Ide contém os dados de identificação da NF-e
type Ide struct {
	CUF      string `xml:"cUF"`      // B02-20: Código da UF do emitente do Documento Fiscal (Obrigatório)
	CNF      string `xml:"cNF"`      // B03-20: Código Numérico que compõe a Chave de Acesso (Obrigatório)
	NatOp    string `xml:"natOp"`    // B04-20: Descrição da Natureza da Operação (Obrigatório)
	Mod      string `xml:"mod"`      // B06-20: Código do Modelo do Documento Fiscal (Obrigatório)
	Serie    string `xml:"serie"`    // B07-20: Série do Documento Fiscal (Obrigatório)
	NFNum    string `xml:"nNF"`      // B08-20: Número do Documento Fiscal (Obrigatório)
	DhEmi    string `xml:"dhEmi"`    // B09-20: Data e hora de emissão do Documento Fiscal (Obrigatório)
	TpNF     string `xml:"tpNF"`     // B11-20: Tipo de Operação (0=Entrada; 1=Saída) (Obrigatório)
	IdDest   string `xml:"idDest"`   // B11a-20: Identificador de destino da operação (Obrigatório)
	CMunFG   string `xml:"cMunFG"`   // B12-20: Código do Município de Ocorrência do Fato Gerador (Obrigatório)
	TpImp    string `xml:"tpImp"`    // B21-20: Tipo de impressão do DANFE (Obrigatório)
	TpEmis   string `xml:"tpEmis"`   // B22-20: Tipo de emissão da NF-e (Obrigatório)
	CDV      string `xml:"cDV"`      // B23-20: Dígito Verificador da Chave de Acesso da NF-e (Obrigatório)
	TpAmb    string `xml:"tpAmb"`    // B24-20: Identificação do Ambiente (1=Produção; 2=Homologação) (Obrigatório)
	FinNFe   string `xml:"finNFe"`   // B25-20: Finalidade de emissão da NF-e (Obrigatório)
	IndFinal string `xml:"indFinal"` // B25a-20: Indicador de operação com consumidor final (Obrigatório)
	IndPres  string `xml:"indPres"`  // B25b-20: Indicador de presença do comprador (Obrigatório)
	ProcEmi  string `xml:"procEmi"`  // B26-20: Processo de emissão da NF-e (Obrigatório)
	VerProc  string `xml:"verProc"`  // B27-20: Versão do processo de emissão (Obrigatório)
}

// Emit contém os dados do emitente da NF-e
type Emit struct {
	CNPJ          string    `xml:"CNPJ,omitempty"`  // C02-20: CNPJ do emitente (Obrigatório)
	XNome         string    `xml:"xNome,omitempty"` // C03-20: Razão social ou nome do emitente (Obrigatório)
	XFant         string    `xml:"xFant"`           // C04-20: Nome fantasia do emitente (Opcional)
	EnderEmit     EnderEmit `xml:"enderEmit"`       // C05-20: Endereço do emitente (Obrigatório)
	IE            string    `xml:"IE,omitempty"`    // C17-20: Inscrição Estadual do emitente (Obrigatório)
	CRT           string    `xml:"CRT"`             // C21-20: Código de regime tributário (Obrigatório)
	IdEstrangeiro string    `xml:"idEstrangeiro,omitempty"`
	CPF           string    `xml:"CPF,omitempty"`
}

// EnderEmit contém o endereço do emitente
type EnderEmit struct {
	XLgr    string `xml:"xLgr"`    // C06-20: Logradouro (Obrigatório)
	Nro     string `xml:"nro"`     // C07-20: Número (Obrigatório)
	XCpl    string `xml:"xCpl"`    // C08-20: Complemento (Opcional)
	XBairro string `xml:"xBairro"` // C09-20: Bairro (Obrigatório)
	CMun    string `xml:"cMun"`    // C10-20: Código do município (Obrigatório)
	XMun    string `xml:"xMun"`    // C11-20: Nome do município (Obrigatório)
	UF      string `xml:"UF"`      // C12-20: Sigla da UF (Obrigatório)
	CEP     string `xml:"CEP"`     // C13-20: Código postal (Obrigatório)
	CPais   string `xml:"cPais"`   // C14-20: Código do país (Obrigatório)
	XPais   string `xml:"xPais"`   // C15-20: Nome do país (Obrigatório)
	Fone    string `xml:"fone"`    // C16-20: Telefone (Opcional)
}

// Dest contém os dados do destinatário da NF-e
type Dest struct {
	CNPJ      string    `xml:"CNPJ,omitempty"`  // E02-20: CNPJ do destinatário (Obrigatório se CPF não for preenchido)
	CPF       string    `xml:"CPF,omitempty"`   // E03-20: CPF do destinatário (Obrigatório se CNPJ não for preenchido)
	XNome     string    `xml:"xNome,omitempty"` // E04-20: Razão social ou nome do destinatário (Obrigatório)
	EnderDest EnderDest `xml:"enderDest"`       // E05-20: Endereço do destinatário (Obrigatório)
	IndIEDest string    `xml:"indIEDest"`       // E16a-20: Indicador da IE do destinatário (Obrigatório)
	IE        string    `xml:"IE,omitempty"`    // E17-20: Inscrição Estadual do destinatário (Opcional)
}

// EnderDest contém o endereço do destinatário
type EnderDest struct {
	XLgr    string `xml:"xLgr"`           // E06-20: Logradouro (Obrigatório)
	Nro     string `xml:"nro"`            // E07-20: Número (Obrigatório)
	XBairro string `xml:"xBairro"`        // E08-20: Bairro (Obrigatório)
	CMun    string `xml:"cMun"`           // E10-20: Código do município (Obrigatório)
	XMun    string `xml:"xMun"`           // E11-20: Nome do município (Obrigatório)
	UF      string `xml:"UF"`             // E12-20: Sigla da UF (Obrigatório)
	CEP     string `xml:"CEP"`            // E13-20: Código postal (Obrigatório)
	CPais   string `xml:"cPais"`          // E14-20: Código do país (Obrigatório)
	XPais   string `xml:"xPais"`          // E15-20: Nome do país (Obrigatório)
	Fone    string `xml:"fone,omitempty"` // E16-20: Telefone (Opcional)
}

// Det contém o detalhamento dos produtos e serviços da NF-e
type Det struct {
	NItem   string  `xml:"nItem,attr"` // H02-20: Número do item (Obrigatório)
	Prod    Prod    `xml:"prod"`       // H03-20: Produto (Obrigatório)
	Imposto Imposto `xml:"imposto"`    // N01-10: Impostos (Obrigatório)
}

// Prod contém as informações do produto
type Prod struct {
	CProd    string `xml:"cProd"`    // I02-20: Código do produto (Obrigatório)
	CEAN     string `xml:"cEAN"`     // I03-20: Código EAN (Opcional)
	XProd    string `xml:"xProd"`    // I04-20: Descrição do produto (Obrigatório)
	NCM      string `xml:"NCM"`      // I05-20: Código NCM (Obrigatório)
	CFOP     string `xml:"CFOP"`     // I08-20: Código CFOP (Obrigatório)
	UCom     string `xml:"uCom"`     // I09-20: Unidade comercial (Obrigatório)
	QCom     string `xml:"qCom"`     // I10-20: Quantidade comercial (Obrigatório)
	VUnCom   string `xml:"vUnCom"`   // I10a-20: Valor unitário comercial (Obrigatório)
	VProd    string `xml:"vProd"`    // I11-20: Valor total do produto (Obrigatório)
	CEANTrib string `xml:"cEANTrib"` // I12-20: Código EAN tributário (Opcional)
	UTrib    string `xml:"uTrib"`    // I13-20: Unidade tributária (Obrigatório)
	QTrib    string `xml:"qTrib"`    // I14-20: Quantidade tributária (Obrigatório)
	VUnTrib  string `xml:"vUnTrib"`  // I14a-20: Valor unitário tributário (Obrigatório)
	IndTot   string `xml:"indTot"`   // I17-20: Indicador de total (Obrigatório)
}

// Imposto contém as informações dos impostos
type Imposto struct {
	ICMS   ICMS   `xml:"ICMS"`   // N02-10: ICMS (Obrigatório)
	PIS    PIS    `xml:"PIS"`    // Q02-10: PIS (Obrigatório)
	COFINS COFINS `xml:"COFINS"` // S02-10: COFINS (Obrigatório)
}

// ICMS contém as informações do ICMS
type ICMS struct {
	ICMSSN102 ICMSSN102 `xml:"ICMSSN102"` // N10c-10: ICMS Simples Nacional 102 (Obrigatório)
}

// ICMSSN102 contém as informações do ICMS para Simples Nacional 102
type ICMSSN102 struct {
	Orig  string `xml:"orig"`  // N11-10: Origem da mercadoria (Obrigatório)
	CSOSN string `xml:"CSOSN"` // N12-10: Código de Situação da Operação Simples Nacional (Obrigatório)
}

// PIS contém as informações do PIS
type PIS struct {
	PISOutr PISOutr `xml:"PISOutr"` // Q10-10: PIS Outras Operações (Obrigatório)
}

// PISOutr contém as informações do PIS para outras operações
type PISOutr struct {
	CST  string `xml:"CST"`  // Q10a-10: Código de Situação Tributária do PIS (Obrigatório)
	VBC  string `xml:"vBC"`  // Q10b-10: Valor da Base de Cálculo do PIS (Obrigatório)
	PPIS string `xml:"pPIS"` // Q10c-10: Alíquota do PIS (Obrigatório)
	VPIS string `xml:"vPIS"` // Q10d-10: Valor do PIS (Obrigatório)
}

// COFINS contém as informações do COFINS
type COFINS struct {
	COFINSOutr COFINSOutr `xml:"COFINSOutr"` // S10-10: COFINS Outras Operações (Obrigatório)
}

// COFINSOutr contém as informações do COFINS para outras operações
type COFINSOutr struct {
	CST     string `xml:"CST"`     // S10a-10: Código de Situação Tributária do COFINS (Obrigatório)
	VBC     string `xml:"vBC"`     // S10b-10: Valor da Base de Cálculo do COFINS (Obrigatório)
	PCOFINS string `xml:"pCOFINS"` // S10c-10: Alíquota do COFINS (Obrigatório)
	VCOFINS string `xml:"vCOFINS"` // S10d-10: Valor do COFINS (Obrigatório)
}

// Total contém os totais da NF-e
type Total struct {
	ICMSTot ICMSTot `xml:"ICMSTot"` // W02-10: Totais do ICMS (Obrigatório)
}

// ICMSTot contém os totais do ICMS
type ICMSTot struct {
	VBC          string `xml:"vBC"`          // W03-10: Valor da BC do ICMS (Obrigatório)
	VICMS        string `xml:"vICMS"`        // W04-10: Valor do ICMS (Obrigatório)
	VICMSDeson   string `xml:"vICMSDeson"`   // W04a-10: Valor do ICMS desonerado (Obrigatório)
	VFCPUFDest   string `xml:"vFCPUFDest"`   // W04c-10: Valor do FCP destinado à UF de destino (Obrigatório)
	VICMSUFDest  string `xml:"vICMSUFDest"`  // W04e-10: Valor do ICMS destinado à UF de destino (Obrigatório)
	VICMSUFRemet string `xml:"vICMSUFRemet"` // W04f-10: Valor do ICMS a ser repassado à UF de origem (Obrigatório)
	VFCP         string `xml:"vFCP"`         // W04g-10: Valor do FCP (Obrigatório)
	VBCST        string `xml:"vBCST"`        // W05-10: Valor da BC do ICMS ST (Opcional)
	VST          string `xml:"vST"`          // W06-10: Valor do ICMS ST (Opcional)
	VFCPST       string `xml:"vFCPST"`       // W06a-10: Valor do FCP retido por substituição tributária (Opcional)
	VFCPSTRet    string `xml:"vFCPSTRet"`    // W06c-10: Valor do FCP retido anteriormente por substituição tributária (Opcional)
	VProd        string `xml:"vProd"`        // W07-10: Valor total dos produtos e serviços (Obrigatório)
	VFrete       string `xml:"vFrete"`       // W08-10: Valor do frete (Opcional)
	VSeg         string `xml:"vSeg"`         // W09-10: Valor do seguro (Opcional)
	VDesc        string `xml:"vDesc"`        // W10-10: Valor do desconto (Opcional)
	VII          string `xml:"vII"`          // W11-10: Valor do II (Opcional)
	VIPI         string `xml:"vIPI"`         // W12-10: Valor do IPI (Opcional)
	VIPIDevol    string `xml:"vIPIDevol"`    // W12a-10: Valor do IPI devolvido (Opcional)
	VPIS         string `xml:"vPIS"`         // W13-10: Valor do PIS (Obrigatório)
	VCOFINS      string `xml:"vCOFINS"`      // W14-10: Valor do COFINS (Obrigatório)
	VOutro       string `xml:"vOutro"`       // W15-10: Valor de outras despesas acessórias (Opcional)
	VNF          string `xml:"vNF"`          // W16-10: Valor total da NF-e (Obrigatório)
	VTotTrib     string `xml:"vTotTrib"`     // W16a-10: Valor total dos tributos federais, estaduais e municipais (Opcional)
}

// Transp contém as informações sobre transporte
type Transp struct {
	ModFrete string `xml:"modFrete"` // X02-10: Modalidade do frete (Obrigatório)
}

// Cobr contém as informações sobre cobrança
type Cobr struct {
	Fat Fat `xml:"fat"` // Y02-10: Informações sobre a fatura (Opcional)
}

// Fat contém as informações sobre a fatura
type Fat struct {
	NFat  string `xml:"nFat"`  // Y03-10: Número da fatura (Opcional)
	VOrig string `xml:"vOrig"` // Y04-10: Valor original da fatura (Opcional)
	VDesc string `xml:"vDesc"` // Y05-10: Valor do desconto da fatura (Opcional)
	VLiq  string `xml:"vLiq"`  // Y06-10: Valor líquido da fatura (Opcional)
}

// InfAdic contém as informações adicionais da NF-e
type InfAdic struct {
	InfCpl string `xml:"infCpl"` // Z03-10: Informações adicionais de interesse do Fisco (Opcional)
}
type Signature struct {
	XMLName        xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Signature"`
	SignedInfo     SignedInfo
	SignatureValue string `xml:"SignatureValue"`
	KeyInfo        KeyInfo
}

type SignedInfo struct {
	CanonicalizationMethod CanonicalizationMethod `xml:"CanonicalizationMethod"`
	SignatureMethod        SignatureMethod        `xml:"SignatureMethod"`
	Reference              Reference              `xml:"Reference"`
}

type CanonicalizationMethod struct {
	Algorithm string `xml:",attr"`
}

type SignatureMethod struct {
	Algorithm string `xml:",attr"`
}

type Transform struct {
	Algorithm string `xml:"Algorithm,attr"`
}

type Transforms struct {
	Transform []Transform `xml:"Transform"`
}

type Reference struct {
	URI          string       `xml:"URI,attr"`
	Transforms   Transforms   `xml:"Transforms"`
	DigestMethod DigestMethod `xml:"DigestMethod"`
	DigestValue  string       `xml:"DigestValue"`
}
type DigestMethod struct {
	Algorithm string `xml:",attr"`
}

type KeyInfo struct {
	X509Data X509Data `xml:"X509Data"`
}

type X509Data struct {
	X509Certificate string `xml:"X509Certificate"`
}
