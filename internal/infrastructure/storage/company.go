package storage

import (
	"database/sql"

	"gitlab.neoway.com.br/companies/domain"
)

type companyStorage struct {
	db *sql.DB
}

func NewCompanyStorage(db *sql.DB) *companyStorage {
	return &companyStorage{db: db}
}

func (self *companyStorage) Insert(company *domain.Company) (*domain.Company, error) {
	err := self.db.QueryRow(
		`INSERT INTO companies(razao_social, cnpj, cidade, uf) 
		VALUES ($1, $2, $3, $4) RETURNING id`,
		company.RazaoSocial, company.CNPJ, company.Cidade, company.UF).Scan(&company.ID)

	if err != nil {
		return nil, err
	}

	return company, nil
}
