package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.neoway.com.br/companies/domain"
)

func (self *handler) postCompany(c *gin.Context) {
	company := &domain.Company{}
	if err := c.BindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, self.getResponseInvalidCompany())
		return
	}

	isValid, err := self.companyService.CheckCompany(c, company)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if !isValid {
		c.JSON(http.StatusBadRequest, self.getResponseInvalidCompany())
		return
	}

	newCompany, err := self.companyService.AddCompany(company)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, newCompany)
}

func (self *handler) getResponseInvalidCompany() map[string]string {
	invalidResult := make(map[string]string)
	invalidResult["message"] = "Invalid company"
	return invalidResult
}
