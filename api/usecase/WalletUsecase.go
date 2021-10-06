package usecase

import (
	"bri-rece/api/models"
	"bri-rece/api/repository"
)

type IWalletUsecase interface {
	CreateWallet(wallet *models.Wallet) (*models.Wallet, error)
	TopUp(wallet *models.Wallet, id string) (*models.Wallet, error)
	WithDraw(wallet *models.Wallet, id string) (*models.Wallet, error)
	GetWalletById(id string) (*models.Wallet, error)
}

type walletUsecaseRepo struct {
	walletRepo 	repository.IWalletRepository
}

func (w *walletUsecaseRepo) GetWalletById(id string) (*models.Wallet, error) {
	return w.walletRepo.FindById(id)
}

func (w *walletUsecaseRepo) WithDraw(wallet *models.Wallet, id string) (*models.Wallet, error) {
	return w.walletRepo.WithDraw(wallet, id)
}

func (w *walletUsecaseRepo) CreateWallet(wallet *models.Wallet) (*models.Wallet, error) {
	return w.walletRepo.CreateWallet(wallet)
}

func (w *walletUsecaseRepo) TopUp(wallet *models.Wallet, id string) (*models.Wallet, error) {
	return w.walletRepo.UpdateBalance(wallet, id)
}

func NewWalletUsecase(walletRepo repository.IWalletRepository) IWalletUsecase  {
	return &walletUsecaseRepo{
		walletRepo,
	}
}