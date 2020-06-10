package registry

import (
    "github.com/micro/go-micro/v2/store"
    mstore "github.com/micro/go-micro/v2/store/memory"
    "github.com/rs/zerolog/log"
    "github.com/sarulabs/di/v2"

    "github.com/xmlking/micro-starter-kit/service/recorder/handler"
    "github.com/xmlking/micro-starter-kit/service/recorder/repository"
    "github.com/xmlking/micro-starter-kit/service/recorder/subscriber"
    configPB "github.com/xmlking/micro-starter-kit/shared/proto/config"
)

// Container - provide di Container
type Container struct {
    ctn di.Container
}

// NewContainer - create new Container
func NewContainer(cfg configPB.Configuration) (*Container, error) {
    builder, err := di.NewBuilder()
    if err != nil {
        log.Fatal().Err(err).Msg("")
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
            Name:  "transaction-repository",
            Scope: di.App,
            Build: func(ctn di.Container) (interface{}, error) {
                store := ctn.Get("store").(store.Store)
                return repository.NewTransactionRepository(store), nil
            },
        },
        {
            Name:  "transaction-subscriber",
            Scope: di.App,
            Build: func(ctn di.Container) (interface{}, error) {
                transRepo := ctn.Get("transaction-repository").(repository.TransactionRepository)
                return subscriber.NewTransactionSubscriber(transRepo), nil
            },
        },
        {
            Name:  "transaction-handler",
            Scope: di.App,
            Build: func(ctn di.Container) (interface{}, error) {
                transRepo := ctn.Get("transaction-repository").(repository.TransactionRepository)
                return handler.NewTransactionHandler(transRepo), nil
            },
        },
        {
            Name:  "store",
            Scope: di.App,
            Build: func(ctn di.Container) (interface{}, error) {
                return mstore.NewStore(), nil
            },
            Close: func(obj interface{}) error {
                // TODO: return mstore.NewStore().Close()
                return nil
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
