package handler

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/thoas/go-funk"
	myErrors "github.com/xmlking/micro-starter-kit/shared/errors"
	log "github.com/xmlking/micro-starter-kit/shared/micro/logger"
	account_entities "github.com/xmlking/micro-starter-kit/srv/account/proto/entities"
	userPB "github.com/xmlking/micro-starter-kit/srv/account/proto/user"
	"github.com/xmlking/micro-starter-kit/srv/account/repository"
	emailerPB "github.com/xmlking/micro-starter-kit/srv/emailer/proto/emailer"
	greeterPB "github.com/xmlking/micro-starter-kit/srv/greeter/proto/greeter"
)

// UserHandler struct
type userHandler struct {
	userRepository   repository.UserRepository
	Publisher        micro.Publisher
	greeterSrvClient greeterPB.GreeterService
}

// NewUserHandler returns an instance of `UserServiceHandler`.
func NewUserHandler(repo repository.UserRepository, pub micro.Publisher, greeterClient greeterPB.GreeterService) userPB.UserServiceHandler {
	return &userHandler{
		userRepository:   repo,
		Publisher:        pub,
		greeterSrvClient: greeterClient,
	}
}

func (h *userHandler) Exist(ctx context.Context, req *userPB.ExistRequest, rsp *userPB.ExistResponse) error {
	log.Info("Received UserHandler.Exist request")
	model := account_entities.UserORM{}
	model.Id = uuid.FromStringOrNil(req.Id.GetValue())
	username := req.Username.GetValue()
	model.Username = &username
	model.Email = req.Email.GetValue()

	exists := h.userRepository.Exist(&model)
	log.Infof("user exists? %t", exists)
	rsp.Result = exists
	return nil
}

func (h *userHandler) List(ctx context.Context, req *userPB.ListRequest, rsp *userPB.ListResponse) error {
	log.Info("Received UserHandler.List request")
	model := account_entities.UserORM{}
	username := req.Username.GetValue()
	model.Username = &username
	model.FirstName = req.FirstName.GetValue()
	model.LastName = req.LastName.GetValue()
	model.Email = req.Email.GetValue()

	total, users, err := h.userRepository.List(req.Limit.GetValue(), req.Page.GetValue(), req.Sort.GetValue(), &model)
	if err != nil {
		return errors.NotFound("account-srv.user.list", "Error %v", err.Error())
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
	log.Info("Received UserHandler.Get request")
	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("account-srv.user.get", "validation error: Missing Id")
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
	log.Info("Received UserHandler.Create request")

	model := account_entities.UserORM{}
	username := req.Username.GetValue()
	model.Username = &username
	model.FirstName = req.FirstName.GetValue()
	model.LastName = req.LastName.GetValue()
	model.Email = req.Email.GetValue()

	if err := h.userRepository.Create(&model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	// send email
	if err := h.Publisher.Publish(ctx, &emailerPB.Message{To: req.Email.GetValue()}); err != nil {
		log.WithError(err, "Received Publisher.Publish request error")
		return myErrors.AppError(myErrors.PSE, err)
	}

	// call greeter
	// if res, err := h.greeterSrvClient.Hello(ctx, &greeterPB.Request{Name: req.GetFirstName().GetValue()}); err != nil {
	if res, err := h.greeterSrvClient.Hello(ctx, &greeterPB.HelloRequest{Name: req.GetFirstName().GetValue()}); err != nil {
		log.WithError(err, "Received greeterService.Hello request error")
		return myErrors.AppError(myErrors.PSE, err)
	} else {
		log.Infof("Got greeterService responce %s", res.Msg)
	}

	return nil
}

func (h *userHandler) Update(ctx context.Context, req *userPB.UpdateRequest, rsp *userPB.UpdateResponse) error {
	log.Info("Received UserHandler.Update request")

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("account-srv.user.update", "validation error: Missing Id")
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
	log.Info("Received UserHandler.Delete request")

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("account-srv.user.update", "validation error: Missing Id")
	}

	model := account_entities.UserORM{}
	model.Id = uuid.FromStringOrNil(id)

	if err := h.userRepository.Delete(&model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	return nil
}
