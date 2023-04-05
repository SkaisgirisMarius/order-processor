package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const connectionString = "user:password@tcp(localhost:3306)/mydatabase?charset=utf8mb4&parseTime=True&loc=Local"

func ConnectToMySQL() (*gorm.DB, error) {
	// Open the database connection
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
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
