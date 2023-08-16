package repository

import (
	"database/sql"
	"encoding/json"
	"log"
	"proxy_service/src/domain"

	"github.com/google/uuid"
)

type ProxyRepository struct {
	db *sql.DB
}

func NewProxyRepository(db *sql.DB) *ProxyRepository {
	return &ProxyRepository{
		db: db,
	}
}

func (r *ProxyRepository) SaveRequest(request *domain.ProxyRequest) (string, error) {
	requestID := uuid.New()

	query := "INSERT INTO request (id, header) VALUES ($1, $2)"
	header, err := json.Marshal(request.Headers)
	if err != nil {
		log.Println(err)
		return "", err
	}

	_, err = r.db.Exec(query, requestID, header)
	if err != nil {
		log.Println("failed to create request: %w", err)
		return "", err
	}
	return requestID.String(), nil
}

func (r *ProxyRepository) SaveResponse(response *domain.ProxyResponse, requestID string) (string, error) {

	query := "INSERT INTO response (id, status, header, length) VALUES ($1, $2, $3, $4)"
	header, err := json.Marshal(response.Headers)
	if err != nil {
		log.Println(err)
		return "", err
	}

	_, err = r.db.Exec(query, requestID, response.Status, header, response.Length)
	if err != nil {
		log.Println("failed to create request: %w", err)
		return "", err
	}
	return requestID, nil
}
