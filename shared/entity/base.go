package entity

import (
	"time"

	"github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
)

// Base contains common columns for all tables.
// CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public; ?
type Base struct {
	ID        *resource.Identifier `gorm:"type:uuid;primary_key;"` // `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt *time.Time           `json:"created_at"`
	UpdatedAt *time.Time           `json:"update_at"`
	DeletedAt *time.Time           `sql:"index" json:"deleted_at"`
}
