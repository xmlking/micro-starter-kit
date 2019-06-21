package handler

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/util/log"

	myErrors "github.com/xmlking/micro-starter-kit/shared/errors"
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
	log.Log("Received UserHandler.Exist request")
	if err := req.Validate(); err != nil {
		return myErrors.ValidationError("go.micro.srv.account.user.exist", "validation error: %v", err)
	}
	// if len(strings.TrimSpace(req.Email)) == 0 {
	// 	return fmt.Errorf("invalid email address")
	// }
	exists := h.userRepository.Exist(req)
	log.Logf("user exists? %t", exists)
	rsp.Result = exists
	return nil
}

func (h *userHandler) List(ctx context.Context, req *pb.UserListQuery, rsp *pb.UserListResponse) error {
	log.Log("Received UserHandler.List request")
	if err := req.Validate(); err != nil {
		return myErrors.ValidationError("go.micro.srv.account.user.list", "validation error: %v", err)
	}
	total, users, err := h.userRepository.List(req)
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
	log.Log("Received UserHandler.Get request")
	if err := req.Validate(); err != nil {
		return myErrors.ValidationError("go.micro.srv.account.user.get", "validation error: %v", err)
	}
	if req.Id == 0 {
		return fmt.Errorf("missing Id")
	}

	user, err := h.userRepository.Get(req)
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
	log.Log("Received UserHandler.Create request")
	if err := req.Validate(); err != nil {
		return myErrors.ValidationError("go.micro.srv.account.user.create", "validation error: %v", err)
	}
	if err := h.userRepository.Create(req); err != nil {
		return err
	}

	// send email
	if err := h.Publisher.Publish(ctx, &emailerPB.Message{To: req.Email, From: req.Email, Subject: "this is email subject", Body: "this is email body"}); err != nil {
		return err
	}

	return nil
}

func (h *userHandler) Update(ctx context.Context, req *pb.UserRequest, rsp *pb.UserResponse) error {
	log.Log("Received UserHandler.Update request")
	if err := req.Validate(); err != nil {
		return myErrors.ValidationError("go.micro.srv.account.user.update", "validation error: %v", err)
	}
	if req.Id == 0 {
		return fmt.Errorf("missing Id")
	}
	if err := h.userRepository.Update(req); err != nil {
		return err
	}

	return nil
}

func (h *userHandler) Delete(ctx context.Context, req *pb.UserRequest, rsp *pb.UserResponse) error {
	log.Log("Received UserHandler.Delete request")
	if err := req.Validate(); err != nil {
		return myErrors.ValidationError("go.micro.srv.account.user.delete", "validation error: %v", err)
	}
	if req.Id == 0 {
		return fmt.Errorf("missing Id")
	}
	if err := h.userRepository.Delete(req); err != nil {
		return err
	}

	return nil
}
