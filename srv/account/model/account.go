package model

import (
	"database/sql/driver"
	// "time"

	"github.com/jinzhu/gorm"
	// "github.com/xmlking/micro-starter-kit/shared/model"
)

// User Entity
type User struct {
	gorm.Model        //model.GormModel
	Username   string `gorm:"size:100;not null"`
	FirstName  string `gorm:"size:255;not null"`
	LastName   string `gorm:"not null"`
	Email      string
	Profile    Profile
}

// Profile Entity
type Profile struct {
	gorm.Model //model.GormModel
	TZ         string // *time.Location
	Avatar     string
	Gender     string
	// FIXME: https://github.com/jinzhu/gorm/issues/1978
	// Gender     GenderType `gorm:"not null;type:ENUM('M', 'F')"`
	// Birthday *time.Time `gorm:"default:null"`
	UserID   uint32
}

// GenderType string
type GenderType string

// GenderTypes
const (
	Male   GenderType = "M"
	Female GenderType = "F"
)

// Scan GenderType
func (u *GenderType) Scan(value interface{}) error { *u = GenderType(value.([]byte)); return nil }

// Value GenderType
func (u GenderType) Value() (driver.Value, error) { return string(u), nil }
