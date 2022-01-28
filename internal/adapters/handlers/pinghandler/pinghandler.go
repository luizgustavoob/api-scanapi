package pinghandler

import "net/http"

type handler struct{}

func New() *handler {
	return &handler{}
}

func (h *handler) GetHttpMethod() string {
	return http.MethodGet
}

func (h *handler) GetRelativePath() string {
	return "/ping"
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`pong`))
}
