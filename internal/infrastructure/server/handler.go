package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.neoway.com.br/companies/domain"
)

type handler struct {
	companyService domain.CompanyService
}

func NewHandler(companyService domain.CompanyService) http.Handler {
	handler := &handler{companyService: companyService}
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger(), handler.recovery())
	v1 := router.Group("/v1")
	v1.GET("/health", func(c *gin.Context) {
		result := make(map[string]string)
		result["message"] = "API no ar!"
		c.JSON(http.StatusOK, result)
	})
	v1.POST("/companies", handler.postCompany)
	return router
}

func (h *handler) recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if recovered := recover(); recovered != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
