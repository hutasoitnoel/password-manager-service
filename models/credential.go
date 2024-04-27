package models

import "gorm.io/gorm"

type Credential struct {
	gorm.Model
	UserId      uint   `gorm:"notNull" json:"user_id"`
	WebsiteUrl  string `gorm:"notNull" json:"website_url"`
	WebsiteName string `gorm:"notNull" json:"website_name"`
	Username    string `gorm:"notNull" json:"username"`
	Password    string `gorm:"notNull" json:"password"`
}
