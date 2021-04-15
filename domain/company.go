package domain

import "context"

type Company struct {
	ID          string `json:"id,omitempty"`
	RazaoSocial string `json:"razaoSocial,omitempty"`
	CNPJ        string `json:"cnpj,omitempty"`
	Cidade      string `json:"cidade,omitempty"`
	UF          string `json:"uf,omitempty"`
}

type CompanyService interface {
	CheckCompany(ctx context.Context, c *Company) (bool, error)
	AddCompany(c *Company) (*Company, error)
}

type CompanyClient interface {
	IsValidCompany(ctx context.Context, c *Company) (bool, error)
}

type CompanyStorage interface {
	Insert(c *Company) (*Company, error)
}
