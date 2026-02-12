package monitoring_service

import (
	"context"

	"gorm.io/gorm"
)

type MonitoringService interface {
	CheckDB(ctx context.Context) error
}

type monitoringService struct {
	db *gorm.DB
}

func NewMonitoringService(db *gorm.DB) MonitoringService {
	return &monitoringService{db: db}
}

func (s *monitoringService) CheckDB(ctx context.Context) error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.PingContext(ctx)
}
