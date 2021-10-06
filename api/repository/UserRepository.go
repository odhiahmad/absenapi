package repository

import (
	"bri-rece/api/models"
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IUserRepository interface {
	SaveUser(user *models.User) (*models.User, error)
	FindAllUsers() (*[]models.User, error)
	FindUserByID(uid string) (*models.User, error)
	UpdateAUser(user *models.User) (*models.User, error)
	DeleteAUser(uid string) (string, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepositoryImpl{
		db,
	}
}

func (u *userRepositoryImpl) SaveUser(newUser *models.User) (*models.User, error) {
	var err error
	err = u.db.Debug().Create(&newUser).Error
	if err != nil {
		return &models.User{}, err
	}
	return newUser, nil
}

func (u *userRepositoryImpl) FindAllUsers() (*[]models.User, error) {
	var err error
	var users []models.User
	err = u.db.Debug().Model(&models.User{}).Limit(100).Preload(clause.Associations).Find(&models.User{}).Error
	if err != nil {
		return &users, err
	}
	return &users, nil
}

func (u *userRepositoryImpl) FindUserByID(id string) (*models.User, error) {
	var err error
	uid, _ := uuid.FromString(id)
	var user models.User
	err = u.db.Debug().Table("users").Where("id  = ?", uid).Preload(clause.Associations).Find(&user).Error
	if err != nil {
		return &user, err
	}
	if gorm.ErrRecordNotFound == err {
		return &models.User{}, errors.New("User Not Found")
	}
	return &user, err
}

func (u *userRepositoryImpl) UpdateAUser(user *models.User) (*models.User, error) {
	u.db = u.db.Debug().Save(user)
	if err := u.db.Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepositoryImpl) DeleteAUser(uid string) (string, error) {
	u.db = u.db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&models.User{}).Delete(&models.User{})

	if u.db.Error != nil {
		return "", u.db.Error

	}
	return string(rune(u.db.RowsAffected)), nil

}

