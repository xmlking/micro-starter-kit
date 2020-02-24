package subscriber

import (
    "context"
    "fmt"

    "github.com/micro/go-micro/v2/config"
    "github.com/micro/go-micro/v2/metadata"
    "github.com/pkg/errors"
    "github.com/xmlking/logger/log"

    "github.com/xmlking/micro-starter-kit/shared/constants"
    transactionPB "github.com/xmlking/micro-starter-kit/srv/recorder/proto/transaction"
    "github.com/xmlking/micro-starter-kit/srv/recorder/repository"
)

type TransactionSubscriber struct {
    repo         repository.TransactionRepository
    accountSrvEp string
    emailerSrvEp string
    greeterSrvEp string
}

// NewTransactionSubscriber returns an instance of `TransactionSubscriber`.
func NewTransactionSubscriber(repo repository.TransactionRepository) *TransactionSubscriber {
    return &TransactionSubscriber{
        repo:         repo,
        accountSrvEp: config.Get("services", constants.ACCOUNTSRV, "endpoint").String(constants.ACCOUNTSRV),
        emailerSrvEp: config.Get("services", constants.EMAILERSRV, "endpoint").String(constants.EMAILERSRV),
        greeterSrvEp: config.Get("services", constants.GREETERSRV, "endpoint").String(constants.GREETERSRV),
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
    case s.accountSrvEp:
        err = s.repo.Write(ctx, fmt.Sprintf("%s#%s", tranId, s.accountSrvEp), event)
    case s.emailerSrvEp:
        err = s.repo.Write(ctx, fmt.Sprintf("%s#%s", tranId, s.emailerSrvEp), event)
    case s.greeterSrvEp:
        err = s.repo.Write(ctx, fmt.Sprintf("%s#%s", tranId, s.greeterSrvEp), event)
    default:
        log.Errorf("TransactionSubscriber: unknown  from: %s", from)
        return fmt.Errorf("TransactionSubscriber: unknown  from: %s", from)
    }
    if err != nil {
        log.Errorw("TransactionSubscriber Error: Unable to save to database", err)
    }
    return err
}
