package manager

import (
	"bri-rece/api/connect"
	"bri-rece/api/repository"
)

type RepoManager interface {
	//UserAuthRepo() repository.IAuthRepository
	UserRepo() repository.IUserRepository
	AccountRepo() repository.IAccountRepository
	WalletRepo() repository.IWalletRepository
	WalletHistoryRepo() repository.IWalletHistoryRepo
	OtpRepo() repository.IOtpRepository
}

type repoManager struct {
	connect connect.Connect
}

func (rm *repoManager) WalletHistoryRepo() repository.IWalletHistoryRepo {
	return repository.NewWalletHistoryRepo(rm.connect.SqlDb())
}

func (rm *repoManager) WalletRepo() repository.IWalletRepository {
	return repository.NewWalletRepository(rm.connect.SqlDb())
}

func (rm *repoManager) AccountRepo() repository.IAccountRepository {
	return repository.NewAccountRepository(rm.connect.SqlDb())
}

func (rm *repoManager) UserRepo() repository.IUserRepository {
	return repository.NewUserRepository(rm.connect.SqlDb())
}

func (rm *repoManager) OtpRepo() repository.IOtpRepository {
	return repository.NewOtpRepository(rm.connect.SqlDb())
}

func NewRepoManager(connect connect.Connect) RepoManager {
	return &repoManager{connect}
}
