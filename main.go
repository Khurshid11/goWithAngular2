package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"log"

	"github.com/labstack/echo/engine/fasthttp"
	
)

var db *gorm.DB

func main() {
	e := echo.New()
	userName := "root"
	password := "password"
	dbName := "db_mks"
	params := "charset=utf8&parseTime=True&loc=Local"

	ConStr := userName + ":" + password + "@/" + dbName + "?" + params

	db, _ = gorm.Open("mysql", ConStr)
  log.Println("Starting server..................")
  	e.Static("/", "./client/dist")
	e.GET("/addUser",AddUser)
	e.POST("/addUser",AddUser)

	e.Run(fasthttp.New(":3000"))
}

type User struct{
	FirstName string  `json:"first_name"`
	LastName string		`json:"last_name"`
	Email string 	`json:"email"`
	Password string	`json:password`
}

func AddUser(c echo.Context) error{
	var user User

	
	if err := c.Bind(&user); err != nil{
		return err
	}

		if !db.HasTable(&user) {
			db.CreateTable(&user)
		}

		db.Create(&user)
	
		log.Println("Testing................",user)

		return nil
		//return c.String(http.StatusOK, "Welcome")
}
