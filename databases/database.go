package databases

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// DB is a global var for connect DB
var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host   string
	Port   string
	User   string
	DBName string
	// Password string
}

// BuildDBConfig use for building DB config
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:   "127.0.0.1",
		Port:   "23306",
		User:   "root",
		DBName: "golang_test",
		// Password: "password",
	}
	return &dbConfig
}

// DbURL use for create DB connection URL
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		// dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
