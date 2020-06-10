package handler

import (
    "context"

    "github.com/jinzhu/gorm"
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/auth"
    "github.com/micro/go-micro/v2/errors"
    "github.com/rs/zerolog/log"
    uuid "github.com/satori/go.uuid"
    "github.com/thoas/go-funk"

    account_entities "github.com/xmlking/micro-starter-kit/service/account/proto/entities"
    userPB "github.com/xmlking/micro-starter-kit/service/account/proto/user"
    "github.com/xmlking/micro-starter-kit/service/account/repository"
    emailerPB "github.com/xmlking/micro-starter-kit/service/emailer/proto/emailer"
    greeterPB "github.com/xmlking/micro-starter-kit/service/greeter/proto/greeter"
    myErrors "github.com/xmlking/micro-starter-kit/shared/errors"
)

// UserHandler struct
type userHandler struct {
    userRepository   repository.UserRepository
    Event            micro.Event
    greeterSrvClient greeterPB.GreeterService
}

// NewUserHandler returns an instance of `UserServiceHandler`.
func NewUserHandler(repo repository.UserRepository, eve micro.Event, greeterClient greeterPB.GreeterService) userPB.UserServiceHandler {
    return &userHandler{
        userRepository:   repo,
        Event:            eve,
        greeterSrvClient: greeterClient,
    }
}

func (h *userHandler) Exist(ctx context.Context, req *userPB.ExistRequest, rsp *userPB.ExistResponse) error {
    log.Info().Msg("Received UserHandler.Exist request")
    model := account_entities.UserORM{}
    model.Id = uuid.FromStringOrNil(req.Id.GetValue())
    username := req.Username.GetValue()
    model.Username = &username
    model.Email = req.Email.GetValue()

    exists := h.userRepository.Exist(&model)
    log.Info().Msgf("user exists? %t", exists)
    rsp.Result = exists
    return nil
}

func (h *userHandler) List(ctx context.Context, req *userPB.ListRequest, rsp *userPB.ListResponse) error {
    log.Info().Msg("Received UserHandler.List request")
    model := account_entities.UserORM{}
    username := req.Username.GetValue()
    model.Username = &username
    model.FirstName = req.FirstName.GetValue()
    model.LastName = req.LastName.GetValue()
    model.Email = req.Email.GetValue()

    total, users, err := h.userRepository.List(req.Limit.GetValue(), req.Page.GetValue(), req.Sort.GetValue(), &model)
    if err != nil {
        return errors.NotFound("mkit.service.account.user.list", "Error %v", err.Error())
    }
    rsp.Total = total

    // newUsers := make([]*accountPB.User, len(users))
    // for index, user := range users {
    // 	tmpUser, _ := user.ToPB(ctx)
    // 	newUsers[index] = &tmpUser
    // 	// *newUsers[index], _ = user.ToPB(ctx) ???
    // }
    newUsers := funk.Map(users, func(user *account_entities.UserORM) *account_entities.User {
        tmpUser, _ := user.ToPB(ctx)
        return &tmpUser
    }).([]*account_entities.User)

    rsp.Results = newUsers
    return nil
}

func (h *userHandler) Get(ctx context.Context, req *userPB.GetRequest, rsp *userPB.GetResponse) error {
    log.Info().Msg("Received UserHandler.Get request")

    id := req.Id.GetValue()
    if id == "" {
        return myErrors.ValidationError("mkit.service.account.user.get", "validation error: Missing Id")
    }
    user, err := h.userRepository.Get(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            rsp.Result = nil
            return nil
        }
        return myErrors.AppError(myErrors.DBE, err)
    }

    tempUser, _ := user.ToPB(ctx)
    rsp.Result = &tempUser

    return nil
}

func (h *userHandler) Create(ctx context.Context, req *userPB.CreateRequest, rsp *userPB.CreateResponse) error {
    log.Info().Msg("Received UserHandler.Create request")

    model := account_entities.UserORM{}
    username := req.Username.GetValue()
    model.Username = &username
    model.FirstName = req.FirstName.GetValue()
    model.LastName = req.LastName.GetValue()
    model.Email = req.Email.GetValue()

    if err := h.userRepository.Create(&model); err != nil {
        return myErrors.AppError(myErrors.DBE, err)
    }

    // send email (TODO: async `go h.Event.Publish(...)`)
    if err := h.Event.Publish(ctx, &emailerPB.Message{To: req.Email.GetValue()}); err != nil {
        log.Error().Err(err).Msg("Received Event.Publish request error")
        return myErrors.AppError(myErrors.PSE, err)
    }

    // call greeter
    // if res, err := h.greeterSrvClient.Hello(ctx, &greeterPB.Request{Name: req.GetFirstName().GetValue()}); err != nil {
    if res, err := h.greeterSrvClient.Hello(ctx, &greeterPB.HelloRequest{Name: req.GetFirstName().GetValue()}); err != nil {
        log.Error().Err(err).Msg("Received greeterService.Hello request error")
        return myErrors.AppError(myErrors.PSE, err)
    } else {
        log.Info().Msgf("Got greeterService responce %s", res.Msg)
    }

    return nil
}

func (h *userHandler) Update(ctx context.Context, req *userPB.UpdateRequest, rsp *userPB.UpdateResponse) error {
    log.Info().Msg("Received UserHandler.Update request")
    // Identify the user
    acc, ok := auth.AccountFromContext(ctx)
    if !ok {
        return errors.Unauthorized("mkit.service.account.user.update", "A valid auth token is required")
    }
    log.Info().Msgf("Caller Account: %v", acc)

    id := req.Id.GetValue()
    if id == "" {
        return myErrors.ValidationError("mkit.service.account.user.update", "validation error: Missing Id")
    }

    model := account_entities.UserORM{}
    username := req.Username.GetValue()
    model.Username = &username
    model.FirstName = req.FirstName.GetValue()
    model.LastName = req.LastName.GetValue()
    model.Email = req.Email.GetValue()

    if err := h.userRepository.Update(id, &model); err != nil {
        return myErrors.AppError(myErrors.DBE, err)
    }

    return nil
}

func (h *userHandler) Delete(ctx context.Context, req *userPB.DeleteRequest, rsp *userPB.DeleteResponse) error {
    log.Info().Msg("Received UserHandler.Delete request")

    id := req.Id.GetValue()
    if id == "" {
        return myErrors.ValidationError("mkit.service.account.user.update", "validation error: Missing Id")
    }

    model := account_entities.UserORM{}
    model.Id = uuid.FromStringOrNil(id)

    if err := h.userRepository.Delete(&model); err != nil {
        return myErrors.AppError(myErrors.DBE, err)
    }

    return nil
}
