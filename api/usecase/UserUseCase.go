package usecase

import (
	"bri-rece/api/models"
	"bri-rece/api/repository"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net/smtp"
	"strings"
)

const (
	CONFIG_SMTP_PORT     = 587
	CONFIG_SMTP_HOST     = "smtp.gmail.com"
	CONFIG_SENDER_NAME   = "RECE Official <company@gmail.com>"
	CONFIG_AUTH_EMAIL    = "odhiahmad15@gmail.com"
	CONFIG_AUTH_PASSWORD = "@Geforce920m"
)

type IUserUseCase interface {
	Register(newUser *models.UserRequest) (*models.UserResponse, error)
	GetUserInfo(id string) *models.User
	UpdateInfo(editUser *models.User) (*models.User, error)
	Unregister(id string) (string, error)
	FindUserById(id string) (*models.User, error)
}

type UserUseCaseRepo struct {
	userRepo      repository.IUserRepository
	accountUsace  IAccountUsecase
	walletUsecase IWalletUsecase
	otpUsecase    IOtpUseCase
}

func (u *UserUseCaseRepo) FindUserById(id string) (*models.User, error) {
	return u.userRepo.FindUserByID(id)
}

func (u *UserUseCaseRepo) Register(newUser *models.UserRequest) (*models.UserResponse, error) {
	otpCode := EncodeToString(6)

	user := models.User{
		Name:        newUser.Name,
		PhoneNumber: newUser.PhoneNumber,
		Email:       newUser.Email,
		BirthPlace:  newUser.BirthPlace,
		BirthDate:   newUser.BirthDate,
		MotherName:  newUser.MotherName,
		BankAccount: newUser.BankAccount,
	}

	user.Prepare()
	userData, err := u.userRepo.SaveUser(&user)
	if err != nil {
		return nil, err
	}

	account := models.Account{
		Username: newUser.Username,
		Password: newUser.Password,
		UserID:   userData.ID,
	}

	account.Prepare()

	accountData, err := u.accountUsace.SaveAccount(&account)
	if err != nil {
		return nil, err
	}

	wallet := models.Wallet{
		Balance:   0,
		AccountID: accountData.ID,
	}
	u.walletUsecase.CreateWallet(&wallet)

	otp := models.Otp{
		OtpCode: otpCode,
		UserID:  userData.ID,
	}

	u.otpUsecase.CreateOtp(&otp)

	//Sending OTP to customer's mail
	to := []string{user.Email}
	cc := []string{"odhiahmad15@gmail.com"}
	subject := "RECE OTP Message"
	message := "Jangan tunjukan kode OTP ini ke siapapun. " + "\n" + "kode: " + otpCode
	err = sendEmail(to, cc, subject, message)
	if err != nil {
		log.Fatal(err.Error())
	}

	// u.walletUsecase.CreateWallet(&wallet)

	return &models.UserResponse{
		ID:          userData.ID,
		Username:    accountData.Username,
		Name:        userData.Name,
		PhoneNumber: userData.PhoneNumber,
		Email:       userData.Email,
		BirthPlace:  userData.BirthPlace,
		BirthDate:   userData.BirthDate,
		MotherName:  userData.MotherName,
		BankAccount: userData.BankAccount,
		CreatedAt:   userData.CreatedAt,
		UpdatedAt:   userData.UpdatedAt,
	}, nil
}

func sendEmail(to []string, cc []string, subject, message string) error {

	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}

func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func (u *UserUseCaseRepo) GetUserInfo(id string) *models.User {
	user, err := u.userRepo.FindUserByID(id)
	if err != nil {
		return nil
	}
	return user
}

func (u *UserUseCaseRepo) Unregister(id string) (string, error) {
	return u.userRepo.DeleteAUser(id)
}

func (u *UserUseCaseRepo) UpdateInfo(editUser *models.User) (*models.User, error) {
	return u.userRepo.UpdateAUser(editUser)
}

func NewUserUseCase(userRepo repository.IUserRepository, service IAccountUsecase, serviceWallet IWalletUsecase, serviceOtp IOtpUseCase) IUserUseCase {
	return &UserUseCaseRepo{
		userRepo,
		service,
		serviceWallet,
		serviceOtp,
	}
}
