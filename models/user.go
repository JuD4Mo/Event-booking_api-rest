package models

import (
	"example.com/event-booking/db"
	"example.com/event-booking/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := `
	INSERT INTO users (email, password) 
	VALUES
	(?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPasswd, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(user.Email, hashedPasswd)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	user.ID = userId
	return err
}
