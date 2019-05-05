package models

import (
	"github.com/jinzhu/gorm"

	common "../common"
)

// Project struct
type Project struct {
	Base
	Name string `json:"project_name"`
}

// ValidateProject function
func (project *Project) ValidateProject() (map[string]interface{}, bool) {

	if project.Name == "" {
		return common.Message(false, "'project_name' is required"), false
	}

	// Email must be unique
	temp := &Project{}

	// check for errors and duplicate emails
	err := GetDB().Table("projects").Where("name = ?", project.Name).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return common.Message(false, "Connection error. Please retry"), false
	}

	if temp.Name != "" {
		return common.Message(false, "'name' already in use by another project."), false
	}

	return common.Message(false, "Requirement passed"), true
}

// CreateProject function
func (project *Project) CreateProject() map[string]interface{} {

	if resp, ok := project.ValidateProject(); !ok {
		return resp
	}

	GetDB().Create(project)

	resp := common.Message(true, "success")
	//resp["auth"] = tk
	resp["project"] = project
	return resp
}

// ListProject function
func ListProject() *Project {
	project := &Project{}
	err := GetDB().Table("projects").Find(project).Error
	if err != nil {
		return nil
	}
	return project
}

// ShowProject function
func ShowProject(id string) *Project {
	project := &Project{}
	err := GetDB().Table("projects").Where("id = ?", id).First(project).Error
	if err != nil {
		return nil
	}
	return project
}

// ModifyProject function
func ModifyProject(id string) *Project {
	project := &Project{}
	err := GetDB().Table("projects").Where("id = ?", id).Update(project).Error
	if err != nil {
		return nil
	}
	return project
}

// DeleteProject function
func DeleteProject(id string) {
	project := &Project{}
	GetDB().Table("projects").Where("id = ?", id).Delete(project)
}
