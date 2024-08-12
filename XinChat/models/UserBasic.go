package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name string
	Password string
	Phone string
	Email string
	Identity string
	ClientIP string
	ClientPort string
	LoginTime time.Time
	HeartBeatTime time.Time
	IsLogout bool
	DeviceInfo string
}
