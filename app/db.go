package app

import (
	"github.com/omidgz/order-api/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (this *App) loadDB() {
	// Open a new connection to our sqlite database.
	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to open the SQLite database.")
	}

	db.AutoMigrate(model.Order{})
	this.DB = db
}
