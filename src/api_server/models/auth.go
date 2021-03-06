package models

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	common "../common"
)

// Token function
type Token struct {
	jwt.StandardClaims
	UserID    uuid.UUID `gorm:"type:char(36);"`
	ProjectID uuid.UUID `gorm:"type:char(36);"`
	Token     string    `json:"token"`
}

// PublishToken function
func PublishToken(username, password string, projectID uuid.UUID) map[string]interface{} {

	fmt.Println(username, projectID)
	user := &User{}
	err := GetDB().Table("users").Where("name = ? and project_id = ?", username, projectID).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.Message(false, "user is not found")
		}
		return common.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return common.Message(false, "Invalid login credentials. Please try again")
	}
	user.Password = ""

	// Create JWT token
	tk := &Token{UserID: user.ID, ProjectID: user.ProjectID}
	tkHS256 := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := tkHS256.SignedString([]byte(os.Getenv("token_password")))
	tk.Token = tokenString //Store the token in the response

	resp := common.Message(true, "success")
	resp["auth"] = tk
	return resp
}
