package models

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
)

//User model
type User struct {
	gorm.Model
	FirstName     string `json:"FirstName,omitempty" gorm:"column:FirstName;type:varchar(256)"`
	LastName      string `json:"LastName,omitempty" gorm:"column:LastName;type:varchar(256)"`
	Email         string `json:"Email,omitempty" gorm:"column:Email;type:varchar(256);unique"`
	EmailVerified bool   `json:"EmailVerified" gorm:"column:EmailVerified;type:varchar(256)"`
	Password      string `json:"-" gorm:"column:Password;type:varchar(256)"`
}

//TestPassword matches stored hash
func (u *User) TestPassword(password string) (bool, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return false, err
	}
	return string(hash) == u.Password, nil
}

//SetPassword for this user
func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}
