package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hackfest/model"
)

func MakeDatabaseConnection(data *Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s "+
		"password=%s "+
		"host=%s "+
		"TimeZone=Asia/Singapore "+
		"port=%s "+
		"dbname=%s",
		data.SupabaseUser, data.SupabasePassword, data.SupabaseHost, data.SupabasePort, data.SupabaseDbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&model.User{},
		&model.Fleet{},
		&model.History{},
		&model.Status{},
		&model.Fleet{},
	); err != nil {
		return nil, err
	}
	return db, nil
}
