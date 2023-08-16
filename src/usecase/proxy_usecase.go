package usecase

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"proxy_service/src/domain"
)

type ProxyUseCase struct {
	recipeRepository domain.ProxyRepository
}

func NewProxyUseCase(proxyRepository domain.ProxyRepository) *ProxyUseCase {
	return &ProxyUseCase{
		recipeRepository: proxyRepository,
	}
}

func getResponseBodyLength(resp *http.Response) int64 {
	if resp.ContentLength != -1 {
		return resp.ContentLength
	}

	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0
	}

	contentLength := int64(len(bodyContent))
	return contentLength
}

func (uc *ProxyUseCase) Proxy(request *domain.ProxyRequest) (domain.ProxyResponse, error) {
	requestID, err := uc.recipeRepository.SaveRequest(request)
	resp, err := makeProxyRequest(*request)
	length := getResponseBodyLength(resp)

	defer resp.Body.Close()
	responsePayload := domain.ProxyResponse{
		ID:      requestID,
		Status:  resp.StatusCode,
		Headers: resp.Header,
		Length:  length,
	}
	if err != nil {
		return responsePayload, err
	}
	uc.recipeRepository.SaveResponse(&responsePayload, requestID)

	return responsePayload, nil
}

func makeProxyRequest(req domain.ProxyRequest) (*http.Response, error) {
	client := &http.Client{}
	requestHeaders := make(http.Header)
	for key, value := range req.Headers {
		requestHeaders.Set(key, value)
	}

	requestBody := []byte{}
	if req.Method != "GET" && req.Method != "HEAD" {
		requestBody, _ = json.Marshal(req)
	}

	proxyReq, err := http.NewRequest(req.Method, req.URL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	proxyReq.Header = requestHeaders

	return client.Do(proxyReq)
}
