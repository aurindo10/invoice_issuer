package nfeentitie

import "encoding/xml"

// NFe é a estrutura principal da Nota Fiscal Eletrônica

type EnviNFe struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/nfe enviNFe" json:"xml_name"`
	Versao  string   `xml:"versao,attr"`
	IdLote  string   `xml:"idLote"`
	IndSinc string   `xml:"indSinc"`
	NFe     NFe      `xml:"NFe"`
}

type NFe struct {
	XMLName   xml.Name   `xml:"http://www.portalfiscal.inf.br/nfe NFe" json:"xml_name"`
	InfNFe    InfNFe     `xml:"infNFe" json:"inf_nfe"`
	Signature *Signature `xml:"Signature,omitempty" json:"signature,omitempty"`
}

type InfNFeSupl struct {
	QrCode string `xml:"qrCode,omitempty" json:"qr_code,omitempty"`
	URL    string `xml:"urlChave,omitempty" json:"url,omitempty"`
}

type Pag struct {
	XMLName xml.Name `xml:"pag" json:"xml_name"`
	DetPag  []DetPag `xml:"detPag" json:"det_pag"`
}

type DetPag struct {
	XMLName xml.Name `xml:"detPag" json:"xml_name"`
	IndPag  string   `xml:"indPag" json:"ind_pag"`
	Tpag    string   `xml:"tPag" json:"t_pag"`
	VPag    string   `xml:"vPag" json:"v_pag"`
}

// InfNFe contém as informações da NF-e
type InfNFe struct {
	XMLName xml.Name `xml:"infNFe" json:"xml_name"`
	Versao  string   `xml:"versao,attr" json:"versao"` // B02-10: Versão do layout da NF-e (Obrigatório)
	Id      string   `xml:"Id,attr" json:"id"`         // B03-10: Identificador único da NF-e (Obrigatório)
	Ide     Ide      `xml:"ide" json:"ide"`            // B04-10: Identificação da NF-e (Obrigatório)
	Emit    Emit     `xml:"emit" json:"emit"`          // C01-10: Emitente da NF-e (Obrigatório)
	Dest    Dest     `xml:"dest" json:"dest"`          // E01-10: Destinatário da NF-e (Obrigatório)
	Det     []Det    `xml:"det" json:"det"`            // H01-10: Detalhamento dos produtos e serviços (Obrigatório)
	Total   Total    `xml:"total" json:"total"`        // W01-10: Totais da NF-e (Obrigatório)
	Transp  Transp   `xml:"transp" json:"transp"`      // X01-10: Informações sobre transporte (Obrigatório)
	Cobr    Cobr     `xml:"cobr" json:"cobr"`          // Y01-10: Informações sobre cobrança (Opcional)
	Pag     Pag      `xml:"pag" json:"pag"`            // Y07-10: Pagamento (Obrigatório)
	InfAdic InfAdic  `xml:"infAdic" json:"inf_adic"`   // Z01-10: Informações adicionais (Opcional)
}

// Ide contém os dados de identificação da NF-e
type Ide struct {
	CUF      string `xml:"cUF" json:"c_uf"`           // B02-20: Código da UF do emitente do Documento Fiscal (Obrigatório)
	CNF      string `xml:"cNF" json:"c_nf"`           // B03-20: Código Numérico que compõe a Chave de Acesso (Obrigatório)
	NatOp    string `xml:"natOp" json:"nat_op"`       // B04-20: Descrição da Natureza da Operação (Obrigatório)
	Mod      string `xml:"mod" json:"mod"`            // B06-20: Código do Modelo do Documento Fiscal (Obrigatório)
	Serie    string `xml:"serie" json:"serie"`        // B07-20: Série do Documento Fiscal (Obrigatório)
	NFNum    string `xml:"nNF" json:"nf_num"`         // B08-20: Número do Documento Fiscal (Obrigatório)
	DhEmi    string `xml:"dhEmi" json:"dh_emi"`       // B09-20: Data e hora de emissão do Documento Fiscal (Obrigatório)
	TpNF     string `xml:"tpNF" json:"tp_nf"`         // B11-20: Tipo de Operação (0=Entrada; 1=Saída) (Obrigatório)
	IdDest   string `xml:"idDest" json:"id_dest"`     // B11a-20: Identificador de destino da operação (Obrigatório)
	CMunFG   string `xml:"cMunFG" json:"c_mun_fg"`    // B12-20: Código do Município de Ocorrência do Fato Gerador (Obrigatório)
	TpImp    string `xml:"tpImp" json:"tp_imp"`       // B21-20: Tipo de impressão do DANFE (Obrigatório)
	TpEmis   string `xml:"tpEmis" json:"tp_emis"`     // B22-20: Tipo de emissão da NF-e (Obrigatório)
	CDV      string `xml:"cDV" json:"c_dv"`           // B23-20: Dígito Verificador da Chave de Acesso da NF-e (Obrigatório)
	TpAmb    string `xml:"tpAmb" json:"tp_amb"`       // B24-20: Identificação do Ambiente (1=Produção; 2=Homologação) (Obrigatório)
	FinNFe   string `xml:"finNFe" json:"fin_nfe"`     // B25-20: Finalidade de emissão da NF-e (Obrigatório)
	IndFinal string `xml:"indFinal" json:"ind_final"` // B25a-20: Indicador de operação com consumidor final (Obrigatório)
	IndPres  string `xml:"indPres" json:"ind_pres"`   // B25b-20: Indicador de presença do comprador (Obrigatório)
	ProcEmi  string `xml:"procEmi" json:"proc_emi"`   // B26-20: Processo de emissão da NF-e (Obrigatório)
	VerProc  string `xml:"verProc" json:"ver_proc"`   // B27-20: Versão do processo de emissão (Obrigatório)
}

// Emit contém os dados do emitente da NF-e
type Emit struct {
	CNPJ          string    `xml:"CNPJ,omitempty" json:"cnpj,omitempty"`    // C02-20: CNPJ do emitente (Obrigatório)
	XNome         string    `xml:"xNome,omitempty" json:"x_nome,omitempty"` // C03-20: Razão social ou nome do emitente (Obrigatório)
	XFant         string    `xml:"xFant" json:"x_fant"`                     // C04-20: Nome fantasia do emitente (Opcional)
	EnderEmit     EnderEmit `xml:"enderEmit" json:"ender_emit"`             // C05-20: Endereço do emitente (Obrigatório)
	IE            string    `xml:"IE,omitempty" json:"ie,omitempty"`        // C17-20: Inscrição Estadual do emitente (Obrigatório)
	CRT           string    `xml:"CRT" json:"crt"`                          // C21-20: Código de regime tributário (Obrigatório)
	IdEstrangeiro string    `xml:"idEstrangeiro,omitempty" json:"id_estrangeiro,omitempty"`
	CPF           string    `xml:"CPF,omitempty" json:"cpf,omitempty"`
}

// EnderEmit contém o endereço do emitente
type EnderEmit struct {
	XLgr    string `xml:"xLgr" json:"x_lgr"`       // C06-20: Logradouro (Obrigatório)
	Nro     string `xml:"nro" json:"nro"`          // C07-20: Número (Obrigatório)
	XCpl    string `xml:"xCpl" json:"x_cpl"`       // C08-20: Complemento (Opcional)
	XBairro string `xml:"xBairro" json:"x_bairro"` // C09-20: Bairro (Obrigatório)
	CMun    string `xml:"cMun" json:"c_mun"`       // C10-20: Código do município (Obrigatório)
	XMun    string `xml:"xMun" json:"x_mun"`       // C11-20: Nome do município (Obrigatório)
	UF      string `xml:"UF" json:"uf"`            // C12-20: Sigla da UF (Obrigatório)
	CEP     string `xml:"CEP" json:"cep"`          // C13-20: Código postal (Obrigatório)
	CPais   string `xml:"cPais" json:"c_pais"`     // C14-20: Código do país (Obrigatório)
	XPais   string `xml:"xPais" json:"x_pais"`     // C15-20: Nome do país (Obrigatório)
	Fone    string `xml:"fone" json:"fone"`        // C16-20: Telefone (Opcional)
}

// Dest contém os dados do destinatário da NF-e
type Dest struct {
	CNPJ      string    `xml:"CNPJ,omitempty" json:"cnpj,omitempty"`    // E02-20: CNPJ do destinatário (Obrigatório se CPF não for preenchido)
	CPF       string    `xml:"CPF,omitempty" json:"cpf,omitempty"`      // E03-20: CPF do destinatário (Obrigatório se CNPJ não for preenchido)
	XNome     string    `xml:"xNome,omitempty" json:"x_nome,omitempty"` // E04-20: Razão social ou nome do destinatário (Obrigatório)
	EnderDest EnderDest `xml:"enderDest" json:"ender_dest"`             // E05-20: Endereço do destinatário (Obrigatório)
	IndIEDest string    `xml:"indIEDest" json:"ind_ie_dest"`            // E16a-20: Indicador da IE do destinatário (Obrigatório)
	IE        string    `xml:"IE,omitempty" json:"ie,omitempty"`        // E17-20: Inscrição Estadual do destinatário (Opcional)
}

// EnderDest contém o endereço do destinatário
type EnderDest struct {
	XLgr    string `xml:"xLgr" json:"x_lgr"`                    // E06-20: Logradouro (Obrigatório)
	Nro     string `xml:"nro" json:"nro"`                       // E07-20: Número (Obrigatório)
	XBairro string `xml:"xBairro" json:"x_bairro"`              // E08-20: Bairro (Obrigatório)
	CMun    string `xml:"cMun" json:"c_mun"`                    // E10-20: Código do município (Obrigatório)
	XMun    string `xml:"xMun" json:"x_mun"`                    // E11-20: Nome do município (Obrigatório)
	UF      string `xml:"UF" json:"uf"`                         // E12-20: Sigla da UF (Obrigatório)
	CEP     string `xml:"CEP" json:"cep"`                       // E13-20: Código postal (Obrigatório)
	CPais   string `xml:"cPais" json:"c_pais"`                  // E14-20: Código do país (Obrigatório)
	XPais   string `xml:"xPais" json:"x_pais"`                  // E15-20: Nome do país (Obrigatório)
	Fone    string `xml:"fone,omitempty" json:"fone,omitempty"` // E16-20: Telefone (Opcional)
}

// Det contém o detalhamento dos produtos e serviços da NF-e
type Det struct {
	NItem   string  `xml:"nItem,attr" json:"n_item"` // H02-20: Número do item (Obrigatório)
	Prod    Prod    `xml:"prod" json:"prod"`         // H03-20: Produto (Obrigatório)
	Imposto Imposto `xml:"imposto" json:"imposto"`   // N01-10: Impostos (Obrigatório)
}

// Prod contém as informações do produto
type Prod struct {
	CProd    string `xml:"cProd" json:"c_prod"`        // I02-20: Código do produto (Obrigatório)
	CEAN     string `xml:"cEAN" json:"c_ean"`          // I03-20: Código EAN (Opcional)
	XProd    string `xml:"xProd" json:"x_prod"`        // I04-20: Descrição do produto (Obrigatório)
	NCM      string `xml:"NCM" json:"ncm"`             // I05-20: Código NCM (Obrigatório)
	CFOP     string `xml:"CFOP" json:"cfop"`           // I08-20: Código CFOP (Obrigatório)
	UCom     string `xml:"uCom" json:"u_com"`          // I09-20: Unidade comercial (Obrigatório)
	QCom     string `xml:"qCom" json:"q_com"`          // I10-20: Quantidade comercial (Obrigatório)
	VUnCom   string `xml:"vUnCom" json:"v_un_com"`     // I10a-20: Valor unitário comercial (Obrigatório)
	VProd    string `xml:"vProd" json:"v_prod"`        // I11-20: Valor total do produto (Obrigatório)
	CEANTrib string `xml:"cEANTrib" json:"c_ean_trib"` // I12-20: Código EAN tributário (Opcional)
	UTrib    string `xml:"uTrib" json:"u_trib"`        // I13-20: Unidade tributária (Obrigatório)
	QTrib    string `xml:"qTrib" json:"q_trib"`        // I14-20: Quantidade tributária (Obrigatório)
	VUnTrib  string `xml:"vUnTrib" json:"v_un_trib"`   // I14a-20: Valor unitário tributário (Obrigatório)
	IndTot   string `xml:"indTot" json:"ind_tot"`      // I17-20: Indicador de total (Obrigatório)
}

// Imposto contém as informações dos impostos
type Imposto struct {
	ICMS   ICMS   `xml:"ICMS" json:"icms"`     // N02-10: ICMS (Obrigatório)
	PIS    PIS    `xml:"PIS" json:"pis"`       // Q02-10: PIS (Obrigatório)
	COFINS COFINS `xml:"COFINS" json:"cofins"` // S02-10: COFINS (Obrigatório)
}

// ICMS contém as informações do ICMS
type ICMS struct {
	ICMSSN102 ICMSSN102 `xml:"ICMSSN102" json:"icms_sn_102"` // N10c-10: ICMS Simples Nacional 102 (Obrigatório)
}

// ICMSSN102 contém as informações do ICMS para Simples Nacional 102
type ICMSSN102 struct {
	Orig  string `xml:"orig" json:"orig"`   // N11-10: Origem da mercadoria (Obrigatório)
	CSOSN string `xml:"CSOSN" json:"csosn"` // N12-10: Código de Situação da Operação Simples Nacional (Obrigatório)
}

// PIS contém as informações do PIS
type PIS struct {
	PISOutr PISOutr `xml:"PISOutr" json:"pis_outr"` // Q10-10: PIS Outras Operações (Obrigatório)
}

// PISOutr contém as informações do PIS para outras operações
type PISOutr struct {
	CST  string `xml:"CST" json:"cst"`    // Q10a-10: Código de Situação Tributária do PIS (Obrigatório)
	VBC  string `xml:"vBC" json:"v_bc"`   // Q10b-10: Valor da Base de Cálculo do PIS (Obrigatório)
	PPIS string `xml:"pPIS" json:"p_pis"` // Q10c-10: Alíquota do PIS (Obrigatório)
	VPIS string `xml:"vPIS" json:"v_pis"` // Q10d-10: Valor do PIS (Obrigatório)
}

// COFINS contém as informações do COFINS
type COFINS struct {
	COFINSOutr COFINSOutr `xml:"COFINSOutr" json:"cofins_outr"` // S10-10: COFINS Outras Operações (Obrigatório)
}

// COFINSOutr contém as informações do COFINS para outras operações
type COFINSOutr struct {
	CST     string `xml:"CST" json:"cst"`          // S10a-10: Código de Situação Tributária do COFINS (Obrigatório)
	VBC     string `xml:"vBC" json:"v_bc"`         // S10b-10: Valor da Base de Cálculo do COFINS (Obrigatório)
	PCOFINS string `xml:"pCOFINS" json:"p_cofins"` // S10c-10: Alíquota do COFINS (Obrigatório)
	VCOFINS string `xml:"vCOFINS" json:"v_cofins"` // S10d-10: Valor do COFINS (Obrigatório)
}

// Total contém os totais da NF-e
type Total struct {
	ICMSTot ICMSTot `xml:"ICMSTot" json:"icms_tot"` // W02-10: Totais do ICMS (Obrigatório)
}

// ICMSTot contém os totais do ICMS
type ICMSTot struct {
	VBC          string `xml:"vBC" json:"v_bc"`                     // W03-10: Valor da BC do ICMS (Obrigatório)
	VICMS        string `xml:"vICMS" json:"v_icms"`                 // W04-10: Valor do ICMS (Obrigatório)
	VICMSDeson   string `xml:"vICMSDeson" json:"v_icms_deson"`      // W04a-10: Valor do ICMS desonerado (Obrigatório)
	VFCPUFDest   string `xml:"vFCPUFDest" json:"v_fcp_uf_dest"`     // W04c-10: Valor do FCP destinado à UF de destino (Obrigatório)
	VICMSUFDest  string `xml:"vICMSUFDest" json:"v_icms_uf_dest"`   // W04e-10: Valor do ICMS destinado à UF de destino (Obrigatório)
	VICMSUFRemet string `xml:"vICMSUFRemet" json:"v_icms_uf_remet"` // W04f-10: Valor do ICMS a ser repassado à UF de origem (Obrigatório)
	VFCP         string `xml:"vFCP" json:"v_fcp"`                   // W04g-10: Valor do FCP (Obrigatório)
	VBCST        string `xml:"vBCST" json:"v_bc_st"`                // W05-10: Valor da BC do ICMS ST (Opcional)
	VST          string `xml:"vST" json:"v_st"`                     // W06-10: Valor do ICMS ST (Opcional)
	VFCPST       string `xml:"vFCPST" json:"v_fcp_st"`              // W06a-10: Valor do FCP retido por substituição tributária (Opcional)
	VFCPSTRet    string `xml:"vFCPSTRet" json:"v_fcp_st_ret"`       // W06c-10: Valor do FCP retido anteriormente por substituição tributária (Opcional)
	VProd        string `xml:"vProd" json:"v_prod"`                 // W07-10: Valor total dos produtos e serviços (Obrigatório)
	VFrete       string `xml:"vFrete" json:"v_frete"`               // W08-10: Valor do frete (Opcional)
	VSeg         string `xml:"vSeg" json:"v_seg"`                   // W09-10: Valor do seguro (Opcional)
	VDesc        string `xml:"vDesc" json:"v_desc"`                 // W10-10: Valor do desconto (Opcional)
	VII          string `xml:"vII" json:"v_ii"`                     // W11-10: Valor do II (Opcional)
	VIPI         string `xml:"vIPI" json:"v_ipi"`                   // W12-10: Valor do IPI (Opcional)
	VIPIDevol    string `xml:"vIPIDevol" json:"v_ipi_devol"`        // W12a-10: Valor do IPI devolvido (Opcional)
	VPIS         string `xml:"vPIS" json:"v_pis"`                   // W13-10: Valor do PIS (Obrigatório)
	VCOFINS      string `xml:"vCOFINS" json:"v_cofins"`             // W14-10: Valor do COFINS (Obrigatório)
	VOutro       string `xml:"vOutro" json:"v_outro"`               // W15-10: Valor de outras despesas acessórias (Opcional)
	VNF          string `xml:"vNF" json:"v_nf"`                     // W16-10: Valor total da NF-e (Obrigatório)
	VTotTrib     string `xml:"vTotTrib" json:"v_tot_trib"`          // W16a-10: Valor total dos tributos federais, estaduais e municipais (Opcional)
}

// Transp contém as informações sobre transporte
type Transp struct {
	ModFrete string `xml:"modFrete" json:"mod_frete"` // X02-10: Modalidade do frete (Obrigatório)
}

// Cobr contém as informações sobre cobrança
type Cobr struct {
	Fat Fat `xml:"fat" json:"fat"` // Y02-10: Informações sobre a fatura (Opcional)
}

// Fat contém as informações sobre a fatura
type Fat struct {
	NFat  string `xml:"nFat" json:"n_fat"`   // Y03-10: Número da fatura (Opcional)
	VOrig string `xml:"vOrig" json:"v_orig"` // Y04-10: Valor original da fatura (Opcional)
	VDesc string `xml:"vDesc" json:"v_desc"` // Y05-10: Valor do desconto da fatura (Opcional)
	VLiq  string `xml:"vLiq" json:"v_liq"`   // Y06-10: Valor líquido da fatura (Opcional)
}

// InfAdic contém as informações adicionais da NF-e
type InfAdic struct {
	InfCpl string `xml:"infCpl" json:"inf_cpl"` // Z03-10: Informações adicionais de interesse do Fisco (Opcional)
}

type Signature struct {
	XMLName        xml.Name   `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"xml_name"`
	SignedInfo     SignedInfo `json:"signed_info"`
	SignatureValue string     `xml:"SignatureValue" json:"signature_value"`
	KeyInfo        KeyInfo    `json:"key_info"`
}

type SignedInfo struct {
	CanonicalizationMethod CanonicalizationMethod `xml:"CanonicalizationMethod" json:"canonicalization_method"`
	SignatureMethod        SignatureMethod        `xml:"SignatureMethod" json:"signature_method"`
	Reference              Reference              `xml:"Reference" json:"reference"`
}

type CanonicalizationMethod struct {
	Algorithm string `xml:",attr" json:"algorithm"`
}

type SignatureMethod struct {
	Algorithm string `xml:",attr" json:"algorithm"`
}

type Transform struct {
	Algorithm string `xml:"Algorithm,attr" json:"algorithm"`
}

type Transforms struct {
	Transform []Transform `xml:"Transform" json:"transform"`
}

type Reference struct {
	URI          string       `xml:"URI,attr" json:"uri"`
	Transforms   Transforms   `xml:"Transforms" json:"transforms"`
	DigestMethod DigestMethod `xml:"DigestMethod" json:"digest_method"`
	DigestValue  string       `xml:"DigestValue" json:"digest_value"`
}
type DigestMethod struct {
	Algorithm string `xml:",attr" json:"algorithm"`
}

type KeyInfo struct {
	X509Data X509Data `xml:"X509Data" json:"x509_data"`
}

type X509Data struct {
	X509Certificate string `xml:"X509Certificate" json:"x509_certificate"`
}
