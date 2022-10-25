package models

import (
	"todo-app/utils"

	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Content string           `json:"content"`
	Status  utils.StatusEnum `json:"status"`
	UserID  uint
}

func (todo *Todo) TableName() string {
	return "todos"
}
