package entity

import (
	"database/sql/driver"
	"time"

	"github.com/jinzhu/gorm"
	accountPB "github.com/xmlking/micro-starter-kit/srv/account/proto/account"
	// "github.com/xmlking/micro-starter-kit/shared/entity"
)

// User Entity
type User struct {
	gorm.Model        //entity.GormModel
	Username   string `gorm:"size:100;not null"`
	FirstName  string `gorm:"size:255;not null"`
	LastName   string `gorm:"not null"`
	Email      string
	Profile    Profile
}

// Profile Entity
type Profile struct {
	gorm.Model        //entity.GormModel
	TZ         string // *time.Location
	Avatar     string
	Gender     string
	// FIXME: https://github.com/jinzhu/gorm/issues/1978
	// Gender     GenderType `gorm:"not null;type:ENUM('M', 'F')"`
	// Birthday *time.Time `gorm:"default:null"`
	UserID uint32
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

// func (m *accountPB.User) ToORM() User {

// }

// ToPB convert entity to PB
func (e *User) ToPB() *accountPB.User {
	return &accountPB.User{
		Id:        uint32(e.Model.ID),
		Username:  e.Username,
		FirstName: e.FirstName,
		LastName:  e.LastName,
		Email:     e.Email,
		CreatedAt: e.Model.CreatedAt.Format(time.RFC3339),
		UpdatedAt: e.Model.UpdatedAt.Format(time.RFC3339),
	}
}

// func (m *accountPB.Profile) ToORM() Profile {

// }

// ToPB convert entity to PB
func (e *Profile) ToPB() *accountPB.Profile {
	return &accountPB.Profile{
		Id:     uint32(e.Model.ID),
		Tz:     e.TZ,
		Avatar: e.Avatar,
		Gender: e.Gender,
		// Birthday:  birthday,
		UserId:    e.UserID,
		CreatedAt: e.Model.CreatedAt.Format(time.RFC3339),
		UpdatedAt: e.Model.UpdatedAt.Format(time.RFC3339),
	}
}
