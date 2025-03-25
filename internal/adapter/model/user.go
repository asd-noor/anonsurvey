package model

import "fmt"

type User struct {
	ID                       string
	Name                     string
	Email                    string
	Password                 string
	IsAdmin                  bool
	LastSurveySubmissionDate string
}

type UserList []User

type UserModeler interface {
	UserModel() User
}

func (u User) InsertQuery() string {
	return ""
}

func (u User) SelectQuery() string {
	return fmt.Sprintf("SELECT * FROM users WHERE ID = %s", u.ID)
}
