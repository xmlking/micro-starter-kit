package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	pb "github.com/xmlking/micro-starter-kit/srv/account/proto/account"
)

// UserRepository interface
type UserRepository interface {
	Exist(model *pb.UserORM) bool
	List(limit, page uint32, sort string, model *pb.UserORM) (total uint32, users []*pb.UserORM, err error)
	Get(id string) (*pb.UserORM, error)
	Create(model *pb.UserORM) error
	Update(id string, model *pb.UserORM) error
	Delete(model *pb.UserORM) error
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
func (repo *userRepository) Exist(model *pb.UserORM) bool {
	log.Info("Received userRepository.Exist request %v", *model)
	var count int
	if model.Username != "" {
		repo.db.Model(&pb.UserORM{}).Where("username = ?", model.Username).Count(&count)
		if count > 0 {
			return true
		}
	}
	if model.Id != "" {
		repo.db.Model(&pb.UserORM{}).Where("id = ?", model.Id).Count(&count)
		if count > 0 {
			return true
		}
	}
	if model.Email != "" {
		repo.db.Model(&pb.UserORM{}).Where("email = ?", model.Email).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

// List
func (repo *userRepository) List(limit, page uint32, sort string, model *pb.UserORM) (total uint32, users []*pb.UserORM, err error) {
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
	// enable auto preloading for `Profile`
	if err = db.Set("gorm:auto_preload", true).Order(sort).Limit(limit).Offset(offset).Find(&users).Count(&total).Error; err != nil {
		log.WithError(err).Error("Error in UserRepository.List")
		return
	}
	return
}

// Find by ID
func (repo *userRepository) Get(id string) (user *pb.UserORM, err error) {
	user = &pb.UserORM{Id: id}
	// enable auto preloading for `Profile`
	if err = repo.db.Set("gorm:auto_preload", true).First(user).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.WithError(err).Error("Error in UserRepository.Get")
	}
	return
}

// Create
func (repo *userRepository) Create(model *pb.UserORM) error {
	if exist := repo.Exist(model); exist == true {
		return errors.New("User already exist")
	}
	// if err := repo.db.Set("gorm:association_autoupdate", false).Create(model).Error; err != nil {
	if err := repo.db.Create(model).Error; err != nil {
		log.WithError(err).Error("Error in UserRepository.Create")
		return err
	}
	return nil
}

// Update TODO: Translation
func (repo *userRepository) Update(id string, model *pb.UserORM) error {
	user := &pb.UserORM{
		Id: id,
	}
	// result := repo.db.Set("gorm:association_autoupdate", false).Save(model)
	result := repo.db.Model(user).Updates(model)
	if err := result.Error; err != nil {
		log.WithError(err).Error("Error in UserRepository.Update")
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		log.Errorf("Error in UserRepository.Update, rowsAffected: %v", rowsAffected)
		return errors.New("No Records Updated, No match was found")
	}
	return nil
}

// Delete
func (repo *userRepository) Delete(model *pb.UserORM) error {
	result := repo.db.Delete(model)
	if err := result.Error; err != nil {
		log.WithError(err).Error("Error in UserRepository.Delete")
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		log.Errorf("Error in UserRepository.Delete, rowsAffected: %v", rowsAffected)
		return errors.New("No Records Deleted, No match was found")
	}
	return nil
}
