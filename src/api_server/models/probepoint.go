package models

import (
	common "../common"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// ProbePoint struct
type ProbePoint struct {
	Base
	ProjectID uuid.UUID `gorm:"type:char(36);" json:"project_id"`
	Name      int64     `sql:"size:60" json:"name"`
	IPAddress string    `sql:"size:60" json:"ip_address"`
}

// ValidateProbePoint function
func (probepoint *ProbePoint) ValidateProbePoint() (map[string]interface{}, bool) {
	// Email must be unique
	temp := &ProbePoint{}

	// check for errors and duplicate emails
	err := GetDB().Table("probepoints").Where("name = ?", probepoint.Name).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return common.Message(false, "Connection error. Please retry"), false
	}

	return common.Message(false, "Requirement passed"), true
}

// CreateProbePoint function
func (probepoint *ProbePoint) CreateProbePoint() map[string]interface{} {

	if resp, ok := probepoint.ValidateProbePoint(); !ok {
		return resp
	}

	GetDB().Create(probepoint)

	resp := common.Message(true, "success")
	//resp["auth"] = tk
	resp["probepoint"] = probepoint
	return resp
}

// ListProbePoint function
func ListProbePoint() []ProbePoint {
	probepoint := []ProbePoint{}
	GetDB().Find(&probepoint)
	return probepoint
}

// ShowProbePoint function
func ShowProbePoint(id string) *ProbePoint {
	probepoint := &ProbePoint{}
	err := GetDB().Table("probepoints").Where("id = ?", id).First(probepoint).Error
	if err != nil {
		return nil
	}
	return probepoint
}

// ModifyProbePoint function
func (probepoint *ProbePoint) ModifyProbePoint(id string) {
	err := GetDB().Table("probepoints").Where("id = ?", id).Update(probepoint).Error
	if err != nil {
		println(err)
		return
	}
}

// DeleteProbePoint function
func DeleteProbePoint(id string) {
	probepoint := &ProbePoint{}
	GetDB().Table("probepoints").Where("id = ?", id).Delete(probepoint)
}
