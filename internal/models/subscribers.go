package models

import "gorm.io/gorm"

type SubscriberInfo struct {
	gorm.Model
	CreatorID        string  `json:"creator-id"`
	CreatorEmail     string  `json:"creator-email"`
	SubscribtionType string  `json:"type"`
	SubscribtionName string  `json:"name"`
	Price            float32 `json:"price"`
}

type AllSubscribtion struct {
	gorm.Model
	SubscriberInfo
	UserID               int    `json:"user-id"`
	Username             string `json:"username" form:"username"`
	Email                string `json:"email" form:"email"`
	BillingStatusPaidOff bool   `json:"billing-status-paid-off"`
}
