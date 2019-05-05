package models

import (
	common "../common"
	uuid "github.com/satori/go.uuid"
)

// DashboardConfig struct
type DashboardConfig struct {
	Base
	UserID   uuid.UUID `gorm:"type:char(36);" json:"user_id"`
	Language string    `sql:"size:60;DEFAULT:'eng'"`
	Timezone string    `sql:"size:60;DEFAULT:'Etc/GMT'"`
}

// ValidateDashboardConfig function
func (dashboard *DashboardConfig) ValidateDashboardConfig() (map[string]interface{}, bool) {

	return common.Message(false, "Requirement passed"), true
}

// CreateDashboardConfig function
func (dashboard *DashboardConfig) CreateDashboardConfig() map[string]interface{} {

	if resp, ok := dashboard.ValidateDashboardConfig(); !ok {
		return resp
	}

	GetDB().Create(dashboard)

	resp := common.Message(true, "success")
	//resp["auth"] = tk
	resp["dashboard"] = dashboard
	return resp
}

// ListDashboardConfig function
func ListDashboardConfig() []DashboardConfig {
	dashboardConfig := []DashboardConfig{}
	GetDB().Find(&dashboardConfig)
	return dashboardConfig
}

// ShowDashboardConfig function
func ShowDashboardConfig(id string) *DashboardConfig {
	dashboard := &DashboardConfig{}
	err := GetDB().Table("dashboards").Where("id = ?", id).First(dashboard).Error
	if err != nil {
		return nil
	}
	return dashboard
}

// ModifyDashboardConfig function
func (dashboard *DashboardConfig) ModifyDashboardConfig(id string) {
	err := GetDB().Table("dashboards").Where("id = ?", id).Update(dashboard).Error
	if err != nil {
		println(err)
		return
	}
}

// DeleteDashboardConfig function
func DeleteDashboardConfig(id string) {
	dashboard := &DashboardConfig{}
	GetDB().Table("dashboards").Where("id = ?", id).Delete(dashboard)
}
