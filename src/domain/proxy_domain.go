package domain

import "net/http"

type ProxyRequest struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type ProxyResponse struct {
	ID      string         `json:"id"`
	Status  int         `json:"status"`
	Headers http.Header `json:"headers"`
	Length  int64       `json:"length"`
}

type ProxyRepository interface {
	SaveRequest(request *ProxyRequest) (string, error)
	SaveResponse(response *ProxyResponse, requestId string) (string, error)
}
