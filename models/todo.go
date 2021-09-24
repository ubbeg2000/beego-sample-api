package models

import (
	"sample-api/database"
)

type Todo struct {
	Base
	Text string `json:"text" gorm:"not null"`
	Done bool   `json:"done" gorm:"default:false"`
}

func AddTodo(t *Todo) (id int64, err error) {
	db := database.Conn

	tx := db.Create(t)

	if tx.Error != nil {
		return 0, tx.Error
	}

	return t.Id, nil
}

func GetTodoById(id int64) (t *Todo, err error) {
	var retval Todo
	db := database.Conn

	tx := db.First(&retval, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &retval, nil
}

func GetAllTodo(page int, limit int) (t *[]Todo, err error) {
	var retval []Todo

	db := database.Conn

	tx := db.Offset((page - 1) * limit).Limit(limit).Find(&retval)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &retval, nil
}

func UpdateTodoById(t *Todo) error {
	db := database.Conn

	tx := db.Model(t).Updates(t)

	return tx.Error
}

func DeleteTodo(id int64) error {
	db := database.Conn

	tx := db.Delete(&Todo{}, id)

	return tx.Error
}
