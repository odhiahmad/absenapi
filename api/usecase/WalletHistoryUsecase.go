package usecase

import (
	"bri-rece/api/models"
	"bri-rece/api/repository"
)

type IWalletHistoryUsecase interface {
	CreateHistory(history *models.WalletHistory) (*models.WalletHistory, error)
}

type walletHistoryUsecase struct {
	walletHistory repository.IWalletHistoryRepo
}

func (w *walletHistoryUsecase) CreateHistory(history *models.WalletHistory) (*models.WalletHistory, error) {
	return w.walletHistory.CreateHistory(history)
}

func NewWalletHistoryUsecase(walletHistory repository.IWalletHistoryRepo) IWalletHistoryUsecase {
	return &walletHistoryUsecase{
		walletHistory,
	}
}