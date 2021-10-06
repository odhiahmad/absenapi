package usecase

import (
	"bri-rece/api/models"
	"bri-rece/api/repository"
)

type IOtpUseCase interface {
	CreateOtp(newOtp *models.Otp) (*models.Otp, error)
	Verification(userId string, otpCode string) (*models.Otp, error)
	// UpdateOtp(id string, newOtp *models.Otp) (error, *models.Otp)
}

type otpUseCaseRepo struct {
	otpRepo repository.IOtpRepository
}

func (u *otpUseCaseRepo) CreateOtp(newOtp *models.Otp) (*models.Otp, error) {
	return u.otpRepo.SaveOtp(newOtp)
}

func (u *otpUseCaseRepo) Verification(userId string, otpCode string) (*models.Otp, error) {
	otp, err := u.otpRepo.FindOtpByUserID(userId, otpCode)
	if err != nil {
		return nil, err
	}

	return otp, nil
}

func NewOtpUseCase(otpRepo repository.IOtpRepository) IOtpUseCase {
	return &otpUseCaseRepo{
		otpRepo,
	}
}
