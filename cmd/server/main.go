package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gitlab.neoway.com.br/companies/domain/company"
	"gitlab.neoway.com.br/companies/internal/infrastructure/client"
	http "gitlab.neoway.com.br/companies/internal/infrastructure/server"
	"gitlab.neoway.com.br/companies/internal/infrastructure/storage"
	httpclient "gitlab.neoway.com.br/companies/pkg/http"
)

const BRASIL_API_URL = "https://brasilapi.com.br/api"

func main() {
	db, err := storage.NewConnection(os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Default().Fatalf("ERROR: %s", err)
		return
	}

	companyStorage := storage.NewCompanyStorage(db)
	companyClient := client.NewCompanyClient(httpclient.NewHTTPClient(60*time.Second), BRASIL_API_URL)
	companyService := company.NewCompanyService(companyStorage, companyClient)

	handler := http.NewHandler(companyService)
	server := http.New("9998", handler)
	server.ListenAndServe()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan
	server.Shutdown()
}
