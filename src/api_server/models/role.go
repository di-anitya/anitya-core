package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"

	common "../common"
)

// Role struct
type Role struct {
	Base
	ProjectID uuid.UUID `gorm:"type:char(36);" json:"project_id"`
	Name      string    `sql:"size:60"`
	IsAdmin   bool      `sql:"size:60"`
}

// ValidateRole function
func (role *Role) ValidateRole() (map[string]interface{}, bool) {

	if role.Name == "" {
		return common.Message(false, "'role_name' is required"), false
	}

	// Email must be unique
	temp := &Role{}

	// check for errors and duplicate emails
	err := GetDB().Table("roles").Where("name = ?", role.Name).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return common.Message(false, "Connection error. Please retry"), false
	}

	if temp.Name != "" {
		return common.Message(false, "'name' already in use by another role."), false
	}

	return common.Message(false, "Requirement passed"), true
}

// CreateRole function
func (role *Role) CreateRole() map[string]interface{} {

	if resp, ok := role.ValidateRole(); !ok {
		return resp
	}

	GetDB().Create(role)

	resp := common.Message(true, "success")
	//resp["auth"] = tk
	resp["role"] = role
	return resp
}

// ListRole function
func ListRole() *Role {
	role := &Role{}
	err := GetDB().Table("roles").Find(role).Error
	if err != nil {
		return nil
	}
	return role
}

// ShowRole function
func ShowRole(id string) *Role {
	role := &Role{}
	err := GetDB().Table("roles").Where("id = ?", id).First(role).Error
	if err != nil {
		return nil
	}
	return role
}

// ModifyRole function
func ModifyRole(id string) *Role {
	role := &Role{}
	err := GetDB().Table("roles").Where("id = ?", id).Update(role).Error
	if err != nil {
		return nil
	}
	return role
}

// DeleteRole function
func DeleteRole(id string) {
	role := &Role{}
	GetDB().Table("roles").Where("id = ?", id).Delete(role)
}
