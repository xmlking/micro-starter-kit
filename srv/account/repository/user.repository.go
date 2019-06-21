package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"

	"github.com/xmlking/micro-starter-kit/srv/account/entity"
)

// UserRepository interface
type UserRepository interface {
	Exist(model *entity.User) bool
	List(limit, page uint32, sort string, model *entity.User) (total uint32, users []*entity.User, err error)
	Get(id uint32) (*entity.User, error)
	Create(model *entity.User) error
	Update(id uint32, model *entity.User) error
	Delete(model *entity.User) error
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
func (repo *userRepository) Exist(model *entity.User) bool {
	var count int
	if model.Username != "" {
		repo.db.Model(&entity.User{}).Where("username = ?", model.Username).Count(&count)
		if count > 0 {
			return true
		}
	}
	if model.ID != 0 {
		repo.db.Model(&entity.User{}).Where("id = ?", model.ID).Count(&count)
		if count > 0 {
			return true
		}
	}
	if model.Email != "" {
		repo.db.Model(&entity.User{}).Where("email = ?", model.Email).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

// List
func (repo *userRepository) List(limit, page uint32, sort string, model *entity.User) (total uint32, users []*entity.User, err error) {
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

	if model.Username != "" {
		db = db.Where("username like ?", "%"+model.Username+"%")
	}
	if model.LastName != "" {
		db = db.Where("last_name like ?", "%"+model.LastName+"%")
	}
	if model.Email != "" {
		db = db.Where("email like ?", "%"+model.Email+"%")
	}
	if err = db.Order(sort).Limit(limit).Offset(offset).Find(&users).Count(&total).Error; err != nil {
		log.Logf("Error in UserRepository: %v", err)
		return
	}
	return
}

// Find by ID
func (repo *userRepository) Get(id uint32) (user *entity.User, err error) {
	// user = &entity.User{Model: gorm.Model{ID: uint(req.Id)}}
	user = &entity.User{}
	if err = repo.db.First(&user, id).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Logf("Error in UserRepository: %v", err)
	}
	return
}

// Create
func (repo *userRepository) Create(model *entity.User) error {
	if exist := repo.Exist(model); exist == true {
		return errors.New("User already exist")
	}
	if err := repo.db.Create(model).Error; err != nil {
		log.Logf("Error in UserRepository: %v", err)
		return err
	}
	return nil
}

// Update TODO: Translation
func (repo *userRepository) Update(id uint32, model *entity.User) error {
	user := &entity.User{
		Model: gorm.Model{ID: uint(id)},
	}
	result := repo.db.Model(user).Updates(model)
	if err := result.Error; err != nil {
		log.Logf("Error in UserRepository: %v", err)
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		log.Logf("Error in UserRepository, rowsAffected: %v", rowsAffected)
		return errors.New("No Records Updated, No match was found")
	}
	return nil
}

// Delete
func (repo *userRepository) Delete(model *entity.User) error {
	result := repo.db.Delete(model)
	if err := result.Error; err != nil {
		log.Logf("Error in UserRepository: %v", err)
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		log.Logf("Error in UserRepository, rowsAffected: %v", rowsAffected)
		return errors.New("No Records Deleted, No match was found")
	}
	return nil
}
