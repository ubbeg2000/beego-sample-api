package models

import (
	db "sample-api/database"
)

type User struct {
	Base
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password,omitempty" gorm:"not null"`
	Name     string `json:"name"`
}

func AddUser(t *User) (id int64, err error) {
	tx := db.Conn.Create(t)

	if tx.Error != nil {
		return 0, tx.Error
	}

	return t.Id, nil
}

func GetUserById(id int64) (u *User, err error) {
	var retval User
	tx := db.Conn.First(&retval, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &retval, nil
}

func GetUserByUsername(uname string) (u *User, err error) {
	var retval User
	tx := db.Conn.Where("username = ?", uname).First(&retval)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &retval, nil
}

func GetAllUser(page int, limit int) (u *[]User, err error) {
	var retval []User

	tx := db.Conn.Offset((page - 1) * limit).Limit(limit).Find(&retval)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &retval, nil
}

func UpdateUserById(u *User) error {
	tx := db.Conn.Model(u).Updates(u)

	return tx.Error
}

func DeleteUser(id int64) error {
	tx := db.Conn.Delete(&User{}, id)

	return tx.Error
}
