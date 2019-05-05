package models

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	common "../common"
)

// User struct
type User struct {
	Base
	ProjectID uuid.UUID `gorm:"primary_key;type:char(36);" json:"project_id"`
	RoleID    uuid.UUID `gorm:"type:char(36);NOT NULL" json:"role_id"`
	Name      string    `json:"user_name"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

// ValidateUser function
func (user *User) ValidateUser() (map[string]interface{}, bool) {

	if user.Name == "" {
		return common.Message(false, "'user_name' is required"), false
	}

	if !strings.Contains(user.Email, "@") {
		return common.Message(false, "'email' is required"), false
	}

	if len(user.Password) < 6 {
		return common.Message(false, "'password' is required"), false
	}

	// Email must be unique
	temp := &User{}

	// check for errors and duplicate emails
	err := GetDB().Table("users").Where("name = ?", user.Name).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		println(err)
		return common.Message(false, "Connection error. Please retry"), false
	}

	if temp.Name != "" {
		return common.Message(false, "'name' already in use by another user."), false
	}

	return common.Message(false, "Requirement passed"), true
}

// CreateUser function
func (user *User) CreateUser() map[string]interface{} {

	if resp, ok := user.ValidateUser(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	GetDB().Create(user)

	resp := common.Message(true, "success")
	//resp["auth"] = tk
	resp["user"] = user
	return resp
}

// ListUser function
func ListUser() []User {
	user := []User{}
	GetDB().Find(&user)
	return user
}

// ShowUser function
func ShowUser(id string) *User {
	user := &User{}
	err := GetDB().Table("users").Where("id = ?", id).First(user).Error
	if err != nil {
		println(err)
		return nil
	}
	return user
}

// ModifyUser function
func (user *User) ModifyUser(id string) {
	err := GetDB().Table("users").Where("id = ?", id).Update(user).Error
	if err != nil {
		println(err)
		return
	}
}

// DeleteUser function
func DeleteUser(id string) {
	user := &User{}
	fmt.Println("id:", id)
	GetDB().Table("users").Unscoped().Where("id = ?", id).Delete(user)
}
