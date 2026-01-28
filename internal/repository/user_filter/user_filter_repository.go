package repository

import (
	"context"
	"database/sql"
	"selection-service/internal/domain"
	"github.com/google/uuid"
)

type UserFilterRepo struct {
	db *sql.DB
}