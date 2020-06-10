package handler

import (
    "context"

    "github.com/golang/protobuf/ptypes/empty"
    "github.com/rs/zerolog/log"

    transactionPB "github.com/xmlking/micro-starter-kit/service/recorder/proto/transaction"
    "github.com/xmlking/micro-starter-kit/service/recorder/repository"
    myErrors "github.com/xmlking/micro-starter-kit/shared/errors"
)

type recorderHandler struct {
    repo repository.TransactionRepository
}

func NewTransactionHandler(repo repository.TransactionRepository) transactionPB.TransactionServiceHandler {
    return &recorderHandler{
        repo: repo,
    }
}

func (h *recorderHandler) Read(ctx context.Context, req *transactionPB.ReadRequest, rsp *transactionPB.TransactionEvent) error {
    if rsp, err := h.repo.Read(ctx, req.GetKey()); err != nil {
        log.Error().Err(err).Msg("Received TransactionRepository.Read request error")
        return myErrors.AppError(myErrors.DBE, err)
    } else {
        log.Info().Msgf("Got transactionService responce %s", rsp.GetReq())
    }
    return nil
}
func (h *recorderHandler) Write(ctx context.Context, req *transactionPB.WriteRequest, rsp *empty.Empty) (err error) {
    if err := h.repo.Write(ctx, req.GetKey(), req.GetEvent()); err != nil {
        log.Error().Err(err).Msg("Received TransactionHandler.Write request error")
        return myErrors.AppError(myErrors.DBE, err)
    }
    return nil
}
