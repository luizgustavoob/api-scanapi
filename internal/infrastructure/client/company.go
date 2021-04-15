package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"gitlab.neoway.com.br/companies/domain"
	httpclient "gitlab.neoway.com.br/companies/pkg/http"
)

type companyClient struct {
	httpClient   httpclient.HTTPClient
	brasilApiUrl string
}

func NewCompanyClient(httpClient httpclient.HTTPClient, brasilApiUrl string) *companyClient {
	return &companyClient{httpClient, brasilApiUrl}
}

func (self *companyClient) IsValidCompany(ctx context.Context, c *domain.Company) (isValid bool, err error) {
	const LIMIT_VALIDATIONS = 2
	resultsChan := make(chan interface{})

	go self.checkCnpj(resultsChan, ctx, c.CNPJ)
	go self.checkUf(resultsChan, ctx, c.UF)

	validations := 0
	for result := range resultsChan {
		validations++
		t := reflect.TypeOf(result).String()
		switch t {
		case "bool":
			isValid = result.(bool)
		default:
			err = result.(error)
		}

		if !isValid || err != nil || validations == LIMIT_VALIDATIONS {
			break
		}
	}
	return
}

func (self *companyClient) checkCnpj(resultChan chan interface{}, ctx context.Context, cnpj string) {
	url := fmt.Sprintf("%s/cnpj/v1/%s", self.brasilApiUrl, cnpj)
	self.checkField(resultChan, ctx, url)
}

func (self *companyClient) checkUf(resultChan chan interface{}, ctx context.Context, uf string) {
	url := fmt.Sprintf("%s/ibge/uf/v1/%s", self.brasilApiUrl, uf)
	self.checkField(resultChan, ctx, url)
}

func (self *companyClient) checkField(resultChan chan interface{}, ctx context.Context, url string) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		resultChan <- err
		return
	}
	req = req.WithContext(ctx)

	res, err := self.httpClient.Do(req)
	if err != nil {
		resultChan <- err
		return
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		resultChan <- true
	case 404:
		resultChan <- false
	default:
		resultChan <- errors.New("Unexpected error")
	}
}
