package migration

import (
	"github.com/MrHenri/gonews/db"
	"github.com/MrHenri/gonews/models"
)

func AutoMigration() {
	db := db.Connect()
	db.Debug().DropTableIfExists(&models.News{})
	//Drops table if already exists
	db.Debug().AutoMigrate(&models.News{})
	//Auto create table based on Model
}
