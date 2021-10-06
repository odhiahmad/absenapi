package usecase

import (
	"bri-rece/api/models"
	"bri-rece/api/models/dto"
	"bri-rece/api/repository"
	"fmt"
)

type IAccountUsecase interface {
	SaveAccount(account *models.Account) (*models.Account, error)
	LoginByUsername(dtoLogin *dto.Login) (string, error)
	UnActiveAccount(id string) (*models.Account, error)
	ActivatedAccount(id string) (*models.Account, error)
	FindByIdAccount(id string) (*models.Account, error)
}

type accountUsecaseRepo struct {
	accountRepo repository.IAccountRepository
}

func (a *accountUsecaseRepo) ActivatedAccount(id string) (*models.Account, error) {
	return a.accountRepo.ActivatedAccount(id)
}

func (a *accountUsecaseRepo) FindByIdAccount(id string) (*models.Account, error) {
	fmt.Println("account id", id)
	return a.accountRepo.FindByIdAccount(id)
}

func (a *accountUsecaseRepo) UnActiveAccount(id string) (*models.Account, error) {
	return a.accountRepo.UnActiveAccount(id)
}

func NewAccountUsecase(accountRepo repository.IAccountRepository) IAccountUsecase {
	return &accountUsecaseRepo{
		accountRepo,
	}
}

func (a *accountUsecaseRepo) SaveAccount(account *models.Account) (*models.Account, error) {
	return a.accountRepo.SaveAccount(account)
}

func (a *accountUsecaseRepo) LoginByUsername(dtoLogin *dto.Login) (string, error) {
	return a.accountRepo.LoginByUsername(dtoLogin)
}
