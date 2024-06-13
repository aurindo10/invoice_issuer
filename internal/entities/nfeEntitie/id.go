package nfeentitie

type NfeInfo struct {
	Cuf        string //state
	Cnpj       string
	Mod        string //55 by default
	Serie      string // 1 or 2
	LastNumber int64
	TpEmis     string // 1 by default
	// cNF    string // randon number
	// cdv    string // calculated number
}
