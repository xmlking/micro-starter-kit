package entities

import (
    "context"
    "time"

    "github.com/jinzhu/gorm"
    uuid "github.com/satori/go.uuid"
)

/**
 * GORM
**/

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

// AfterToPB implements the posthook interface for the Profile type. This allows
// us to customize conversion behavior. In this example, we set the User's Age
// based on the Birthday, instead of storing it separately in the DB
func (m *ProfileORM) AfterToPB(ctx context.Context, profile *Profile) error {
    profile.Age = uint32(time.Now().Sub(*m.Birthday).Hours() / 24 / 365)
    return nil
}
