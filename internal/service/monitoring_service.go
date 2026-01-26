package service

import (
	"context"
	"database/sql"
)

type MonitoringService interface {
	CheckDB(ctx context.Context) error
}

type monitoringService struct {
	db *sql.DB
}

func NewMonitoringService(db *sql.DB) MonitoringService {
	return &monitoringService{db: db}
}

func (s *monitoringService) CheckDB(ctx context.Context) error {
	return s.db.PingContext(ctx)
}
