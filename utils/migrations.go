package utils

import (
	"github.com/BinayRajbanshi/go-auth/database"
	"github.com/BinayRajbanshi/go-auth/models"
)

func Migrate() {
	database.DB.AutoMigrate(&models.User{})
}
