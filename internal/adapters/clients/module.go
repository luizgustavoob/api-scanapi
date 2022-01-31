package clients

import (
	"net"
	"net/http"
	"time"

	"github.com/companies/internal/adapters/clients/brasilapiclient"
	"github.com/companies/internal/core/ports"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	newClient,
)

func newClient() ports.CompanyClient {
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 60 * time.Second,
			}).DialContext,
			MaxIdleConns:          1,
			IdleConnTimeout:       time.Second,
			ExpectContinueTimeout: time.Second,
			DisableKeepAlives:     true,
		},
	}

	return brasilapiclient.New(httpClient)
}
