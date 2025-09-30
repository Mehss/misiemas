package config

import (
	"log"
	"os"
	"tripatra-dct-service-config/database/model"
	"tripatra-dct-service-config/database/model/module"
	"tripatra-dct-service-config/database/model/supplier"
	"tripatra-dct-service-config/database/model/user"

	"gorm.io/gorm"
)

// Migrate function to create or update tables based on your models
func Migrate(db *gorm.DB) error {
	// AutoMigrate will create tables, missing columns, and indexes
	// if env migration mode is false (production)
	if os.Getenv("MIGRATION_MODE") == "false" {
		return nil
	} else {
		// Ensure schemas exist
		// Create the dev schema
		if err := db.Exec("CREATE SCHEMA IF NOT EXISTS user_management").Error; err != nil {
			return err
		}
		if err := db.Exec("CREATE SCHEMA IF NOT EXISTS settings").Error; err != nil {
			return err
		}

		err := db.AutoMigrate(
			&user.User{},
			&user.Permission{},
			&module.Module{},
			&supplier.Vendor{},
			&model.Notifications{},
			// Add other models here
		)
		if err != nil {
			log.Printf("Error during AutoMigrate: %v", err) // Log the error and continue
			// return err
		}
	}

	// // Add foreign key for Role if it doesn't exist
	// if !db.Migrator().HasConstraint(&user.User{}, "Role") {
	// 	err = db.Migrator().CreateConstraint(&user.User{}, "Role")
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// // Add foreign key for CreatedBy if it doesn't exist
	// if !db.Migrator().HasConstraint(&user.User{}, "CreatedBy") {
	// 	err = db.Migrator().CreateConstraint(&user.User{}, "CreatedBy")
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// // Add foreign key for ModifiedBy if it doesn't exist
	// if !db.Migrator().HasConstraint(&user.User{}, "ModifiedBy") {
	// 	err = db.Migrator().CreateConstraint(&user.User{}, "ModifiedBy")
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil
}
