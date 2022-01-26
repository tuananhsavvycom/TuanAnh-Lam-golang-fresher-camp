package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	User_id  int    `json:"user_id,omitempty" gorm:"column:user_id;"`
	Name     string `json:"name" gorm:"column:name;"`
	Password string `json:"password" gorm:"column:password;"`
	Email    string `json:"email" gorm:"column:email;"`
	Phone    int    `json:"phone" gorm:"column:phone;"`
	Address  string `json:"address" gorm:"column:address;"`
}

type UserUpdate struct {
	Name *string `json:"name" gorm:"column:name"`
}

func (User) TableName() string {
	return "Users"
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("DBconnectionString")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println(db, err)
	//insert
	newUser := User{
		Name:     "TuanAnh",
		Password: "123456",
		Email:    "lta@gmail.com",
		Phone:    123456789,
		Address:  "Ha Noi",
	}

	if err := db.Create(&newUser).Error; err != nil {
		fmt.Println(err)
	}

	//search array
	var users []User
	db.Where("name = ?", "TuanAnh").Find(&users)
	fmt.Println(users)

	//search one
	var user User
	if err := db.Where("user_id = 1").First(&user); err != nil {
		log.Println(err)
	}
	fmt.Println(user)

	//update
	newName := "Lam Tuan Anh"
	if err := db.Table(User{}.TableName()).Where("user_id = 2").Updates(&UserUpdate{
		Name: &newName,
	}); err != nil {
		log.Println(err)
	}

	//delete
	db.Table(User{}.TableName()).Where("user_id = 2").Delete(nil)

}
