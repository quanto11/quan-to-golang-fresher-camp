package main

import (
	"GoEx8/component"
	"GoEx8/modules/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

type User struct {
	UserId          int    `json:"id,omitempty" gorm:"column:id;"`
	UserFirstname   string `json:"user_firstname" gorm:"column:user_firstname;"`
	UserLastname    string `json:"user_lastname" gorm:"column:user_lastname;"`
	UserPhonenumber string `json:"user_phonenumber" gorm:"column:user_phonenumber;"`
	Username        string `json:"username" gorm:"column:username;"`
	Password        string `json:"password" gorm:"column:password;"`
}

type UserUpdate struct {
	UserFirstname   *string `json:"user_firstname" gorm:"column:user_firstname"`
	UserLastname    *string `json:"user_lastname" gorm:"column:user_lastname"`
	UserPhonenumber *string `json:"user_phonenumber" gorm:"column:user_phonenumber;"`
	Username        *string `json:"username" gorm:"column:username;"`
	Password        *string `json:"password" gorm:"column:password;"`
}

func (User) TableName() string {
	return "users"
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

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

		user.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})

				return
			}

			var data User

			if err := db.Where("id = ?", id).First(&data).Error; err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})

				return
			}

			c.JSON(http.StatusOK, data)
		})

		user.GET("", ginuser.ListUser(appCtx))

		user.PATCH("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(401, map[string]interface{}{
					"error": err.Error(),
				})

				return
			}

			var data UserUpdate

			if err := c.ShouldBind(&data); err != nil {
				c.JSON(401, map[string]interface{}{
					"error": err.Error(),
				})

				return
			}

			if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
				c.JSON(401, map[string]interface{}{
					"error": "Update failed",
				})

				return
			}

			c.JSON(http.StatusOK, map[string]interface{}{
				"update successfully": 1,
			})
		})

		user.DELETE("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(401, map[string]interface{}{
					"error": err.Error(),
				})

				return
			}

			if err := db.Table(User{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
				c.JSON(401, map[string]interface{}{
					"error": err.Error(),
				})

				return
			}

			c.JSON(200, map[string]interface{}{"ok": 1})
		})
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
