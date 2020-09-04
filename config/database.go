package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// DB ... gorm db
var DB *gorm.DB

// DBConfig default
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig ... DB configuration
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		DBName:   "simple_go",
		Password: "",
	}

	return &dbConfig
}

// DbURL ... Setting up the db config
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
