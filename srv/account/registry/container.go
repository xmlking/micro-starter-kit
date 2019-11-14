package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/sarulabs/di/v2"

	log "github.com/sirupsen/logrus"
	"github.com/xmlking/micro-starter-kit/shared/config"
	"github.com/xmlking/micro-starter-kit/shared/database"
	logger "github.com/xmlking/micro-starter-kit/shared/log"
	"github.com/xmlking/micro-starter-kit/srv/account/handler"
	pb "github.com/xmlking/micro-starter-kit/srv/account/proto/account"
	"github.com/xmlking/micro-starter-kit/srv/account/repository"
)

// Container - provide di Container
type Container struct {
	ctn di.Container
}

// NewContainer - create new Container
func NewContainer(cfg config.ServiceConfiguration) (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := builder.Add([]di.Def{
		{
			Name:  "config",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return cfg, nil
			},
		},
		{
			Name:  "user-repository",
			Scope: di.App,
			Build: buildUserRepository,
		},
		{
			Name:  "profile-repository",
			Scope: di.App,
			Build: buildProfileRepository,
		},
		{
			Name:  "user-handler",
			Scope: di.App,
			Build: buildUserHandler,
		},
		{
			Name:  "profile-handler",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				repo := ctn.Get("profile-repository").(repository.ProfileRepository)
				logger := logger.NewLogger(cfg.Log).WithFields(map[string]interface{}{"component": "ProfileHandler"})
				return handler.NewProfileHandler(repo, logger), nil
			},
		},
		{
			Name:  "database",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return database.GetDatabaseConnection(cfg.Database)
			},
			Close: func(obj interface{}) error {
				return obj.(*gorm.DB).Close()
			},
		},
	}...); err != nil {
		return nil, err
	}

	return &Container{
		ctn: builder.Build(),
	}, nil
}

// Resolve object
func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

// Clean Container
func (c *Container) Clean() error {
	return c.ctn.Clean()
}

// Delete Container
func (c *Container) Delete() error {
	return c.ctn.Delete()
}

func buildUserRepository(ctn di.Container) (interface{}, error) {
	db := ctn.Get("database").(*gorm.DB)
	db.AutoMigrate(&pb.UserORM{})
	return repository.NewUserRepository(db), nil
}

func buildProfileRepository(ctn di.Container) (interface{}, error) {
	db := ctn.Get("database").(*gorm.DB)
	db.AutoMigrate(&pb.ProfileORM{})
	return repository.NewProfileRepository(db), nil
}

func buildUserHandler(ctn di.Container) (interface{}, error) {
	repo := ctn.Get("user-repository").(repository.UserRepository)
	return handler.NewUserHandler(repo, nil, nil), nil // FIXME inject Publisher, and greeter service
}
