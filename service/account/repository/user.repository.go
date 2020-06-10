package repository

import (
    "github.com/jinzhu/gorm"
    "github.com/pkg/errors"
    "github.com/rs/zerolog/log"
    uuid "github.com/satori/go.uuid"

    account_entities "github.com/xmlking/micro-starter-kit/service/account/proto/entities"
)

// UserRepository interface
type UserRepository interface {
    Exist(model *account_entities.UserORM) bool
    List(limit, page uint32, sort string, model *account_entities.UserORM) (total uint32, users []*account_entities.UserORM, err error)
    Get(id string) (*account_entities.UserORM, error)
    Create(model *account_entities.UserORM) error
    Update(id string, model *account_entities.UserORM) error
    Delete(model *account_entities.UserORM) error
}

// userRepository struct
type userRepository struct {
    db *gorm.DB
}

// NewUserRepository returns an instance of `UserRepository`.
func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{
        db: db,
    }
}

// Exist
func (repo *userRepository) Exist(model *account_entities.UserORM) bool {
    log.Info().Msgf("Received userRepository.Exist request %v", *model)
    var count int
    if model.Username != nil && len(*model.Username) > 0 {
        repo.db.Model(&account_entities.UserORM{}).Where("username = ?", model.Username).Count(&count)
        if count > 0 {
            return true
        }
    }
    if len(model.Id.String()) > 0 {
        repo.db.Model(&account_entities.UserORM{}).Where("id = ?", model.Id.String()).Count(&count)
        if count > 0 {
            return true
        }
    }
    if model.Email != "" {
        repo.db.Model(&account_entities.UserORM{}).Where("email = ?", model.Email).Count(&count)
        if count > 0 {
            return true
        }
    }
    return false
}

// List
func (repo *userRepository) List(limit, page uint32, sort string, model *account_entities.UserORM) (total uint32, users []*account_entities.UserORM, err error) {
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

    if model.Username != nil && len(*model.Username) > 0 {
        db = db.Where("username like ?", "%"+*model.Username+"%")
    }
    if model.FirstName != "" {
        db = db.Where("first_name like ?", "%"+model.FirstName+"%")
    }
    if model.LastName != "" {
        db = db.Where("last_name like ?", "%"+model.LastName+"%")
    }
    if model.Email != "" {
        db = db.Where("email like ?", "%"+model.Email+"%")
    }
    // enable auto preloading for `Profile`
    if err = db.Set("gorm:auto_preload", true).Order(sort).Limit(limit).Offset(offset).Find(&users).Count(&total).Error; err != nil {
        log.Error().Err(err).Msg("Error in UserRepository.List")
        return
    }
    return
}

// Find by ID
func (repo *userRepository) Get(id string) (user *account_entities.UserORM, err error) {
    u2, err := uuid.FromString(id)
    if err != nil {
        return
    }
    user = &account_entities.UserORM{Id: u2}
    // enable auto preloading for `Profile`
    if err = repo.db.Set("gorm:auto_preload", true).First(user).Error; err != nil && err != gorm.ErrRecordNotFound {
        log.Error().Err(err).Msg("Error in UserRepository.Get")
    }
    return
}

// Create
func (repo *userRepository) Create(model *account_entities.UserORM) error {
    if exist := repo.Exist(model); exist {
        return errors.New("user already exist")
    }
    // if err := repo.db.Set("gorm:association_autoupdate", false).Create(model).Error; err != nil {
    if err := repo.db.Create(model).Error; err != nil {
        log.Error().Err(err).Msg("Error in UserRepository.Create")
        return err
    }
    return nil
}

// Update TODO: Translation
func (repo *userRepository) Update(id string, model *account_entities.UserORM) error {
    u2, err := uuid.FromString(id)
    if err != nil {
        return err
    }
    user := &account_entities.UserORM{
        Id: u2,
    }
    // result := repo.db.Set("gorm:association_autoupdate", false).Save(model)
    result := repo.db.Model(user).Updates(model)
    if err := result.Error; err != nil {
        log.Error().Err(err).Msg("Error in UserRepository.Update")
        return err
    }
    if rowsAffected := result.RowsAffected; rowsAffected == 0 {
        log.Error().Msgf("Error in UserRepository.Update, rowsAffected: %v", rowsAffected)
        return errors.New("no records updated, No match was found")
    }
    return nil
}

// Delete
func (repo *userRepository) Delete(model *account_entities.UserORM) error {
    result := repo.db.Delete(model)
    if err := result.Error; err != nil {
        log.Error().Err(err).Msg("Error in UserRepository.Delete")
        return err
    }
    if rowsAffected := result.RowsAffected; rowsAffected == 0 {
        log.Error().Msgf("Error in UserRepository.Delete, rowsAffected: %v", rowsAffected)
        return errors.New("no records deleted, No match was found")
    }
    return nil
}
