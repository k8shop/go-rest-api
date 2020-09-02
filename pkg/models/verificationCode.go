package models

import (
	"time"

	"crypto/rand"

	"github.com/jinzhu/gorm"
)

//VerificationCode model
type VerificationCode struct {
	gorm.Model
	Code      string    `json:"Code,omitempty" gorm:"column:Code;type:varchar(256)"`
	Expiry    time.Time `json:"Expiry,omitempty" gorm:"column:Expiry;type:timestamp"`
	AccountID int       `json:"AccountID,omitempty" gorm:"column:AccountID;type:int"`
	Used      bool      `json:"Used"`
}

//NewVerificationCode for account id
func NewVerificationCode(accountID int) (*VerificationCode, error) {
	code, err := generateCode(6)
	if err != nil {
		return nil, err
	}
	return &VerificationCode{
		Code:      code,
		AccountID: accountID,
		Expiry:    time.Now().Add(24 * time.Hour),
	}, nil
}

func generateCode(length int) (string, error) {
	codeChars := "1234567890"
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	codeCharsLength := len(codeChars)
	for i := 0; i < length; i++ {
		buffer[i] = codeChars[int(buffer[i])%codeCharsLength]
	}

	return string(buffer), nil
}
