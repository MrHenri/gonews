package migration

import (
	"github.com/MrHenri/gonews/db"
	"github.com/MrHenri/gonews/models"
)

func AutoMigration() {
	db := db.Connect()
	defer db.Close()
	db.AutoMigrate(models.News{})
}
