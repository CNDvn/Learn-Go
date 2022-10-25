package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Todos    []Todo
}

func (user *User) TableName() string {
	return "users"
}
