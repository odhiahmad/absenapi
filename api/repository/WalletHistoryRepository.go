package repository

import (
	"bri-rece/api/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type IWalletHistoryRepo interface {
	CreateHistory(history *models.WalletHistory) (*models.WalletHistory, error)
	GetHistoryById(id string) (*models.WalletHistory, error)
}

type walletHistoryRepo struct {
	db *gorm.DB
}

func (w *walletHistoryRepo) CreateHistory(history *models.WalletHistory) (*models.WalletHistory, error) {
	var err error
	err = w.db.Debug().Create(&history).Error
	if err != nil {
		return &models.WalletHistory{}, err
	}
	return history, nil
}

func (w *walletHistoryRepo) GetHistoryById(id string) (*models.WalletHistory, error) {
	uid, _ := uuid.FromString(id)
	var err error
	history := models.WalletHistory{}
	if err = w.db.Debug().Table("wallet_histories").Where("wallet_id = ?", uid).Find(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}

func NewWalletHistoryRepo(db *gorm.DB) IWalletHistoryRepo  {
	return &walletHistoryRepo{
		db,
	}
}