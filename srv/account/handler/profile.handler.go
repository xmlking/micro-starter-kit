package handler

import (
	"context"
	"time"

	// "github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
	"github.com/thoas/go-funk"

	ptypes1 "github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"
	myErrors "github.com/xmlking/micro-starter-kit/shared/errors"
	profilePB "github.com/xmlking/micro-starter-kit/srv/account/proto/profile"
	entityPB "github.com/xmlking/micro-starter-kit/srv/account/proto/user"
	"github.com/xmlking/micro-starter-kit/srv/account/repository"
)

// ProfileHandler struct
type profileHandler struct {
	profileRepository repository.ProfileRepository
	contextLogger     log.FieldLogger
}

// NewProfileHandler returns an instance of `ProfileServiceHandler`.
func NewProfileHandler(repo repository.ProfileRepository, logger log.FieldLogger) profilePB.ProfileServiceHandler {
	return &profileHandler{
		profileRepository: repo,
		contextLogger:     logger,
	}
}

func (ph *profileHandler) List(ctx context.Context, req *profilePB.ListRequest, rsp *profilePB.ListResponse) error {
	ph.contextLogger.Info("Received ProfileHandler.List request")
	model := entityPB.ProfileORM{
		Id:     req.UserId.GetValue(),
		Gender: req.Gender.GetValue(),
	}

	total, profiles, err := ph.profileRepository.List(req.Limit.GetValue(), req.Page.GetValue(), req.Sort.GetValue(), &model)
	if err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}
	rsp.Total = total
	// newProfiles := make([]*pb.Profile, len(profiles))
	// for index, profile := range profiles {
	// 	tempProfile, _ := profile.ToPB(ctx)
	// 	newProfiles[index] = &tempProfile
	// }
	newProfiles := funk.Map(profiles, func(profile *entityPB.ProfileORM) *entityPB.Profile {
		tempProfile, _ := profile.ToPB(ctx)
		return &tempProfile
	}).([]*entityPB.Profile)

	rsp.Results = newProfiles
	return nil
}

func (ph *profileHandler) Get(ctx context.Context, req *profilePB.GetRequest, rsp *profilePB.GetResponse) error {
	ph.contextLogger.Info("Received ProfileHandler.Get request")
	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("account-srv.profile.get", "validation error: Missing Id")
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

func (ph *profileHandler) Create(ctx context.Context, req *profilePB.CreateRequest, rsp *profilePB.CreateResponse) error {
	ph.contextLogger.Debug("Received ProfileHandler.Create request")
	model := entityPB.ProfileORM{}
	userID := req.UserId.GetValue()
	model.UserId = &userID
	model.Tz = req.Tz.GetValue()
	model.Gender = req.Gender.GetValue()
	model.Avatar = req.Avatar.GetValue()
	if req.Birthday != nil {
		var t time.Time
		var err error
		if t, err = ptypes1.Timestamp(req.Birthday); err != nil {
			return myErrors.ValidationError("account-srv.profile.rceate", "Invalid birthday: %v", err)
		}
		model.Birthday = &t
	}

	if err := ph.profileRepository.Create(&model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}
	return nil
}
