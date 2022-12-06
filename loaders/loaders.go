package loaders

import (
	"bbs_backend/configs"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Create a new connection to database
func SetDBConnection() *gorm.DB{
	db_user := configs.GetEnvConfigs("DB_USER")
	db_passwd := configs.GetEnvConfigs("DB_PASSWARD")
	db_host := configs.GetEnvConfigs("DB_HOST")
	db_name := configs.GetEnvConfigs("DB_NAME")
	db_port := configs.GetEnvConfigs("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",db_user, db_passwd, db_host, db_port, db_name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	return db
}

// Close a connection between app and database
func CloseDBConnection(db *gorm.DB){
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection to database")
	}

	dbSQL.Close()
}