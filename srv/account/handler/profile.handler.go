package handler

import (
	"context"

	// "github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"

	myErrors "github.com/xmlking/micro-starter-kit/shared/errors"
	"github.com/xmlking/micro-starter-kit/srv/account/entity"
	pb "github.com/xmlking/micro-starter-kit/srv/account/proto/account"
	"github.com/xmlking/micro-starter-kit/srv/account/repository"
)

// ProfileHandler struct
type profileHandler struct {
	profileRepository repository.ProfileRepository
}

// NewProfileHandler returns an instance of `ProfileServiceHandler`.
func NewProfileHandler(repo repository.ProfileRepository) pb.ProfileServiceHandler {
	return &profileHandler{
		profileRepository: repo,
	}
}

func (h *profileHandler) List(ctx context.Context, req *pb.ProfileListQuery, rsp *pb.ProfileListResponse) error {
	log.Log("Received ProfileHandler.List request")
	if err := req.Validate(); err != nil {
		return myErrors.ValidationError("go.micro.srv.account.profile.list", "validation error: %v", err)
	}
	model := &entity.Profile{
		UserID: req.UserId.GetValue(),
		Gender: req.Gender.GetValue(),
	}
	total, profiles, err := h.profileRepository.List(req.Limit.GetValue(), req.Page.GetValue(), req.Sort.GetValue(), model)
	if err != nil {
		return err
	}
	rsp.Total = total
	newProfiles := make([]*pb.Profile, len(profiles))
	for index, profile := range profiles {
		newProfiles[index] = profile.ToPB()
	}
	rsp.Results = newProfiles
	return nil
}

func (h *profileHandler) Get(ctx context.Context, req *pb.ProfileRequest, rsp *pb.ProfileResponse) error {
	log.Log("Received ProfileHandler.Get request")
	if err := req.Validate(); err != nil {
		return myErrors.ValidationError("go.micro.srv.account.profile.get", "validation error: %v", err)
	}
	if req.Id.GetValue() == 0 {
		return myErrors.ValidationError("go.micro.srv.account.profile.get", "validation error: Missing Id")
	}
	profile, err := h.profileRepository.Get(req.Id.GetValue())
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			rsp.Result = nil
			return nil
		}
		return err
	}

	rsp.Result = profile.ToPB()
	return nil
}

func (h *profileHandler) Create(ctx context.Context, req *pb.ProfileRequest, rsp *pb.ProfileResponse) error {
	log.Log("Received ProfileHandler.Create request")
	if err := req.Validate(); err != nil {
		return myErrors.ValidationError("go.micro.srv.account.profile.rceate", "validation error: %v", err)
	}
	model := &entity.Profile{
		UserID: req.UserId.GetValue(),
		TZ:     req.Tz.GetValue(),
		Gender: req.Gender.GetValue(),
		Avatar: req.Avatar.GetValue(),
	}

	if err := h.profileRepository.Create(model); err != nil {
		return err
	}
	return nil
}
