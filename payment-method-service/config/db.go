package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	username := EnvConfig.MYSQL_USERNAME
	password := EnvConfig.MYSQL_PASSWORD
	host := EnvConfig.MYSQL_HOST
	port := EnvConfig.MYSQL_PORT
	dbname := EnvConfig.MYSQL_DATABASE

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}

func GetDB() *gorm.DB {
	return DB
}
