package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToMySQL(dataSourceName string) (*gorm.DB, error) {
	// Open the database connection
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	err = sqlDB.Ping()
	if err != nil {
		sqlDB.Close()
		return nil, err
	}

	return db, nil
}
