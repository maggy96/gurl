package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

type Url struct {
	ID uint `json:"id"`
	Payload string `json:"payload"`
}


func main() {
	db, _ = gorm.Open("sqlite3", "./api.db")
	defer db.Close()
	db.AutoMigrate(&Url{})
	r := setupRouter()
	r.Run(":8080")
}

func shorten(c *gin.Context)  {
	url := c.Param("url")
	u1 := Url{Payload: url[1:len(url)]}
	db.Create(&u1)
	c.JSON(200, u1)
}

func resolve(c *gin.Context)  {
	id := c.Param("id")
	var url Url
	db.Where("id = ?", id).First(&url)
	c.Redirect(http.StatusTemporaryRedirect, url.Payload)
}
