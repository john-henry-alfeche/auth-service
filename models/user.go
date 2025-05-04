package models

import (
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type User struct {
	Id       string `gorm:"type:char(27);primaryKey"`
	Username string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(255);not null"`
}

func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	if u.Id == "" {
		u.Id = ksuid.New().String()
	}
	return
}
