package models

import "gorm.io/gorm"

type CreatorsMediaActivities struct {
	gorm.Model
	FileName string `json:"file-name"`
	FileExtension string `json:"file-extension"`
	FileType string `json:"file-type"`
	FilePath string `json:"file-path"`
	SenderID string `json:"sender-id"`
	SenderEmail string `json:"sender-email"`
	SpecialSubscriberStatus bool `json:"special-subscriber-status"`
	SubscriberID string `json:"subscriber-id"`
}