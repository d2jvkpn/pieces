package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	addr string = ":8080"
	db   *gorm.DB

	config map[string]string = map[string]string{
		"user":     "root",
		"password": "rover",
		"uri":      "172.17.0.2:3306",
		"db":       "rover",
		"table":    "players",
	}
)

func main() {
	var (
		err    error
		router *gin.Engine
	)

	if db, err = NewDB(config); err != nil {
		log.Fatal(err)
	}

	router = gin.Default()
	router.GET("/", Hello)
	router.Run(addr)
}

func Hello(c *gin.Context) {
	results := make([]map[string]interface{}, 0)

	state := fmt.Sprintf("select * from %s limit 20;", config["table"])
	if err := db.Raw(state).Scan(&results).Error; err != nil {
		fmt.Printf(">>> error: %v\n", err)
		c.JSON(http.StatusInternalServerError, nil)
	} else {
		c.JSON(http.StatusOK, results)
	}
	return
}

func NewDB(mp map[string]string) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mp["user"], mp["password"], mp["uri"], mp["db"],
	)

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
