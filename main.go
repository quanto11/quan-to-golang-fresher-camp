package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	UserId          int    `json:"user_id,omitempty" gorm:"column:user_id"`
	UserFirstname   string `json:"user_firstname" gorm:"column:user_firstname"`
	UserLastname    string `json:"user_lastname" gorm:"column:user_lastname"`
	UserPhonenumber string `json:"user_phonenumber" gorm:"column:user_phonenumber"`
	Username        string `json:"username" gorm:"column:username"`
	Password        string `json:"password" gorm:"column:password"`
}

type UserUpdate struct {
	UserFirstname *string `json:"user_firstname" gorm:"column:user_firstname"`
	UserLastname  *string `json:"user_lastname" gorm:"column:user_lastname"`
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
	newUser := User{UserFirstname: "To", UserLastname: "Quan", UserPhonenumber: "0962xxx166", Username: "tohongquan", Password: "abc123"}
	if err := db.Create(&newUser); err != nil {
		fmt.Println(err)
	}

	// SELECT
	//SELECT list user
	var users []User
	db.Where("status=?", 1).Find(&users)
	fmt.Println("List users:")
	for _, i := range users {
		fmt.Printf("\tId:%v\tFirstname:%v\tLastname:%v\tPhonenumber:%v\tUsername:%v\tPassword:%v\n", i.UserId, i.UserFirstname, i.UserLastname, i.UserPhonenumber, i.Username, i.Password)
	}

	//SELECT a user
	var user User
	if err := db.Where("user_id = 1").First(&user); err == nil {
		fmt.Println(err)
	}
	fmt.Printf("User\t:\tId:%v\tFirstname:%v\tLastname:%v\tPhonenumber:%v\tUsername:%v\tPassword:%v\n", user.UserId, user.UserFirstname, user.UserLastname, user.UserPhonenumber, user.Username, user.Password)
	//UPDATE
	var updateUser User
	updateUser.UserFirstname = "Hong"
	updateUser.UserLastname = "Quan"
	db.Table(User{}.TableName()).Where("user_id = 5").Updates(&updateUser)
	//possible UPDATE EMPTY value
	emptyStr := ""
	db.Table(User{}.TableName()).Where("user_id = 4").Updates(UserUpdate{&emptyStr, &emptyStr})

	//DELETE
	db.Table(User{}.TableName()).Where("user_id = 6").Delete(nil)
}
