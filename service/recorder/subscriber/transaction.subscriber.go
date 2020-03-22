package subscriber

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2/metadata"
	"github.com/pkg/errors"
	"github.com/xmlking/logger/log"

	transactionPB "github.com/xmlking/micro-starter-kit/service/recorder/proto/transaction"
	"github.com/xmlking/micro-starter-kit/service/recorder/repository"
	"github.com/xmlking/micro-starter-kit/shared/constants"
)

type TransactionSubscriber struct {
	repo repository.TransactionRepository
}

// NewTransactionSubscriber returns an instance of `TransactionSubscriber`.
func NewTransactionSubscriber(repo repository.TransactionRepository) *TransactionSubscriber {
	return &TransactionSubscriber{
		repo: repo,
	}
}

// Handle is a method to record transaction event, Method can be of any name
func (s *TransactionSubscriber) Handle(ctx context.Context, event *transactionPB.TransactionEvent) (err error) {
	md, _ := metadata.FromContext(ctx)
	tranId := md[constants.TransID]

	if len(tranId) == 0 {
		log.Errorf("TransactionSubscriber: missing  TranID")
		return errors.New("TransactionSubscriber: missing  TranID")
	}
	switch from := md["Micro-From-Service"]; from {
	case constants.ACCOUNT_SERVICE:
		err = s.repo.Write(ctx, fmt.Sprintf("%s#%s", tranId, from), event)
	case constants.EMAILER_SERVICE:
		err = s.repo.Write(ctx, fmt.Sprintf("%s#%s", tranId, from), event)
	case constants.GREETER_SERVICE:
		err = s.repo.Write(ctx, fmt.Sprintf("%s#%s", tranId, from), event)
	default:
		log.Errorf("TransactionSubscriber: unknown  from: %s", from)
		return fmt.Errorf("TransactionSubscriber: unknown  from: %s", from)
	}
	if err != nil {
		log.WithError(err).Error("TransactionSubscriber Error: Unable to save to database")
	}
	return err
}
