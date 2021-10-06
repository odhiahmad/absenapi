package connect

import (
	"bri-rece/api/models"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	cfg *viper.Viper
	db  *gorm.DB
	err error
)

type Connect interface {
	SqlDb() *gorm.DB
	Config() *viper.Viper
	ApiServer() string
}

type connect struct{}

func NewConnect() Connect {
	return &connect{}
}

func (i *connect) SqlDb() *gorm.DB {
	dbUser := i.Config().GetString("DB_USER")
	dbPassword := i.Config().GetString("DB_PASSWORD")
	dbHost := i.Config().GetString("DB_HOST")
	dbPort := i.Config().GetString("DB_PORT")
	dbName := i.Config().GetString("DB_NAME")

	db, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s ", dbHost, dbUser, dbPassword, dbName, dbPort)))

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection established")
	}
	db.Debug().AutoMigrate(&models.User{}, &models.Account{}, &models.Wallet{}, &models.WalletHistory{}, &models.Otp{})
	db.Debug().Model(&models.Account{})
	return db
}

func (i *connect) Config() *viper.Viper {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.ReadInConfig()
	cfg = viper.GetViper()
	return cfg
}

func (i *connect) ApiServer() string {
	host := i.Config().GetString("HTTPHOST")
	port := i.Config().GetString("HTTPPORT")
	return fmt.Sprintf("%s:%s", host, port)
}
