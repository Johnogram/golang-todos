package migrations

import (
	"log"

	"github.com/Johnogram/golang-todos/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	log.SetPrefix("db: ")
	log.SetFlags(0)

	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return
	}

	db.AutoMigrate(&models.Task{})
}
