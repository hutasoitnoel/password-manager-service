package models

import "gorm.io/gorm"

type IdentificationCard struct {
	gorm.Model
	UserId         uint   `gorm:"not null" json:"user_id"`
	IDNumber       string `gorm:"type:varchar(255)" json:"id_number"`
	Type           string `gorm:"type:varchar(50)" json:"type"`
	Name           string `gorm:"type:varchar(255)" json:"name"`
	Province       string `gorm:"type:varchar(100)" json:"province"`
	City           string `gorm:"type:varchar(100)" json:"city"`
	DOB            string `gorm:"type:varchar(50)" json:"dob"`
	Gender         string `gorm:"type:varchar(10)" json:"gender"`
	BloodType      string `gorm:"type:varchar(3)" json:"blood_type"`
	Address        string `gorm:"type:text" json:"address"`
	RtRw           string `gorm:"type:varchar(50)" json:"rt_rw"`
	SubDistrict    string `gorm:"type:varchar(100)" json:"sub_district"`
	District       string `gorm:"type:varchar(100)" json:"district"`
	Religion       string `gorm:"type:varchar(50)" json:"religion"`
	MaritalStatus  string `gorm:"type:varchar(50)" json:"marital_status"`
	Occupation     string `gorm:"type:varchar(100)" json:"occupation"`
	Citizenship    string `gorm:"type:varchar(50)" json:"citizenship"`
	ExpirationDate string `gorm:"type:varchar(50)" json:"expiration_date"`
	PlaceOfIssue   string `gorm:"type:varchar(100)" json:"place_of_issue"`
}
