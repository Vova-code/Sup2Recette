package database

import (
	"GoApi/models/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var recipes = []*database.Recipe{
	{
		Title:       "Poulet Frites",
		Steps:       []string{"Faire les frites", "Cuire le poulet"},
		Evaluations: 0,
	},
	{
		Title:       "Salade",
		Steps:       []string{"Laver la salade", "Couper les tomates", "Faire la vinaigrette"},
		Evaluations: 2,
	},
}

func ConnectToDatabase() *gorm.DB {
	dsn := ""
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Impossible to connect to database: %v", err)
	}

	return db
}

func InitDb() {
	db := ConnectToDatabase()

	if db.Migrator().HasTable(&database.Recipe{}) {
		err := db.Migrator().DropTable(&database.Recipe{})
		if err != nil {
			log.Fatalf("Failed to drop table: %v", err)
		}
	}

	err := db.AutoMigrate(&database.Recipe{})
	if err != nil {
		log.Fatalf("Impossible to migrate table Recipe: %v", err)
	}

	db.Create(&recipes)
}
