CREATE TABLE companies
(
    id SERIAL,
    razao_social TEXT NOT NULL,
    cnpj TEXT NOT NULL,
    cidade TEXT NOT NULL,
    uf TEXT NOT NULL,
    CONSTRAINT companies_pkey PRIMARY KEY (id)
)