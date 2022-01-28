package brasilapiclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/api-scanapi/internal/core/entities"
	"github.com/api-scanapi/internal/core/ports"
)

const (
	url = "https://brasilapi.com.br/api"
)

type (
	HTTPClient interface {
		Get(endpoint string) (*http.Response, error)
	}

	client struct {
		httpClient HTTPClient
	}
)

func New(httpClient HTTPClient) ports.CompanyClient {
	return &client{
		httpClient: httpClient,
	}
}

func (c *client) IsValidCompany(ctx context.Context, company *entities.Company) (bool, error) {
	channelResults := make(chan interface{})
	var results []interface{}

	limit_validations := 2
	go c.checkCnpj(ctx, company.CNPJ, channelResults)
	go c.checkUf(ctx, company.UF, channelResults)

	i := 1
	for i <= limit_validations {
		result := <-channelResults
		results = append(results, result)
		i++
	}

	for _, result := range results {
		var isValid bool
		var err error

		t := reflect.TypeOf(result).String()
		switch t {
		case "bool":
			isValid = result.(bool)
		default:
			err = result.(error)
		}

		if err != nil {
			return false, err
		}

		if !isValid {
			return false, nil
		}
	}

	return true, nil
}

func (c *client) checkCnpj(ctx context.Context, cnpj string, channelResults chan interface{}) {
	endpoint := fmt.Sprintf("%s/cnpj/v1/%s", url, cnpj)
	c.checkField(ctx, endpoint, channelResults)
}

func (c *client) checkUf(ctx context.Context, uf string, channelResults chan interface{}) {
	endpoint := fmt.Sprintf("%s/ibge/uf/v1/%s", url, uf)
	c.checkField(ctx, endpoint, channelResults)
}

func (c *client) checkField(ctx context.Context, endpoint string, channelResults chan interface{}) {
	resp, err := c.httpClient.Get(endpoint)
	if err != nil {
		channelResults <- err
		return
	}

	switch resp.StatusCode {
	case 200:
		channelResults <- true
	case 404:
		channelResults <- false
	default:
		channelResults <- errors.New("unexpected error")
	}
}
