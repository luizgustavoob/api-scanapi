package ports

import "net/http"

type Handler interface {
	GetHttpMethod() string
	GetRelativePath() string
	http.Handler
}
