package repository

import (
	"bri-rece/api/models"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type IOtpRepository interface {
	SaveOtp(otp *models.Otp) (*models.Otp, error)
	FindOtpByUserID(userId string, otpCode string) (*models.Otp, error)
	UpdateOtp(otp *models.Otp, uid string) (error, *models.Otp)
}

type otpRepositoryImpl struct {
	db *gorm.DB
}

func NewOtpRepository(db *gorm.DB) IOtpRepository {
	return &otpRepositoryImpl{
		db,
	}
}

func (u *otpRepositoryImpl) SaveOtp(newOtp *models.Otp) (*models.Otp, error) {
	var err error
	err = u.db.Debug().Create(&newOtp).Error
	if err != nil {
		return &models.Otp{}, err
	}
	return newOtp, nil
}

func (u *otpRepositoryImpl) FindOtpByUserID(userId string, otpCode string) (*models.Otp, error) {
	var err error
	var otp models.Otp
	uid, _ := uuid.FromString(userId)

	if err = u.db.Debug().Table("otps").Where("user_id = ?", uid).Where("otp_code = ?", otpCode).Find(&otp).Error; err != nil {
		return nil, err
	}

	return &otp, nil
}

func (u *otpRepositoryImpl) UpdateOtp(user *models.Otp, uid string) (error, *models.Otp) {

	return nil, nil
}
