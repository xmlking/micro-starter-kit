package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	userPB "github.com/xmlking/micro-starter-kit/srv/account/proto/user"
)

// ProfileRepository interface
type ProfileRepository interface {
	Exist(model *userPB.ProfileORM) bool
	List(limit, page uint32, sort string, model *userPB.ProfileORM) (total uint32, profiles []*userPB.ProfileORM, err error)
	Get(id string) (*userPB.ProfileORM, error)
	Create(model *userPB.ProfileORM) error
}

// profileRepository struct
type profileRepository struct {
	db *gorm.DB
}

// NewProfileRepository returns an instance of `ProfileRepository`.
func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{
		db: db,
	}
}

// Exist
func (repo *profileRepository) Exist(model *userPB.ProfileORM) bool {
	var count int
	userID := model.UserId
	if userID != nil && *userID != "" {
		repo.db.Model(&userPB.ProfileORM{}).Where("user_id = ?", *userID).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

// List
func (repo *profileRepository) List(limit, page uint32, sort string, model *userPB.ProfileORM) (total uint32, profiles []*userPB.ProfileORM, err error) {
	db := repo.db

	if limit == 0 {
		limit = 10
	}
	var offset uint32
	if page > 1 {
		offset = (page - 1) * limit
	} else {
		offset = 0
	}
	if sort == "" {
		sort = "created_at desc"
	}

	userID := model.UserId
	if userID != nil && *userID != "" {
		db = db.Where("user_id = ?", *userID)
	}
	if model.Gender != "" {
		db = db.Where("gender = ?", model.Gender)
	}

	if err = db.Order(sort).Limit(limit).Offset(offset).Find(&profiles).Count(&total).Error; err != nil {
		log.WithError(err).Error("Error in ProfileRepository.List")
		return
	}
	return
}

// Find by ID
func (repo *profileRepository) Get(id string) (profile *userPB.ProfileORM, err error) {
	profile = &userPB.ProfileORM{Id: id}
	if err = repo.db.First(profile).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.WithError(err).Error("Error in ProfileRepository.Get")
	}
	return
}

// Create
func (repo *profileRepository) Create(model *userPB.ProfileORM) error {
	if exist := repo.Exist(model); exist {
		return fmt.Errorf("profile already exist")
	}

	if err := repo.db.Create(model).Error; err != nil {
		log.WithError(err).Error("Error in ProfileRepository.Create")
		return err
	}
	return nil
}
