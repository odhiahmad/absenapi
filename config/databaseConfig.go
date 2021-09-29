package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/odhiahmad/absenapi/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//setup database
func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		panic("Gagal load file env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	// dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	dsn := fmt.Sprintf("host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=absenapi port=" + dbPort + " sslmode=disable TimeZone=Asia/Shanghai")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Gagal membuat koneksi ke database")
	}
	db.AutoMigrate(&entity.User{})
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		panic("Gagal menutup koneksi dari database")
	}

	dbSQL.Close()
}
