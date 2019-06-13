package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"

	"github.com/xmlking/micro-starter-kit/srv/account/model"
	pb "github.com/xmlking/micro-starter-kit/srv/account/proto/account"
)

// UserRepository interface
type UserRepository interface {
	Exist(req *pb.UserRequest) bool
	List(req *pb.UserListQuery) (total uint32, users []*model.User, err error)
	Get(req *pb.UserRequest) (*model.User, error)
	Create(req *pb.UserRequest) error
	Update(req *pb.UserRequest) error
	Delete(req *pb.UserRequest) error
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
func (repo *userRepository) Exist(req *pb.UserRequest) bool {
	var count int
	if req.Username != "" {
		repo.db.Model(&model.User{}).Where("username = ?", req.Username).Count(&count)
		if count > 0 {
			return true
		}
	}
	if req.Id != 0 {
		repo.db.Model(&model.User{}).Where("id = ?", req.Id).Count(&count)
		if count > 0 {
			return true
		}
	}
	if req.Email != "" {
		repo.db.Model(&model.User{}).Where("email = ?", req.Email).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

// List
func (repo *userRepository) List(req *pb.UserListQuery) (total uint32, users []*model.User, err error) {
	db := repo.db

	var limit, offset uint32
	if req.Limit > 0 {
		limit = req.Limit
	} else {
		limit = 10
	}
	if req.Page > 1 {
		offset = (req.Page - 1) * limit
	} else {
		offset = 0
	}

	var sort string
	if req.Sort != "" {
		sort = req.Sort
	} else {
		sort = "created_at desc"
	}

	if req.Username != "" {
		db = db.Where("username like ?", "%"+req.Username+"%")
	}
	if req.LastName != "" {
		db = db.Where("last_name like ?", "%"+req.LastName+"%")
	}
	if req.Email != "" {
		db = db.Where("email like ?", "%"+req.Email+"%")
	}
	if err = db.Order(sort).Limit(limit).Offset(offset).Find(&users).Count(&total).Error; err != nil {
		log.Logf("Error in UserRepository: %v", err)
		return
	}
	return
}

// Find by ID
func (repo *userRepository) Get(req *pb.UserRequest) (user *model.User, err error) {
	// user = &model.User{Model: gorm.Model{ID: uint(req.Id)}}
	user = &model.User{}
	if err = repo.db.First(&user, req.Id).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Logf("Error in UserRepository: %v", err)
	}
	return
}

// Create
func (repo *userRepository) Create(req *pb.UserRequest) error {
	if exist := repo.Exist(req); exist == true {
		return fmt.Errorf("User already exist")
	}
	user := &model.User{
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	if err := repo.db.Create(user).Error; err != nil {
		log.Logf("Error in UserRepository: %v", err)
		return err
	}
	return nil
}

// Update TODO: Translation
func (repo *userRepository) Update(req *pb.UserRequest) error {
	user := &model.User{
		Model: gorm.Model{ID: uint(req.Id)},
	}
	result := repo.db.Model(user).Updates(req)
	if err := result.Error; err != nil {
		log.Logf("Error in UserRepository: %v", err)
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		log.Logf("Error in UserRepository, rowsAffected: %v", rowsAffected)
		return fmt.Errorf("No Records Updated, No match was found")
	}
	return nil
}

// Delete
func (repo *userRepository) Delete(req *pb.UserRequest) error {
	user := &model.User{
		Model: gorm.Model{ID: uint(req.Id)},
	}
	result := repo.db.Delete(user)
	if err := result.Error; err != nil {
		log.Logf("Error in UserRepository: %v", err)
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		log.Logf("Error in UserRepository, rowsAffected: %v", rowsAffected)
		return fmt.Errorf("No Records Deleted, No match was found")
	}
	return nil
}
