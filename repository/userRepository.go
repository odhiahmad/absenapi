package repository

import (
	"log"

	"github.com/odhiahmad/absenapi/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(username string, password string) interface{}
	IsDuplicateUsername(username string) (tx *gorm.DB)
	FindByUsername(username string) entity.User
	ProfileUser(userId string) entity.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	
	return user
}

func (db *userConnection) UpdateUser(user entity.User) entity.User {

	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser entity.User
		db.connection.Find(&tempUser, user.Id)
		user.Password = tempUser.Password
	}

	db.connection.Save(&user)
	
	return user
}

func (db *userConnection) VerifyCredential(username string, password string) interface{} {
	var user entity.User
	res := db.connection.Where("username = ?", username).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateUsername(username string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("username = ?", username).Take(&user)
}

func (db *userConnection) FindByUsername(username string) entity.User {
	var user entity.User
	db.connection.Where("username = ?", username).Take(&user)

	return user
}

func (db *userConnection) ProfileUser(userId string) entity.User {
	var user entity.User
	db.connection.Find(&user, userId)

	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
