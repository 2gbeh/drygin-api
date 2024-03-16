package models

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Company string `gorm:"type:varchar(50)"`
	LastName string `gorm:"type:varchar(50)"`
	FirstName string `gorm:"type:varchar(50)"`
	EmailAddress string `gorm:"type:varchar(50)"`
	JobTitle string `gorm:"type:varchar(50)"`
	BusinessPhone string `gorm:"type:varchar(25)"`
	HomePhone string `gorm:"type:varchar(25)"`
	MobilePhone string `gorm:"type:varchar(25)"`
	FaxNumber string `gorm:"type:varchar(25)"`
	Address string 
	City string `gorm:"type:varchar(50)"`
	StateProvince string `gorm:"type:varchar(50)"`
	ZipPostalCode string `gorm:"type:varchar(15)"`
	CountryRegion string `gorm:"type:varchar(50)"`
	WebPage string 
	Notes string 
	Attachments []byte
}

type Tags struct {
	Data []Tag `json:"data"`
}