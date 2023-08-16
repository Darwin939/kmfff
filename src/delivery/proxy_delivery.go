package delivery

import (
	"encoding/json"
	"net/http"

	"proxy_service/src/domain"
	"proxy_service/src/usecase"
)

type ProxyHandler struct {
	recipeUseCase *usecase.ProxyUseCase
}

func NewProxyHandler(recipeUseCase *usecase.ProxyUseCase) *ProxyHandler {
	return &ProxyHandler{
		recipeUseCase: recipeUseCase,
	}
}

func (h *ProxyHandler) Proxy(w http.ResponseWriter, r *http.Request, recipeUseCase *usecase.ProxyUseCase) {
	var requestBody domain.ProxyRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	response, err := recipeUseCase.Proxy(&requestBody)
	if err != nil {
		http.Error(w, "Failed to fetch", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, response, http.StatusOK)
}


func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
