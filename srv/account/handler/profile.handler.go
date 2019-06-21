package handler

import (
	"context"
	"fmt"

	// "github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"

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
	total, profiles, err := h.profileRepository.List(req)
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
	if req.Id.GetValue() == 0 {
		return fmt.Errorf("missing Id")
	}

	profile, err := h.profileRepository.Get(req)
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
	if err := h.profileRepository.Create(req); err != nil {
		return err
	}
	return nil
}
