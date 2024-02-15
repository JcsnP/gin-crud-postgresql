package db

import (
	"fmt"
	"log"
	"os"

	"github.com/JcsnP/gin-clean/pkg/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDatabase() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error loading .env file")
	}

	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresDb := os.Getenv("POSTGRES_DB")
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable Timezone=Asia/Bangkok", postgresHost, postgresUser, postgresPassword, postgresDb, postgresPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("can not connect to database")
	}

	AutoMigrate(db)

	return db
}

func GetDBInstance() *gorm.DB {
	return db
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(models.User{},)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}