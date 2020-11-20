package gorm_usage

import "time"

type Company struct {
	Code   string  `gorm:"primaryKey;size:64;not null"`
	Name   string  `gorm:"size:64;not null"`
	Users  []*User `gorm:"foreignKey:CompanyID;references:Code"`
}

type User struct {
	Code    string  `gorm:"primaryKey;size:64;not null"`
	Name    string  `gorm:"size:64;not null"`
	Age     int64  `gorm:"size:8;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	CompanyID string  `gorm:"size:64;not null"`
	Company *Company `gorm:"foreignKey:CompanyID;references:Code"`
	//ProfileID string  `gorm:"size:64;not null"`
	//Profile *Profile `gorm:"foreignKey:ProfileID;references:Code"`
}

/*
type Profile struct {
	Code    string `gorm:"primaryKey;size:64;not null"`
	User    *User `gorm:"foreignKey:ProfileID;references:Code"`
}
*/
