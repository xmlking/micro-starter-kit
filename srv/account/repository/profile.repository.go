package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/xmlking/micro-starter-kit/srv/account/entity"
)

// ProfileRepository interface
type ProfileRepository interface {
	Exist(model *entity.Profile) bool
	List(limit, page uint32, sort string, model *entity.Profile) (total uint32, profiles []*entity.Profile, err error)
	Get(id uint32) (*entity.Profile, error)
	Create(model *entity.Profile) error
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
func (repo *profileRepository) Exist(model *entity.Profile) bool {
	var count int
	if model.UserID != 0 {
		repo.db.Model(&entity.Profile{}).Where("user_id = ?", model.UserID).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

// List
func (repo *profileRepository) List(limit, page uint32, sort string, model *entity.Profile) (total uint32, profiles []*entity.Profile, err error) {
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

	if model.UserID != 0 {
		db = db.Where("user_id = ?", model.UserID)
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
func (repo *profileRepository) Get(id uint32) (profile *entity.Profile, err error) {
	// profile = &entity.Profile{Model: gorm.Model{ID: uint(req.Id)}}
	profile = &entity.Profile{}
	if err = repo.db.First(&profile, id).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.WithError(err).Error("Error in ProfileRepository.Get")
	}
	return
}

// Create
func (repo *profileRepository) Create(model *entity.Profile) error {
	if exist := repo.Exist(model); exist == true {
		return fmt.Errorf("Profile already exist")
	}

	if err := repo.db.Create(model).Error; err != nil {
		log.WithError(err).Error("Error in ProfileRepository.Create")
		return err
	}
	return nil
}
