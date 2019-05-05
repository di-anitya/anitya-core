package models

import (
	common "../common"
	uuid "github.com/satori/go.uuid"
)

// NotificationSlack struct
type NotificationSlack struct {
	Base
	Text      string    `json:"text"`
	Username  string    `json:"username"`
	IconEmoji string    `json:"icon_emoji"`
	IconURL   string    `json:"icon_url"`
	Channel   string    `json:"channel"`
	ProjectID uuid.UUID `gorm:"type:char(36);" json:"project_id"`
}

// NotificationHistory struct
type NotificationHistory struct {
	Base
	ProjectID uuid.UUID `gorm:"type:char(36);" json:"project_id"`
}

//////////////////////////////////////////////////////////////////////////

// ValidateNotificationSlack function
func (notification_slack *NotificationSlack) ValidateNotificationSlack() (map[string]interface{}, bool) {

	return common.Message(false, "Requirement passed"), true
}

// CreateNotificationSlack function
func (notification_slack *NotificationSlack) CreateNotificationSlack() map[string]interface{} {

	if resp, ok := notification_slack.ValidateNotificationSlack(); !ok {
		return resp
	}

	GetDB().Create(notification_slack)

	resp := common.Message(true, "success")
	//resp["auth"] = tk
	resp["notification_slack"] = notification_slack
	return resp
}

// ListNotificationSlack function
func ListNotificationSlack() []NotificationSlack {
	notificationSlack := []NotificationSlack{}
	GetDB().Find(&notificationSlack)
	return notificationSlack
}

// ShowNotificationSlack function
func ShowNotificationSlack(id string) *NotificationSlack {
	notificationSlack := &NotificationSlack{}
	err := GetDB().Table("notification_slacks").Where("id = ?", id).First(notificationSlack).Error
	if err != nil {
		return nil
	}
	return notificationSlack
}

// ModifyNotificationSlack function
func (notification_slack *NotificationSlack) ModifyNotificationSlack(id string) {
	err := GetDB().Table("notification_slacks").Where("id = ?", id).Update(notification_slack).Error
	if err != nil {
		println(err)
		return
	}
}

// DeleteNotificationSlack function
func DeleteNotificationSlack(id string) {
	notificationSlack := &NotificationSlack{}
	GetDB().Table("notification_slacks").Where("id = ?", id).Delete(notificationSlack)
}

//////////////////////////////////////////////////////////////////////////

// ValidateNotificationHistory function
func (notification_history *NotificationHistory) ValidateNotificationHistory() (map[string]interface{}, bool) {

	return common.Message(false, "Requirement passed"), true
}

// CreateNotificationHistory function
func (notification_history *NotificationHistory) CreateNotificationHistory() map[string]interface{} {

	if resp, ok := notification_history.ValidateNotificationHistory(); !ok {
		return resp
	}

	GetDB().Create(notification_history)

	resp := common.Message(true, "success")
	//resp["auth"] = tk
	resp["notification_history"] = notification_history
	return resp
}

// ListNotificationHistory function
func ListNotificationHistory() []NotificationHistory {
	notificationHistory := []NotificationHistory{}
	GetDB().Find(&notificationHistory)
	return notificationHistory
}

// ShowNotificationHistory function
func ShowNotificationHistory(id string) *NotificationHistory {
	notificationHistory := &NotificationHistory{}
	err := GetDB().Table("notification_histories").Where("id = ?", id).First(notificationHistory).Error
	if err != nil {
		return nil
	}
	return notificationHistory
}

// ModifyNotificationHistory function
func (notification_history *NotificationHistory) ModifyNotificationHistory(id string) {
	err := GetDB().Table("notification_histories").Where("id = ?", id).Update(notification_history).Error
	if err != nil {
		println(err)
		return
	}
}

// DeleteNotificationHistory function
func DeleteNotificationHistory(id string) {
	notificationHistory := &NotificationHistory{}
	GetDB().Table("notification_histories").Where("id = ?", id).Delete(notificationHistory)
}
