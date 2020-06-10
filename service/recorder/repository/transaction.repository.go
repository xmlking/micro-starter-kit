package repository

import (
    "context"
    "time"

    "github.com/golang/protobuf/proto"
    "github.com/micro/go-micro/v2/store"
    "github.com/rs/zerolog/log"

    transactionPB "github.com/xmlking/micro-starter-kit/service/recorder/proto/transaction"
)

// TransactionRepository interface
type TransactionRepository interface {
    Read(ctx context.Context, key string) (transation *transactionPB.TransactionEvent, err error)
    Write(ctx context.Context, key string, transation *transactionPB.TransactionEvent) error
}

// transactionRepository struct
type transactionRepository struct {
    store store.Store
}

// NewProfileRepository returns an instance of `TransactionRepository`.
func NewTransactionRepository(store store.Store) TransactionRepository {
    return &transactionRepository{
        store: store,
    }
}

// Read: returns matching Records
func (repo *transactionRepository) Read(ctx context.Context, key string) (transation *transactionPB.TransactionEvent, err error) {
    var records []*store.Record
    records, err = repo.store.Read(key)
    if len(records) > 0 {
        transation = &transactionPB.TransactionEvent{}
        err = proto.Unmarshal(records[0].Value, transation)
    }
    return
}

// Write:
func (repo *transactionRepository) Write(ctx context.Context, key string, transation *transactionPB.TransactionEvent) error {
    log.Debug().Msgf("Writing to database: key: %s, transation: %v", key, transation)
    data, error := proto.Marshal(transation)
    if error != nil {
        return error
    }
    rec := &store.Record{
        Key:    key,
        Value:  data,
        Expiry: 100 * time.Second,
    }
    return repo.store.Write(rec)
}
