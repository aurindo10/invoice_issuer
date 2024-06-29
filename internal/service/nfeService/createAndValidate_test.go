package nfeservice_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	nfeentitie "github.com/aurindo10/invoice_issuer/internal/entities/nfeEntitie"
	nfeservice "github.com/aurindo10/invoice_issuer/internal/service/nfeService"
)

func TestCreateXmlAndValidateService(t *testing.T) {

	// Instância dos parâmetros
	params := nfeservice.CreateAndValidateNFeServiceParams{
		ClientInfo: &nfeentitie.Dest{
			CNPJ:  "12494754000124",
			XNome: "Client Name",
			EnderDest: nfeentitie.EnderDest{
				XLgr:    "Street Name",
				Nro:     "123",
				XBairro: "Neighborhood",
				CMun:    "1234567",
				XMun:    "City Name",
				UF:      "SP",
				CEP:     "12345678",
				CPais:   "1058",
				XPais:   "Brasil",    // Corrigido para "Brasil"
				Fone:    "123456789", // Adicionado valor para fone
			},
			IndIEDest: "1",
		},
		CompanyInfo: &nfeentitie.Emit{
			CNPJ:  "12494754000124",
			XNome: "Company Name",
			XFant: "Company Fantasy Name",
			EnderEmit: nfeentitie.EnderEmit{
				XLgr:    "Company Street",
				Nro:     "456",
				XCpl:    "Complement",
				XBairro: "Company Neighborhood",
				CMun:    "7654321",
				XMun:    "Company City",
				UF:      "RJ",
				CEP:     "87654321",
				CPais:   "1058",
				XPais:   "Brasil",    // Corrigido para "Brasil"
				Fone:    "987654321", // Adicionado valor para fone
			},
			IE:  "12345678",
			CRT: "3",
		},
		Ide: &nfeentitie.Ide{
			CUF:      "21",
			CNF:      "12345678",
			NatOp:    "Venda",
			Mod:      "55",
			Serie:    "1",
			NFNum:    "1234",
			DhEmi:    "2024-06-13T14:00:00-03:00",
			TpNF:     "1",
			IdDest:   "1",
			CMunFG:   "3550308",
			TpImp:    "1",
			TpEmis:   "1",
			CDV:      "1",
			TpAmb:    "2",
			FinNFe:   "1",
			IndFinal: "1",
			IndPres:  "1",
			ProcEmi:  "0",
			VerProc:  "1.0",
		},
		Pagamento: &[]nfeentitie.DetPag{
			{
				IndPag: "0",
				Tpag:   "01",
				VPag:   "100.00",
			},
		},
		Produtos: &[]nfeentitie.Det{
			{
				NItem: "1",
				Prod: nfeentitie.Prod{
					CProd:   "001",
					XProd:   "Product 1",
					NCM:     "12345678",
					CFOP:    "5102",
					UCom:    "UN",
					QCom:    "1.00",
					VUnCom:  "100.00",
					VProd:   "100.00",
					UTrib:   "UN",
					QTrib:   "1.00",
					VUnTrib: "100.00",
					IndTot:  "1",
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
							CST:  "99",
							VBC:  "0.00",
							PPIS: "0.00",
							VPIS: "0.00",
						},
					},
					COFINS: nfeentitie.COFINS{
						COFINSOutr: nfeentitie.COFINSOutr{
							CST:     "99",
							VBC:     "0.00",
							PCOFINS: "0.00",
							VCOFINS: "0.00",
						},
					},
				},
			},
		},
		Total: &nfeentitie.Total{
			ICMSTot: nfeentitie.ICMSTot{
				VBC:          "0.00",
				VICMS:        "0.00",
				VICMSDeson:   "0.00",
				VFCPUFDest:   "0.00",
				VICMSUFDest:  "0.00",
				VICMSUFRemet: "0.00",
				VFCP:         "0.00",
				VBCST:        "0.00",
				VST:          "0.00",
				VFCPST:       "0.00",
				VFCPSTRet:    "0.00",
				VProd:        "100.00",
				VFrete:       "0.00",
				VSeg:         "0.00",
				VDesc:        "0.00",
				VII:          "0.00",
				VIPI:         "0.00",
				VIPIDevol:    "0.00",
				VPIS:         "0.00",
				VCOFINS:      "0.00",
				VOutro:       "0.00",
				VNF:          "100.00",
				VTotTrib:     "0.00",
			},
		},
		Frete: &nfeentitie.Transp{
			ModFrete: "9",
		},
		Cobra: &nfeentitie.Cobr{
			Fat: nfeentitie.Fat{
				NFat:  "1234",
				VOrig: "100.00",
				VDesc: "0.00",
				VLiq:  "100.00",
			},
		},
		InfAdc: &nfeentitie.InfAdic{
			InfCpl: "Informações adicionais.",
		},
	}
	service := nfeservice.NewCreateAndValidateNFe()
	error := service.CreateAndValidateNFeService(&params)
	if error != nil {
		t.Errorf("expected no error, got this error %v", error)
	}
}
func TestCreateXmlAndValidateServiceEndToEnd(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Chame o serviço real aqui
		var params nfeservice.CreateAndValidateNFeServiceParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Verificação de nil para cada campo relevante
		if params.ClientInfo == nil || params.CompanyInfo == nil || params.Ide == nil || params.Pagamento == nil ||
			params.Produtos == nil || params.Total == nil || params.Frete == nil || params.Cobra == nil || params.InfAdc == nil {
			http.Error(w, "missing required field", http.StatusBadRequest)
			return
		}

		service := nfeservice.NewCreateAndValidateNFe()
		if err := service.CreateAndValidateNFeService(&params); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))

	defer ts.Close()

	// JSON dos parâmetros
	jsonParams := `{
		"client_info": {
			"cnpj": "31918175000106",
			"x_nome": "Client Name",
			"ender_dest": {
				"x_lgr": "Street Name",
				"nro": "123",
				"x_bairro": "Neighborhood",
				"c_mun": "2110708",
				"x_mun": "City Name",
				"uf": "MA",
				"cep": "65790000",
				"c_pais": "1058",
				"x_pais": "Brasil",
				"fone": "123456789"
			},
			"ind_ie_dest": "1"
		},
		"company_info": {
			"cnpj": "12494754000124",
			"x_nome": "Company Name",
			"x_fant": "Company Fantasy Name",
			"ender_emit": {
				"x_lgr": "Company Street",
				"nro": "456",
				"x_cpl": "Complement",
				"x_bairro": "Company Neighborhood",
				"c_mun": "2110708",
				"x_mun": "Company City",
				"uf": "MA",
				"cep": "65790000",
				"c_pais": "1058",
				"x_pais": "Brasil",
				"fone": "987654321"
			},
			"ie": "12345678",
			"crt": "3"
		},
		"ide": {
			"c_uf": "21",
			"nat_op": "Venda",
			"mod": "55",
			"serie": "1",
			"dh_emi": "2024-06-13T14:00:00-03:00",
			"tp_nf": "1",
			"id_dest": "1",
			"nf_num":"1",
			"c_mun_fg": "2110708",
			"tp_imp": "1",
			"tp_emis": "1",
			"c_dv": "1",
			"tp_amb": "2",
			"fin_nfe": "1",
			"ind_final": "1",
			"ind_pres": "1",
			"proc_emi": "0",
			"ver_proc": "1.0"
		},
		"pagamento": [{
			"ind_pag": "0",
			"t_pag": "01",
			"v_pag": "100.00"
		}],
		"produtos": [{
			"n_item": "1",
			"prod": {
				"c_prod": "001",
				"x_prod": "Product 1",
				"ncm": "12345678",
				"cfop": "5102",
				"u_com": "UN",
				"q_com": "1.00",
				"v_un_com": "100.00",
				"v_prod": "100.00",
				"c_ean_trib": "1234567890123",
				"u_trib": "UN",
				"q_trib": "1.00",
				"v_un_trib": "100.00",
				"ind_tot": "1"
			},
			"imposto": {
				"icms": {
					"icms_sn_102": {
						"orig": "0",
						"csosn": "102"
					}
				},
				"pis": {
					"pis_outr": {
						"cst": "99",
						"v_bc": "0.00",
						"p_pis": "0.00",
						"v_pis": "0.00"
					}
				},
				"cofins": {
					"cofins_outr": {
						"cst": "99",
						"v_bc": "0.00",
						"p_cofins": "0.00",
						"v_cofins": "0.00"
					}
				}
			}
		}],
		"total": {
			"icms_tot": {
				"v_bc": "0.00",
				"v_icms": "0.00",
				"v_icms_deson": "0.00",
				"v_fcp_uf_dest": "0.00",
				"v_icms_uf_dest": "0.00",
				"v_icms_uf_remet": "0.00",
				"v_fcp": "0.00",
				"v_bc_st": "0.00",
				"v_st": "0.00",
				"v_fcp_st": "0.00",
				"v_fcp_st_ret": "0.00",
				"v_prod": "100.00",
				"v_frete": "0.00",
				"v_seg": "0.00",
				"v_desc": "0.00",
				"v_ii": "0.00",
				"v_ipi": "0.00",
				"v_ipi_devol": "0.00",
				"v_pis": "0.00",
				"v_cofins": "0.00",
				"v_outro": "0.00",
				"v_nf": "100.00",
				"v_tot_trib": "0.00"
			}
		},
		"frete": {
			"mod_frete": "9"
		},
		"cobra": {
			"fat": {
				"n_fat": "1234",
				"v_orig": "100.00",
				"v_desc": "0.00",
				"v_liq": "100.00"
			}
		},
		"inf_adc": {
			"inf_cpl": "Informações adicionais."
		}
	}`

	request, err := http.NewRequest("POST", ts.URL, bytes.NewBuffer([]byte(jsonParams)))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := ts.Client().Do(request)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %v", response.Status)
	}

}
