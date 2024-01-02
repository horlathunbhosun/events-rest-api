package models

import (
	"errors"
	"time"

	"github.com/horlathunbhosun/events-rest-api/db"
	"github.com/horlathunbhosun/events-rest-api/internal/utility"
)

type User struct {
	ID           int64
	Email        string `binding:"required"`
	Password     string `binding:"required"`
	DatedCreated time.Time
}

func (u *User) Save() error {
	u.DatedCreated = time.Now()
	query := `INSERT INTO users(email,password, datecreated) 
		VALUES (?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	passwordHashed, err := utility.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, passwordHashed, u.DatedCreated)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = userId

	return err
}

func (u *User) ValidateUserCredential() error {
	query := "SELECT id,password FROM users WHERE email=?"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("credential invalid")
	}

	passwordIsValid := utility.CompareHashedPassword(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credential invalid")
	}

	return nil

}
