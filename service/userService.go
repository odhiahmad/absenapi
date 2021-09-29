package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/odhiahmad/absenapi/dto"
	"github.com/odhiahmad/absenapi/entity"
	"github.com/odhiahmad/absenapi/repository"
)

type UserService interface {
	CreateUser(user dto.UserCreateDTO) entity.User
	UpdateUser(user dto.UserUpdateDTO) entity.User
	Profile(userId string) entity.User
	FindByUsername(user string) entity.User
	IsDuplicateUsername(user string) bool
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) CreateUser(user dto.UserCreateDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.userRepository.InsertUser((userToCreate))
	return res
}

func (service *userService) UpdateUser(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.userRepository.UpdateUser((userToUpdate))
	return res
}

func (service *userService) Profile(userId string) entity.User {
	return service.userRepository.ProfileUser(userId)
}

func (service *userService) FindByUsername(username string) entity.User {
	return service.userRepository.FindByUsername(username)
}

func (service *userService) IsDuplicateUsername(username string) bool {
	res := service.userRepository.IsDuplicateUsername(username)
	return !(res.Error == nil)
}
