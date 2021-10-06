package manager

import (
	"bri-rece/api/connect"
	"bri-rece/api/usecase"
)

type ServiceManager interface {
	UserUseCase() usecase.IUserUseCase
	AccountUseCase() usecase.IAccountUsecase
	WalletUseCase() usecase.IWalletUsecase
	WalletHistoryUseCase() usecase.IWalletHistoryUsecase
	OtpUseCase() usecase.IOtpUseCase
}

type serviceManager struct {
	repo RepoManager
}

func (sm *serviceManager) WalletHistoryUseCase() usecase.IWalletHistoryUsecase {
	return usecase.NewWalletHistoryUsecase(sm.repo.WalletHistoryRepo())
}

func (sm *serviceManager) WalletUseCase() usecase.IWalletUsecase {
	return usecase.NewWalletUsecase(sm.repo.WalletRepo())
}

func (sm *serviceManager) AccountUseCase() usecase.IAccountUsecase {
	return usecase.NewAccountUsecase(sm.repo.AccountRepo())
}

func (sm *serviceManager) UserUseCase() usecase.IUserUseCase {
	return usecase.NewUserUseCase(sm.repo.UserRepo(), sm.AccountUseCase(), sm.WalletUseCase(), sm.OtpUseCase())
}

func (sm *serviceManager) OtpUseCase() usecase.IOtpUseCase {
	return usecase.NewOtpUseCase(sm.repo.OtpRepo())
}

func NewServiceManager(connect connect.Connect) ServiceManager {
	return &serviceManager{repo: NewRepoManager(connect)}
}
