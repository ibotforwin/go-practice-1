package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           uint   `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email" gorm:"unique"`
	Password     []byte `json:"-"`
	IsAmbassador bool   `json:"-"`
}

func (user *User) SetPassword(password string) {
	//Encrypt password, and the second return item is an error which we are ignoring.
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword

}
func (user *User) CheckPassword(password string) error {
	//Encrypt password, and the second return item is an error which we are ignoring.
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	return err
}
