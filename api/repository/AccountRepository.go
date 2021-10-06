package repository

import (
	"bri-rece/api/middlewares"
	"bri-rece/api/models"
	"bri-rece/api/models/dto"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IAccountRepository interface {
	SaveAccount(account *models.Account) (*models.Account, error)
	LoginByUsername(dtoLogin *dto.Login) (string, error)
	UnActiveAccount(id string) (*models.Account, error)
	ActivatedAccount(id string) (*models.Account, error)
	FindByIdAccount(id string) (*models.Account, error)
}

type accountRepository struct {
	db *gorm.DB
}

func (a *accountRepository) FindByIdAccount(id string) (*models.Account, error) {
	uid, _ := uuid.FromString(id)
	accountDb := models.Account{}
	if err := a.db.Debug().Table("accounts").Where("id = ?", uid).Preload(clause.Associations).Find(&accountDb).Error; err != nil {
		return nil, err
	}
	return &accountDb, nil
}

func (a *accountRepository) UnActiveAccount(id string) (*models.Account, error) {
	uid, _ := uuid.FromString(id)
	var account models.Account
	a.db = a.db.Debug().Model(&account).Where("id = ?", uid).Take(&account).UpdateColumns(
		map[string]interface{}{
			"is_active":  false,
			"deleted_at": time.Now(),
		},
	)
	if err := a.db.Error; err != nil {
		return &account, err
	}
	return &account, nil
}

func (a *accountRepository) ActivatedAccount(id string) (*models.Account, error) {
	uid, _ := uuid.FromString(id)
	var account models.Account
	a.db = a.db.Debug().Model(&account).Where("id = ?", uid).Take(&account).UpdateColumns(
		map[string]interface{}{
			"is_active":  true,
			"deleted_at": time.Now(),
		},
	)
	if err := a.db.Error; err != nil {
		return &account, err
	}
	return &account, nil
}

func NewAccountRepository(db *gorm.DB) IAccountRepository {
	return &accountRepository{
		db,
	}
}



func (a *accountRepository) SaveAccount(account *models.Account) (*models.Account, error) {
	var err error
	err = a.db.Debug().Create(&account).Error
	if err != nil {
		return &models.Account{}, err
	}
	return account, nil
}

func (a *accountRepository) LoginByUsername(dtoLogin *dto.Login) (string, error) {

	var err error
	var accounts models.Account

	if err = a.db.Debug().Table("accounts").Where("username = ?", dtoLogin.Username).Find(&accounts).Error; err != nil {
		return "", err
	}
	err = models.VerifyPassword(accounts.Password, dtoLogin.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return middlewares.CreateToken(accounts.ID)
}
