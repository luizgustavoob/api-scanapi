package companieshandler

import (
	"encoding/json"
	"net/http"

	"github.com/api-scanapi/internal/core/entities"
	"github.com/api-scanapi/internal/core/ports"
)

type handler struct {
	validator ports.CompanyValidator
	inserter  ports.CompanyInserter
}

func New(
	validator ports.CompanyValidator,
	inserter ports.CompanyInserter,
) *handler {
	return &handler{
		validator: validator,
		inserter:  inserter,
	}
}

func (h *handler) GetHttpMethod() string {
	return http.MethodPost
}

func (h *handler) GetRelativePath() string {
	return "/companies"
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var company entities.Company

	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		result := make(map[string]string)
		result["error"] = "Invalid company"
		json, _ := json.Marshal(result)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
		return
	}

	isValid, err := h.validator.CheckCompany(r.Context(), &company)
	if err != nil {
		result := make(map[string]string)
		result["error"] = err.Error()
		json, _ := json.Marshal(result)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(json)
		return
	}

	if !isValid {
		result := make(map[string]string)
		result["error"] = "Invalid company"
		json, _ := json.Marshal(result)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
		return
	}

	newCompany, err := h.inserter.AddCompany(&company)
	if err != nil {
		result := make(map[string]string)
		result["error"] = err.Error()
		json, _ := json.Marshal(result)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(json)
		return
	}

	json, _ := json.Marshal(newCompany)
	w.WriteHeader(http.StatusCreated)
	w.Write(json)
}
