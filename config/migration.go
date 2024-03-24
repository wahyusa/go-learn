package config

import "go-learn/models"

func MigrateDatabase() {
	// DEV ONLY auto truncate
	DBCON.Migrator().DropTable(&models.User{})

	// generate tables
	DBCON.AutoMigrate(&models.User{})
}
