package companiespostgresql

import (
	"database/sql"

	"github.com/companies/internal/core/entities"
	"github.com/companies/internal/core/ports"
)

type (
	DatabaseWriter interface {
		QueryRow(query string, args ...interface{}) *sql.Row
	}

	companiesWriter struct {
		db DatabaseWriter
	}
)

func NewWriter(db DatabaseWriter) ports.CompanyWriter {
	return &companiesWriter{
		db: db,
	}
}

func (w *companiesWriter) Insert(company *entities.Company) (*entities.Company, error) {
	err := w.db.QueryRow(
		`INSERT INTO companies(razao_social, cnpj, cidade, uf) 
		VALUES ($1, $2, $3, $4) RETURNING id`,
		company.RazaoSocial, company.CNPJ, company.Cidade, company.UF).
		Scan(&company.ID)
	if err != nil {
		return nil, err
	}

	return company, nil
}
