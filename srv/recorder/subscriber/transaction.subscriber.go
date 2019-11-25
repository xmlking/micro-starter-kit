package subscriber

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/metadata"
	log "github.com/sirupsen/logrus"

	"github.com/xmlking/micro-starter-kit/shared/constants"
	pb "github.com/xmlking/micro-starter-kit/srv/recorder/proto/recorder"
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
	log.Debug("in NewTransactionSubscriber")
	return &TransactionSubscriber{
		repo:         repo,
		accountSrvEp: config.Get("services", "accountsrv", "endpoint").String("accountsrv"),
		emailerSrvEp: config.Get("services", "emailersrv", "endpoint").String("emailersrv"),
		greeterSrvEp: config.Get("services", "greeterrv", "endpoint").String("greeterrv"),
	}
}

// Handle is a method to record transaction event, Method can be of any name
func (s *TransactionSubscriber) Handle(ctx context.Context, event *pb.TransationEvent) (err error) {
	md, _ := metadata.FromContext(ctx)
	log.Debugf("TransactionSubscriber Struct: Received event %v with metadata %v", event, md)
	tranId := md[constants.TransID]

	if len(tranId) == 0 {
		log.Errorf("TransactionSubscriber: missing  TranID")
		return fmt.Errorf("TransactionSubscriber: missing  TranID")
	}
	switch from := md["Micro-From-Service"]; from {
	case s.accountSrvEp:
		err = s.repo.Write(ctx, "accountsrv"+tranId, event)
	case s.emailerSrvEp:
		err = s.repo.Write(ctx, "emailersrv"+tranId, event)
	case s.greeterSrvEp:
		err = s.repo.Write(ctx, "greetersrv"+tranId, event)
	default:
		log.Errorf("TransactionSubscriber: unknown  from: %s", from)
		return fmt.Errorf("TransactionSubscriber: unknown  from: %s", from)
	}
	if err != nil {
		log.WithError(err).Error("TransactionSubscriber Error: Unable to save to database")
	}
	return err
}
