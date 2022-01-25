package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	User_id          int    `json:"user_id,omitempty" gorm:"column:user_id"`
	User_firstname   string `json:"user_firstname" gorm:"column:user_firstname"`
	User_lastname    string `json:"user_lastname" gorm:"column:user_lastname"`
	User_phonenumber string `json:"user_phonenumber" gorm:"column:user_phonenumber"`
	Username         string `json:"username" gorm:"column:username"`
	Password         string `json:"password" gorm:"column:password"`
}

type UserUpdate struct {
	User_firstname *string `json:"user_firstname" gorm:"column:user_firstname"`
	User_lastname  *string `json:"user_lastname" gorm:"column:user_lastname"`
}

func (User) TableName() string {
	return "users"
}

func main() {
	dsn := os.Getenv("DBConnection")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// INSERT
	newUser := User{User_firstname: "To", User_lastname: "Quan", User_phonenumber: "0962xxx166", Username: "tohongquan", Password: "abc123"}
	if err := db.Create(&newUser); err != nil {
		fmt.Println(err)
	}

	// SELECT
	//SELECT list user
	var users []User
	db.Where("status=?", 1).Find(&users)
	fmt.Println("List users:", users)
	//SELECT a user
	var user User
	if err := db.Where("user_id = 1").First(&user); err != nil {
		fmt.Println(err)
	}
	fmt.Println("user:", user)

	//UPDATE
	var updateUser User
	updateUser.User_firstname = "Hong"
	updateUser.User_lastname = "Quan"
	db.Table(User{}.TableName()).Where("user_id = 5").Updates(&updateUser)
	//possible UPDATE EMPTY value
	emptyStr := ""
	db.Table(User{}.TableName()).Where("user_id = 4").Updates(UserUpdate{&emptyStr, &emptyStr})

	//DELETE
	db.Table(User{}.TableName()).Where("user_id = 6").Delete(nil)
}
