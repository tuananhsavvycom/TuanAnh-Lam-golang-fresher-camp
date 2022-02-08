package usermodel

import (
	"errors"
	"strings"
)

const EntityName = "User"

type User struct {
	//common.SQLModel `json:",inline"`
	UserId          int    `json:"id,omitempty" gorm:"column:id;"`
	UserFirstname   string `json:"user_firstname" gorm:"column:user_firstname;"`
	UserLastname    string `json:"user_lastname" gorm:"column:user_lastname;"`
	UserPhonenumber string `json:"user_phonenumber" gorm:"column:user_phonenumber;"`
	Username        string `json:"username" gorm:"column:username;"`
	Password        string `json:"password" gorm:"column:password;"`
	Status          int
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	UserId          int    `json:"id,omitempty" gorm:"column:id;"`
	UserFirstname   string `json:"user_firstname" gorm:"column:user_firstname"`
	UserLastname    string `json:"user_lastname" gorm:"column:user_lastname"`
	UserPhonenumber string `json:"user_phonenumber" gorm:"column:user_phonenumber;"`
	Username        string `json:"username" gorm:"column:username;"`
	Password        string `json:"password" gorm:"column:password;"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (user *UserCreate) Validate() error {
	user.UserFirstname = strings.TrimSpace(user.UserFirstname)
	user.UserLastname = strings.TrimSpace(user.UserLastname)

	if len(user.UserFirstname) == 0 {
		return errors.New("user first name can not be blank")
	}

	if len(user.UserLastname) == 0 {
		return errors.New("user last name can not be blank")
	}

	return nil
}

type UserUpdate struct {
	UserFirstname   *string `json:"user_firstname" gorm:"column:user_firstname"`
	UserLastname    *string `json:"user_lastname" gorm:"column:user_lastname"`
	UserPhonenumber *string `json:"user_phonenumber" gorm:"column:user_phonenumber;"`
	Username        *string `json:"username" gorm:"column:username;"`
	Password        *string `json:"password" gorm:"column:password;"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}
