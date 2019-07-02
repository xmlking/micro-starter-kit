package handler

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	log "github.com/sirupsen/logrus"

	myErrors "github.com/xmlking/micro-starter-kit/shared/errors"
	"github.com/xmlking/micro-starter-kit/srv/account/entity"
	pb "github.com/xmlking/micro-starter-kit/srv/account/proto/account"
	"github.com/xmlking/micro-starter-kit/srv/account/repository"
	emailerPB "github.com/xmlking/micro-starter-kit/srv/emailer/proto/emailer"
)

// UserHandler struct
type userHandler struct {
	userRepository repository.UserRepository
	Publisher      micro.Publisher
}

// NewUserHandler returns an instance of `UserServiceHandler`.
func NewUserHandler(repo repository.UserRepository, pub micro.Publisher) pb.UserServiceHandler {
	return &userHandler{
		userRepository: repo,
		Publisher:      pub,
	}
}

func (h *userHandler) Exist(ctx context.Context, req *pb.UserRequest, rsp *pb.UserExistResponse) error {
	log.Info("Received UserHandler.Exist request")
	if err := req.Validate(); err != nil {
		return myErrors.ValidationError("go.micro.srv.account.user.exist", "validation error: %v", err)
	}

	model := &entity.User{
		Model:    gorm.Model{ID: uint(req.Id.GetValue())},
		Username: req.Username.GetValue(),
		Email:    req.Email.GetValue(),
	}

	exists := h.userRepository.Exist(model)
	log.Info("user exists? %t", exists)
	rsp.Result = exists
	return nil
}

func (h *userHandler) List(ctx context.Context, req *pb.UserListQuery, rsp *pb.UserListResponse) error {
	log.Info("Received UserHandler.List request")
	if err := req.Validate(); err != nil {
		return myErrors.ValidationError("go.micro.srv.account.user.list", "validation error: %v", err)
	}

	model := &entity.User{
		Username: req.Username.GetValue(),
		LastName: req.LastName.GetValue(),
		Email:    req.Email.GetValue(),
	}

	total, users, err := h.userRepository.List(req.Limit.GetValue(), req.Page.GetValue(), req.Sort.GetValue(), model)
	if err != nil {
		return errors.NotFound("go.micro.srv.account.user.list", "Error %v", err.Error())
	}
	rsp.Total = total
	newUsers := make([]*pb.User, len(users))
	for index, user := range users {
		newUsers[index] = user.ToPB()
	}
	rsp.Results = newUsers
	return nil
}

func (h *userHandler) Get(ctx context.Context, req *pb.UserRequest, rsp *pb.UserResponse) error {
	log.Info("Received UserHandler.Get request")
	if err := req.Validate(); err != nil {
		return myErrors.ValidationError("go.micro.srv.account.user.get", "validation error: %v", err)
	}
	if req.Id.GetValue() == 0 {
		return myErrors.ValidationError("go.micro.srv.account.user.get", "validation error: Missing Id")
	}
	user, err := h.userRepository.Get(req.Id.GetValue())
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			rsp.Result = nil
			return nil
		}
		return err
	}

	rsp.Result = user.ToPB()
	return nil
}

func (h *userHandler) Create(ctx context.Context, req *pb.UserRequest, rsp *pb.UserResponse) error {
	log.Info("Received UserHandler.Create request")
	if err := req.Validate(); err != nil {
		return myErrors.ValidationError("go.micro.srv.account.user.create", "validation error: %v", err)
	}

	model := &entity.User{
		Username:  req.Username.GetValue(),
		FirstName: req.FirstName.GetValue(),
		LastName:  req.LastName.GetValue(),
		Email:     req.Email.GetValue(),
	}

	if err := h.userRepository.Create(model); err != nil {
		return myErrors.ValidationError("go.micro.srv.account.user.create", "Create Failed: %v", err)
	}

	// send email
	if err := h.Publisher.Publish(ctx, &emailerPB.Message{To: req.Email.GetValue(), From: req.Email.GetValue(), Subject: "this is email subject", Body: "this is email body"}); err != nil {
		return err
	}

	return nil
}

func (h *userHandler) Update(ctx context.Context, req *pb.UserRequest, rsp *pb.UserResponse) error {
	log.Info("Received UserHandler.Update request")
	if err := req.Validate(); err != nil {
		return myErrors.ValidationError("go.micro.srv.account.user.update", "validation error: %v", err)
	}
	if req.Id.GetValue() == 0 {
		return myErrors.ValidationError("go.micro.srv.account.user.update", "validation error: Missing Id")
	}

	model := &entity.User{
		Username:  req.Username.GetValue(),
		FirstName: req.FirstName.GetValue(),
		LastName:  req.LastName.GetValue(),
		Email:     req.Email.GetValue(),
	}

	if err := h.userRepository.Update(req.Id.GetValue(), model); err != nil {
		return myErrors.ValidationError("go.micro.srv.account.user.update", "Update Failed: %v", err)
	}

	return nil
}

func (h *userHandler) Delete(ctx context.Context, req *pb.UserRequest, rsp *pb.UserResponse) error {
	log.Info("Received UserHandler.Delete request")
	if err := req.Validate(); err != nil {
		return myErrors.ValidationError("go.micro.srv.account.user.delete", "validation error: %v", err)
	}

	if req.Id.GetValue() == 0 {
		return myErrors.ValidationError("go.micro.srv.account.user.delete", "validation error: Missing Id")
	}

	model := &entity.User{
		Model: gorm.Model{ID: uint(req.Id.GetValue())},
	}

	if err := h.userRepository.Delete(model); err != nil {
		return myErrors.ValidationError("go.micro.srv.account.user.delete", "Delete Failed: %v", err)
	}

	return nil
}
