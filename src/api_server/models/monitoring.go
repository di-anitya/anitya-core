package models

import (
	common "../common"
	uuid "github.com/satori/go.uuid"
)

// HTTPMonitoringConfig struct
type HTTPMonitoringConfig struct {
	Base
	ProjectID uuid.UUID `gorm:"type:char(36);" json:"project_id"`
	URL       string    `sql:"size:60" json:"url"`
	Duration  string    `sql:"size:60" json:"duration"`
}

// DNSMonitoringConfig struct
type DNSMonitoringConfig struct {
	Base
	ProjectID uuid.UUID `gorm:"type:char(36);" json:"project_id"`
	URL       string    `sql:"size:60" json:"url"`
	Duration  string    `sql:"size:60" json:"duration"`
}

// ValidateHTTPMonitoringConfig function
func (http_monitoring_configs *HTTPMonitoringConfig) ValidateHTTPMonitoringConfig() (map[string]interface{}, bool) {
	if http_monitoring_configs.URL == "" {
		return common.Message(false, "'url' is required"), false
	}

	if http_monitoring_configs.Duration == "" {
		return common.Message(false, "'duration' is required"), false
	}

	return common.Message(false, "Requirement passed"), true
}

// CreateHTTPMonitoringConfig function
func (http_monitoring_configs *HTTPMonitoringConfig) CreateHTTPMonitoringConfig() map[string]interface{} {

	if resp, ok := http_monitoring_configs.ValidateHTTPMonitoringConfig(); !ok {
		return resp
	}

	GetDB().Create(http_monitoring_configs)

	resp := common.Message(true, "success")
	//resp["auth"] = tk
	resp["monitoring_http"] = http_monitoring_configs
	return resp
}

// ListHTTPMonitoringConfig function
func ListHTTPMonitoringConfig() *HTTPMonitoringConfig {
	httpMonitoringConfigs := &HTTPMonitoringConfig{}
	err := GetDB().Table("http_monitoring_configs").Find(httpMonitoringConfigs).Error
	if err != nil {
		return nil
	}
	return httpMonitoringConfigs
}

// ShowHTTPMonitoringConfig function
func ShowHTTPMonitoringConfig(id string) *HTTPMonitoringConfig {
	httpMonitoringConfigs := &HTTPMonitoringConfig{}
	err := GetDB().Table("http_monitoring_configs").Where("id = ?", id).First(httpMonitoringConfigs).Error
	if err != nil {
		return nil
	}
	return httpMonitoringConfigs
}

// ModifyHTTPMonitoringConfig function
func ModifyHTTPMonitoringConfig(id string) *HTTPMonitoringConfig {
	httpMonitoringConfigs := &HTTPMonitoringConfig{}
	err := GetDB().Table("http_monitoring_configs").Where("id = ?", id).Update(httpMonitoringConfigs).Error
	if err != nil {
		return nil
	}
	return httpMonitoringConfigs
}

// DeleteHTTPMonitoringConfig function
func DeleteHTTPMonitoringConfig(id string) {
	httpMonitoringConfigs := &HTTPMonitoringConfig{}
	GetDB().Table("http_monitoring_configs").Where("id = ?", id).Delete(httpMonitoringConfigs)
}

// ValidateDNSMonitoringConfig function
func (dns_monitoring_configs *DNSMonitoringConfig) ValidateDNSMonitoringConfig() (map[string]interface{}, bool) {
	if dns_monitoring_configs.URL == "" {
		return common.Message(false, "'url' is required"), false
	}

	if dns_monitoring_configs.Duration == "" {
		return common.Message(false, "'duration' is required"), false
	}

	return common.Message(false, "Requirement passed"), true
}

// CreateDNSMonitoringConfig function
func (dns_monitoring_configs *DNSMonitoringConfig) CreateDNSMonitoringConfig() map[string]interface{} {

	if resp, ok := dns_monitoring_configs.ValidateDNSMonitoringConfig(); !ok {
		return resp
	}

	GetDB().Create(dns_monitoring_configs)

	resp := common.Message(true, "success")
	//resp["auth"] = tk
	resp["monitoring_dns"] = dns_monitoring_configs
	return resp
}

// ListDNSMonitoringConfig function
func ListDNSMonitoringConfig() *DNSMonitoringConfig {
	dnsMonitoringConfigs := &DNSMonitoringConfig{}
	err := GetDB().Table("dns_monitoring_configs").Find(dnsMonitoringConfigs).Error
	if err != nil {
		return nil
	}
	return dnsMonitoringConfigs
}

// ShowDNSMonitoringConfig function
func ShowDNSMonitoringConfig(id string) *DNSMonitoringConfig {
	dnsMonitoringConfigs := &DNSMonitoringConfig{}
	err := GetDB().Table("dns_monitoring_configs").Where("id = ?", id).First(dnsMonitoringConfigs).Error
	if err != nil {
		return nil
	}
	return dnsMonitoringConfigs
}

// ModifyDNSMonitoringConfig function
func ModifyDNSMonitoringConfig(id string) *DNSMonitoringConfig {
	dnsMonitoringConfigs := &DNSMonitoringConfig{}
	err := GetDB().Table("dns_monitoring_configs").Where("id = ?", id).Update(dnsMonitoringConfigs).Error
	if err != nil {
		return nil
	}
	return dnsMonitoringConfigs
}

// DeleteDNSMonitoringConfig function
func DeleteDNSMonitoringConfig(id string) {
	dnsMonitoringConfigs := &DNSMonitoringConfig{}
	GetDB().Table("dns_monitoring_configs").Where("id = ?", id).Delete(dnsMonitoringConfigs)
}
