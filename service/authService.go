package service

import (
	"log"

	"github.com/odhiahmad/absenapi/entity"
	"github.com/odhiahmad/absenapi/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(username string, password string) interface{}
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(username string, password string) interface{} {
	res := service.userRepository.VerifyCredential(username, password)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Username == username && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
