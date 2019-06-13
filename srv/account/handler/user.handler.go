package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"

	"github.com/xmlking/micro-starter-kit/srv/account/model"
	pb "github.com/xmlking/micro-starter-kit/srv/account/proto/account"
	"github.com/xmlking/micro-starter-kit/srv/account/repository"
)

// UserHandler struct
type userHandler struct {
	userRepository repository.UserRepository
}

// NewUserHandler returns an instance of `UserServiceHandler`.
func NewUserHandler(repo repository.UserRepository) pb.UserServiceHandler {
	return &userHandler{
		userRepository: repo,
	}
}

func (h *userHandler) Exist(ctx context.Context, req *pb.UserRequest, rsp *pb.UserExistResponse) error {
	log.Log("Received UserHandler.Exist request")
	// if len(strings.TrimSpace(req.Email)) == 0 {
	// 	return fmt.Errorf("invalid email address")
	// }
	exists := h.userRepository.Exist(req)
	log.Logf("user exists? %t", exists)
	rsp.Exists = exists
	if exists {
		rsp.Msg = "User Found"
		rsp.Code = "200"
	} else {
		rsp.Msg = "User Not Found"
		rsp.Code = "404"
	}

	return nil
}

func (h *userHandler) List(ctx context.Context, req *pb.UserListQuery, rsp *pb.UserListResponse) error {
	log.Log("Received UserHandler.List request")
	total, users, err := h.userRepository.List(req)
	if err != nil {
		return err
	}
	rsp.Total = total
	newUsers := make([]*pb.User, len(users))
	for index, user := range users {
		newUsers[index] = userModelToProto(user)
	}
	rsp.Users = newUsers
	rsp.Msg = fmt.Sprintf("%v Total Users Found", total) // "Users Found"
	rsp.Code = "200"
	return nil
}

func (h *userHandler) Get(ctx context.Context, req *pb.UserRequest, rsp *pb.UserResponse) error {
	log.Log("Received UserHandler.Get request")
	if req.Id == 0 {
		return fmt.Errorf("missing Id")
	}

	user, err := h.userRepository.Get(req)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			rsp.User = nil
			rsp.Msg = "User Not Found"
			rsp.Code = "404"
			return nil
		}
		return err
	}

	rsp.User = userModelToProto(user)
	rsp.Msg = "User Found"
	rsp.Code = "200"
	return nil
}

func (h *userHandler) Create(ctx context.Context, req *pb.UserRequest, rsp *pb.UserResponse) error {
	log.Log("Received UserHandler.Create request")
	if err := h.userRepository.Create(req); err != nil {
		return err
	}
	rsp.Msg = "User Created"
	rsp.Code = "200"
	return nil
}

func (h *userHandler) Update(ctx context.Context, req *pb.UserRequest, rsp *pb.UserResponse) error {
	log.Log("Received UserHandler.Update request")
	if req.Id == 0 {
		return fmt.Errorf("missing Id")
	}
	if err := h.userRepository.Update(req); err != nil {
		return err
	}
	rsp.Msg = "User Updated"
	rsp.Code = "200"
	return nil
}

func (h *userHandler) Delete(ctx context.Context, req *pb.UserRequest, rsp *pb.UserResponse) error {
	log.Log("Received UserHandler.Delete request")
	if req.Id == 0 {
		return fmt.Errorf("missing Id")
	}
	if err := h.userRepository.Delete(req); err != nil {
		return err
	}
	rsp.Msg = "User Deleted"
	rsp.Code = "200"
	return nil
}

func userModelToProto(user *model.User) (proto *pb.User) {
	return &pb.User{
		Id:        uint32(user.Model.ID),
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.Model.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.Model.UpdatedAt.Format(time.RFC3339),
	}
}
