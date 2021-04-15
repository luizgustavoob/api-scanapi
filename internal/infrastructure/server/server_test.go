package http_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	http "gitlab.neoway.com.br/companies/internal/infrastructure/server"
)

func TestServer_ListenAndServe(t *testing.T) {
	server := http.New("9998", nil)
	server.ListenAndServe()

	stopChan := make(chan bool)

	go func() {
		time.Sleep(1 * time.Second)
		stopChan <- true
	}()

	var result bool
	result = <-stopChan
	server.Shutdown()

	assert.True(t, result)
}
