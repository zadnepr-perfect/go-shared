package loggingdb

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
)

type RequestLogsModel struct {
	Conn *pgx.Conn
}

func NewRequestLogsModel(conn *pgx.Conn) *RequestLogsModel {
	return &RequestLogsModel{Conn: conn}
}

func (m *RequestLogsModel) LogRequest(ctx context.Context, endpoint string, status int, request, response, ip string) error {
	query := `
		INSERT INTO request_logs (created_at, endpoint, status, request, response, ip)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := m.Conn.Exec(ctx, query, time.Now(), endpoint, status, request, response, ip)
	if err != nil {
		log.Printf("Failed to log request: %v", err)
		return err
	}

	log.Println("Request logged successfully")
	return nil
}

// GetRowCount возвращает количество строк в таблице request_logs
func (m *RequestLogsModel) GetRowCount(ctx context.Context) (int, error) {
	var rowCount int
	err := m.Conn.QueryRow(ctx, "SELECT COUNT(*) FROM request_logs").Scan(&rowCount)
	if err != nil {
		return 0, err
	}
	return rowCount, nil
}
