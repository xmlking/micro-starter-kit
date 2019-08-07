package handler

import (
	"context"
	"time"

	// "github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"

	ptypes1 "github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"
	myErrors "github.com/xmlking/micro-starter-kit/shared/errors"
	pb "github.com/xmlking/micro-starter-kit/srv/account/proto/account"
	"github.com/xmlking/micro-starter-kit/srv/account/repository"
)

// ProfileHandler struct
type profileHandler struct {
	profileRepository repository.ProfileRepository
	contextLogger     log.FieldLogger
}

// NewProfileHandler returns an instance of `ProfileServiceHandler`.
func NewProfileHandler(repo repository.ProfileRepository, logger log.FieldLogger) pb.ProfileServiceHandler {
	return &profileHandler{
		profileRepository: repo,
		contextLogger:     logger,
	}
}

func (ph *profileHandler) List(ctx context.Context, req *pb.ProfileListQuery, rsp *pb.ProfileListResponse) error {
	ph.contextLogger.Info("Received ProfileHandler.List request")
	if err := req.Validate(); err != nil {
		return myErrors.ValidationError("go.micro.srv.account.profile.list", "validation error: %v", err)
	}
	model := pb.ProfileORM{
		Id:     req.UserId.GetValue(),
		Gender: req.Gender.GetValue(),
	}

	total, profiles, err := ph.profileRepository.List(req.Limit.GetValue(), req.Page.GetValue(), req.Sort.GetValue(), &model)
	if err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}
	rsp.Total = total
	newProfiles := make([]*pb.Profile, len(profiles))
	for index, profile := range profiles {
		tempProfile, _ := profile.ToPB(ctx)
		newProfiles[index] = &tempProfile
	}
	rsp.Results = newProfiles
	return nil
}

func (ph *profileHandler) Get(ctx context.Context, req *pb.ProfileRequest, rsp *pb.ProfileResponse) error {
	ph.contextLogger.Info("Received ProfileHandler.Get request")
	if err := req.Validate(); err != nil {
		return myErrors.ValidationError("go.micro.srv.account.profile.get", "validation error: %v", err)
	}
	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("go.micro.srv.account.profile.get", "validation error: Missing Id")
	}

	profile, err := ph.profileRepository.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			rsp.Result = nil
			return nil
		}
		return myErrors.AppError(myErrors.DBE, err)
	}

	tempProfile, _ := profile.ToPB(ctx)
	rsp.Result = &tempProfile
	return nil
}

func (ph *profileHandler) Create(ctx context.Context, req *pb.ProfileRequest, rsp *pb.ProfileResponse) error {
	ph.contextLogger.Debug("Received ProfileHandler.Create request")
	if err := req.Validate(); err != nil {
		return myErrors.ValidationError("go.micro.srv.account.profile.rceate", "validation error: %v", err)
	}
	model := pb.ProfileORM{}
	userID := req.UserId.GetValue()
	model.UserId = &userID
	model.Tz = req.Tz.GetValue()
	model.Gender = req.Gender.GetValue()
	model.Avatar = req.Avatar.GetValue()
	if req.Birthday != nil {
		var t time.Time
		var err error
		if t, err = ptypes1.Timestamp(req.Birthday); err != nil {
			return myErrors.ValidationError("go.micro.srv.account.profile.rceate", "Invalid birthday: %v", err)
		}
		model.Birthday = &t
	}

	if err := ph.profileRepository.Create(&model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}
	return nil
}
