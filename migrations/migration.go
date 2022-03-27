package migrations

import (
	"fmt"
	"log"

	"github.com/Johnogram/golang-todos/config"
	"github.com/Johnogram/golang-todos/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migrate() {
	log.SetPrefix("db: ")
	log.SetFlags(0)

	config, err := config.GetConfig()

	if err != nil {
		log.Fatal(err)
		return
	}

	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.MySQL.Username,
		config.MySQL.Password,
		config.MySQL.Host,
		config.MySQL.Port,
		config.MySQL.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return
	}

	db.AutoMigrate(&models.Task{})
}
