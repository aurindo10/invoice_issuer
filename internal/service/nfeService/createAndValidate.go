package nfeservice

import (
	"encoding/xml"

	nfeentitie "github.com/aurindo10/invoice_issuer/internal/entities/nfeEntitie"
	nfeusecase "github.com/aurindo10/invoice_issuer/internal/usecase/nfeUseCase"
)

type CreateAndValidateNFe struct {
}

func (c *CreateAndValidateNFe) CreateAndValidateNFeService() error {
	nfe := nfeentitie.NFe{
		XMLName: xml.Name{Space: "http://www.portalfiscal.inf.br/nfe", Local: "NFe"},
		InfNFe: nfeentitie.InfNFe{
			Versao: "4.00",
			Id:     "NFe21240543150071000183650010000005861149202408",
			Pag: nfeentitie.Pag{
				XMLName: xml.Name{Local: "pag"},
				DetPag: []nfeentitie.DetPag{
					{
						XMLName: xml.Name{Local: "detPag"},
						IndPag:  "0",
						Tpag:    "01",
						VPag:    "100.00",
					},
				},
			},
			Dest: nfeentitie.Dest{
				CNPJ:      "43150071000183",
				IndIEDest: "1",
				EnderDest: nfeentitie.EnderDest{
					XLgr:    "PRAÇA 1 DE MAIO",
					Nro:     "18",
					XBairro: "CENTRO",
					CMun:    "2110708",
					XMun:    "SAO DOMINGOS DO MARANHAO",
					UF:      "MA",
					CEP:     "65790000",
					CPais:   "1058",
					XPais:   "BRASIL",
					Fone:    "99992030680",
				},
			},
			Ide: nfeentitie.Ide{
				CUF:      "21",
				CNF:      "14920240",
				NatOp:    "Venda de mercadoria adquirida ou recebida de terceiros",
				Mod:      "65",
				Serie:    "1",
				NFNum:    "586",
				DhEmi:    "2024-05-29T15:02:08-03:00",
				TpNF:     "1",
				IdDest:   "1",
				CMunFG:   "2110708",
				TpImp:    "4",
				TpEmis:   "1",
				CDV:      "8",
				TpAmb:    "1",
				FinNFe:   "1",
				IndFinal: "1",
				IndPres:  "1",
				ProcEmi:  "0",
				VerProc:  "VHSYS",
			},
			Emit: nfeentitie.Emit{
				CNPJ:  "43150071000183",
				XNome: "SOL MATERIAIS ELETRICOS LTDA",
				XFant: "SOL MATERIAIS",
				EnderEmit: nfeentitie.EnderEmit{
					XLgr:    "PRAÇA 1 DE MAIO",
					Nro:     "18",
					XCpl:    "SALA B",
					XBairro: "CENTRO",
					CMun:    "2110708",
					XMun:    "SAO DOMINGOS DO MARANHAO",
					UF:      "MA",
					CEP:     "65790000",
					CPais:   "1058",
					XPais:   "BRASIL",
					Fone:    "99992030680",
				},
				IE:  "127211047",
				CRT: "1",
			},
			Det: []nfeentitie.Det{
				{
					NItem: "1",
					Prod: nfeentitie.Prod{
						CProd:    "166652",
						CEAN:     "7899563917008",
						XProd:    "CABO DE COBRE NU 16MM",
						NCM:      "74130000",
						CFOP:     "5102",
						UCom:     "m",
						QCom:     "46.0000",
						VUnCom:   "19.950000",
						VProd:    "917.70",
						CEANTrib: "7899563917008",
						UTrib:    "m",
						QTrib:    "46.0000",
						VUnTrib:  "19.950000",
						IndTot:   "1",
					},
					Imposto: nfeentitie.Imposto{
						ICMS: nfeentitie.ICMS{
							ICMSSN102: nfeentitie.ICMSSN102{
								Orig:  "0",
								CSOSN: "102",
							},
						},
						PIS: nfeentitie.PIS{
							PISOutr: nfeentitie.PISOutr{
								CST:  "49",
								VBC:  "917.70",
								PPIS: "0.00",
								VPIS: "0.00",
							},
						},
						COFINS: nfeentitie.COFINS{
							COFINSOutr: nfeentitie.COFINSOutr{
								CST:     "49",
								VBC:     "917.70",
								PCOFINS: "0.00",
								VCOFINS: "0.00",
							},
						},
					},
				},
			},
			Total: nfeentitie.Total{
				ICMSTot: nfeentitie.ICMSTot{
					VFCPUFDest:   "0.00",
					VICMSUFDest:  "0.00",
					VICMSUFRemet: "0.00",
					VBC:          "0.00",
					VICMS:        "0.00",
					VICMSDeson:   "0.00",
					VFCP:         "0.00",
					VBCST:        "0.00",
					VST:          "0.00",
					VFCPST:       "0.00",
					VFCPSTRet:    "0.00",
					VProd:        "5092.80",
					VFrete:       "0.00",
					VSeg:         "0.00",
					VDesc:        "0.00",
					VII:          "0.00",
					VIPI:         "0.00",
					VIPIDevol:    "0.00",
					VPIS:         "0.00",
					VCOFINS:      "0.00",
					VOutro:       "0.00",
					VNF:          "5092.80",
					VTotTrib:     "1706.49",
				},
			},
			Transp: nfeentitie.Transp{
				ModFrete: "9",
			},
			Cobr: nfeentitie.Cobr{
				Fat: nfeentitie.Fat{
					NFat:  "1",
					VOrig: "5092.80",
					VDesc: "0.00",
					VLiq:  "5092.80",
				},
			},

			InfAdic: nfeentitie.InfAdic{
				InfCpl: "Informações adicionais",
			},
		},
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
