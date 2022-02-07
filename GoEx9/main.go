package main

import (
	"GoEx8/component"
	"GoEx8/modules/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

//type User struct {
//	UserId          int    `json:"id,omitempty" gorm:"column:id;"`
//	UserFirstname   string `json:"user_firstname" gorm:"column:user_firstname;"`
//	UserLastname    string `json:"user_lastname" gorm:"column:user_lastname;"`
//	UserPhonenumber string `json:"user_phonenumber" gorm:"column:user_phonenumber;"`
//	Username        string `json:"username" gorm:"column:username;"`
//	Password        string `json:"password" gorm:"column:password;"`
//}
//
//type UserUpdate struct {
//	UserFirstname   *string `json:"user_firstname" gorm:"column:user_firstname"`
//	UserLastname    *string `json:"user_lastname" gorm:"column:user_lastname"`
//	UserPhonenumber *string `json:"user_phonenumber" gorm:"column:user_phonenumber;"`
//	Username        *string `json:"username" gorm:"column:username;"`
//	Password        *string `json:"password" gorm:"column:password;"`
//}
//
//func (User) TableName() string {
//	return "users"
//}
//
//func (UserUpdate) TableName() string {
//	return User{}.TableName()
//}

func main() {

	dsn := os.Getenv("DBConnection")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//CRUD

	appCtx := component.NewAppContext(db)

	user := r.Group("/users")
	{
		user.POST("", ginuser.CreateUser(appCtx))

		user.GET("/:id", ginuser.GetUser(appCtx))

		user.GET("", ginuser.ListUser(appCtx))

		user.PATCH("/:id", ginuser.UpdateUser(appCtx))

		user.DELETE("/:id", ginuser.DeleteUser(appCtx))
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
