package repository

import (
    "github.com/jinzhu/gorm"
    "github.com/pkg/errors"
    "github.com/rs/zerolog/log"
    go_uuid1 "github.com/satori/go.uuid"

    account_entities "github.com/xmlking/micro-starter-kit/service/account/proto/entities"
)

// ProfileRepository interface
type ProfileRepository interface {
    Exist(model *account_entities.ProfileORM) bool
    List(limit, page uint32, sort string, model *account_entities.ProfileORM) (total uint32, profiles []*account_entities.ProfileORM, err error)
    Get(id string) (*account_entities.ProfileORM, error)
    GetByUserID(userId string) (*account_entities.ProfileORM, error)
    Create(model *account_entities.ProfileORM) error
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
func (repo *profileRepository) Exist(model *account_entities.ProfileORM) bool {
    var count int
    userID := model.UserId
    if userID != nil && len(*userID) > 0 {
        repo.db.Model(&account_entities.ProfileORM{}).Where("user_id = ?", *userID).Count(&count)
        if count > 0 {
            return true
        }
    }
    return false
}

// List
func (repo *profileRepository) List(limit, page uint32, sort string, model *account_entities.ProfileORM) (total uint32, profiles []*account_entities.ProfileORM, err error) {
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
    if userID != nil && len(*userID) > 0 {
        db = db.Where("user_id = ?", *userID)
    }
    if model.PreferredTheme != nil && len(*model.PreferredTheme) > 0 {
        db = db.Where("preferred_theme = ?", *model.PreferredTheme)
    }
    if model.Gender != account_entities.Profile_GenderType_name[0] {
        db = db.Where("gender = ?", model.Gender)
    }

    if err = db.Order(sort).Limit(limit).Offset(offset).Find(&profiles).Count(&total).Error; err != nil {
        log.Error().Err(err).Msg("Error in ProfileRepository.List")
        return
    }
    return
}

// Find by ID
func (repo *profileRepository) Get(id string) (profile *account_entities.ProfileORM, err error) {
    println("Get")
    println("id")
    profile = &account_entities.ProfileORM{Id: go_uuid1.FromStringOrNil(id)}

    if err = repo.db.First(profile).Error; err != nil && err != gorm.ErrRecordNotFound {
        log.Error().Err(err).Msg("Error in ProfileRepository.Get")
    }
    println(profile.Id.String())
    println(profile.UserId.String())
    return
}

// Find by UserID
func (repo *profileRepository) GetByUserID(userId string) (profile *account_entities.ProfileORM, err error) {
    println("GetByUserID")
    println("userId")
    user_uuid := go_uuid1.FromStringOrNil(userId)
    profile = &account_entities.ProfileORM{UserId: &user_uuid}
    if err = repo.db.Where(&profile).First(&profile).Error; err != nil && err != gorm.ErrRecordNotFound {
        // if err = repo.db.First(profile).Error; err != nil && err != gorm.ErrRecordNotFound {
        log.Error().Err(err).Msg("Error in ProfileRepository.GetByUserID")
    }
    println(profile.Id.String())
    println(profile.UserId.String())
    return
}

// Create
func (repo *profileRepository) Create(model *account_entities.ProfileORM) error {
    if exist := repo.Exist(model); exist {
        return errors.New("profile already exist")
    }

    if err := repo.db.Create(model).Error; err != nil {
        log.Error().Err(err).Msg("Error in ProfileRepository.Create")
        return err
    }
    return nil
}
