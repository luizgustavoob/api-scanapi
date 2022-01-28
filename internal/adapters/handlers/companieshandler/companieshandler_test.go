package companieshandler_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/api-scanapi/internal/adapters/handlers/companieshandler"
	"github.com/api-scanapi/internal/core/entities"
	"github.com/api-scanapi/internal/core/ports/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_New(t *testing.T) {
	validator := new(mocks.CompanyValidator)
	inserter := new(mocks.CompanyInserter)

	handler := companieshandler.New(validator, inserter)

	assert.Equal(t, http.MethodPost, handler.GetHttpMethod())
	assert.Equal(t, "/companies", handler.GetRelativePath())
}

func TestHandler_ServeHTTP_WhenRequestCompanyIsInvalid_ShouldReturnError(t *testing.T) {
	validator := new(mocks.CompanyValidator)
	inserter := new(mocks.CompanyInserter)

	handler := companieshandler.New(validator, inserter)

	request := httptest.NewRequest("POST", "/companies", bytes.NewReader([]byte(`{`)))
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	expectedBody := []byte(`{"error": "Invalid company"}`)
	body, _ := ioutil.ReadAll(response.Body)

	assert.JSONEq(t, string(expectedBody), string(body))
	assert.Equal(t, http.StatusBadRequest, response.Result().StatusCode)
}

func TestHandler_ServeHTTP_WhenClientCompanyReturnError_ShouldReturnError(t *testing.T) {
	validator := new(mocks.CompanyValidator)
	validator.On("CheckCompany", mock.Anything, mock.Anything).Return(false, errors.New("error"))

	inserter := new(mocks.CompanyInserter)

	handler := companieshandler.New(validator, inserter)

	request := httptest.NewRequest("POST", "/companies", bytes.NewReader([]byte(`{}`)))
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	expectedBody := []byte(`{"error": "error"}`)
	body, _ := ioutil.ReadAll(response.Body)

	assert.JSONEq(t, string(expectedBody), string(body))
	assert.Equal(t, http.StatusInternalServerError, response.Result().StatusCode)
}

func TestHandler_ServeHTTP_WhenCompanyIsInvalid_ShouldReturnError(t *testing.T) {
	validator := new(mocks.CompanyValidator)
	validator.On("CheckCompany", mock.Anything, mock.Anything).Return(false, nil)

	inserter := new(mocks.CompanyInserter)

	handler := companieshandler.New(validator, inserter)

	request := httptest.NewRequest("POST", "/companies", bytes.NewReader([]byte(`{}`)))
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	expectedBody := []byte(`{"error": "Invalid company"}`)
	body, _ := ioutil.ReadAll(response.Body)

	assert.JSONEq(t, string(expectedBody), string(body))
	assert.Equal(t, http.StatusBadRequest, response.Result().StatusCode)
}

func TestHandler_ServeHTTP_WhenInserterReturnError_ShouldReturnError(t *testing.T) {
	validator := new(mocks.CompanyValidator)
	validator.On("CheckCompany", mock.Anything, mock.Anything).Return(true, nil)

	inserter := new(mocks.CompanyInserter)
	inserter.On("AddCompany", mock.Anything).Return(nil, errors.New("error"))

	handler := companieshandler.New(validator, inserter)

	request := httptest.NewRequest("POST", "/companies", bytes.NewReader([]byte(`{}`)))
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	expectedBody := []byte(`{"error": "error"}`)
	body, _ := ioutil.ReadAll(response.Body)

	assert.JSONEq(t, string(expectedBody), string(body))
	assert.Equal(t, http.StatusInternalServerError, response.Result().StatusCode)
}

func TestHandler_ServeHTTP_WhenSuccess_ShouldReturnSuccess(t *testing.T) {
	expected := &entities.Company{
		ID:          "1",
		RazaoSocial: "NEOWAY TECNOLOGIA INTEGRADA ASSESSORIA E NEGOCIOS SA",
		CNPJ:        "05337875000105",
		Cidade:      "FLORIANOPOLIS",
		UF:          "SC",
	}

	validator := new(mocks.CompanyValidator)
	validator.On("CheckCompany", mock.Anything, mock.Anything).Return(true, nil)

	inserter := new(mocks.CompanyInserter)
	inserter.On("AddCompany", mock.Anything).Return(expected, nil)

	handler := companieshandler.New(validator, inserter)

	company := `{"razaoSocial": "NEOWAY TECNOLOGIA INTEGRADA ASSESSORIA E NEGOCIOS SA", "cnpj": "05337875000105", "cidade": "FLORIANOPOLIS", "uf": "SC"}`
	companySave := `{"id": "1", "razaoSocial": "NEOWAY TECNOLOGIA INTEGRADA ASSESSORIA E NEGOCIOS SA", "cnpj": "05337875000105", "cidade": "FLORIANOPOLIS", "uf": "SC"}`

	request := httptest.NewRequest("POST", "/companies", bytes.NewReader([]byte(company)))
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	expectedBody := []byte(companySave)
	body, _ := ioutil.ReadAll(response.Body)

	assert.JSONEq(t, string(expectedBody), string(body))
	assert.Equal(t, http.StatusCreated, response.Result().StatusCode)
}
