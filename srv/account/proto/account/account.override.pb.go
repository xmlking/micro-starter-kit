package accountsrv

import (
	"context"
	"errors"
	fmt "fmt"
	"net/mail"
	strings "strings"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

/**
 * GORM
**/

// AfterToPB implements the posthook interface for the Profile type. This allows
// us to customize conversion behavior. In this example, we set the User's Age
// based on the Birthday, instead of storing it separately in the DB
func (m *ProfileORM) AfterToPB(ctx context.Context, profile *Profile) error {
	profile.Age = uint32(time.Now().Sub(*m.Birthday).Hours() / 24 / 365)
	return nil
}

// BeforeCreate implements the GORM BeforeCreate interface for the UserORM type.
// you can use this method to generate new UUID for CREATE operation or let database create it with this annotation:
// {type: "uuid", primary_key: true, not_null:true, default: "uuid_generate_v4()"}];
// we prefer First method as it works with both SQLite & PostgreSQL
func (m *UserORM) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("Id", uuid.String())
}

// BeforeCreate implements the GORM BeforeCreate interface for the ProfileORM type.
func (m *ProfileORM) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("Id", uuid.String())
}

/**
 * Validate
**/
// PATCH for FIXME: https://github.com/envoyproxy/protoc-gen-validate/issues/223

func (m *UserRequest) _validateEmail(addr string) error {
	return _validateEmail(addr)
}
func (m *UserRequest) _validateHostname(host string) error {
	return _validateHostname(host)
}
func (m *UserRequest) _validateUuid(uuid string) error {
	return _validateUuid(uuid)
}
func (m *ProfileRequest) _validateEmail(addr string) error {
	return _validateEmail(addr)
}
func (m *ProfileRequest) _validateHostname(host string) error {
	return _validateHostname(host)
}
func (m *ProfileRequest) _validateUuid(uuid string) error {
	return _validateUuid(uuid)
}
func (m *UserListQuery) _validateEmail(addr string) error {
	return _validateEmail(addr)
}
func (m *UserListQuery) _validateHostname(host string) error {
	return _validateHostname(host)
}
func (m *ProfileListQuery) _validateUuid(uuid string) error {
	return _validateUuid(uuid)
}

func _validateUuid(uuid string) error {
	if matched := _account_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}
func _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}
func _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return _validateHostname(parts[1])
}
