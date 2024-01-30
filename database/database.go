package database

import (
	"log"
	"os"

	"github.com/gabriel-kimutai/ano2tes/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBinstance struct {
	Db *gorm.DB
}

var DB DBinstance

// 5ua8NUCyOVo2A32O

func ConnectDB() {
	// dsn := fmt.Sprintf(
	// 	// "host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
	// 	"user=postgres.rmuaaruwspzbnskfyhzj password=%s host=aws-0-eu-central-1.pooler.supabase.com port=6543 dbname=postgres",
	// 	os.Getenv("DB_PASSWORD"),
	// )
	dsn := "postgres://postgres.rmuaaruwspzbnskfyhzj:tQrk8eoMfvSqCiUm@aws-0-eu-central-1.pooler.supabase.com:6543/postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to db. \n", err)
		os.Exit(2)
	}
	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.User{}, &models.Note{})
	DB = DBinstance{
		Db: db,
	}
}
