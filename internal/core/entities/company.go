package entities

type Company struct {
	ID          string `json:"id,omitempty"`
	RazaoSocial string `json:"razaoSocial,omitempty"`
	CNPJ        string `json:"cnpj,omitempty"`
	Cidade      string `json:"cidade,omitempty"`
	UF          string `json:"uf,omitempty"`
}
