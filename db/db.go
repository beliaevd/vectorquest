package db

import (
	"Vectorquest/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var db *gorm.DB

var db_conf = config.DbConfig()

func init() {

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", db_conf.DbHost, db_conf.DbUser, db_conf.DbName, db_conf.DbPass)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}
	db = conn
}

func GetDB() *gorm.DB {
	return db
}
