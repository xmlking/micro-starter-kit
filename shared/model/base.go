package model

import (
	"time"

	"github.com/google/uuid"
)

// GormModel Gorm Base Model
// CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public; ?
type GormModel struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
