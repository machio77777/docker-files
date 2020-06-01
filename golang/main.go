package main

import (
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
	
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)

type Users struct {
    ID    int
    Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	db, err := sqlConnect()
	if err != nil {
        panic(err.Error())
	}
	
	defer db.Close()

	result := []*Users{}
	db.Find(&result)

	bytes, err := json.Marshal(result)

	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, string(bytes))
	})
	r.Run(":8080")
}

func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
    USER := "user"
    PASS := "password"
    PROTOCOL := "tcp(mysql_host:3306)"
	DBNAME := "sample_db"
	
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
    return gorm.Open(DBMS, CONNECT)
}