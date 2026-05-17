package storage

import (
	"github.com/NormoBoat/shortlink/models"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

func New() *Storage {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err)
	}

	makeMigrate(db)

	return &Storage{
		DB: db,
	}
}

func makeMigrate(db *gorm.DB) {
	mustMigrate(db, models.User{})
	mustMigrate(db, models.Symlink{})
}

func mustMigrate(db *gorm.DB, model any) {
	if err := db.AutoMigrate(&model); err != nil {
		log.Panic().Err(err)
	}
}
