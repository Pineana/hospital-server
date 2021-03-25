package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;type:int(64);AUTO_INCREMENT"`
	UserName  string    `gorm:"not null;type:varchar(256)"`
	Password  string    `gorm:"not null;type:varchar(256)"`
	UserType  uint8     `gorm:"not null;type:tinyint(1)"`
	PatientID uint      `gorm:"not null;type:int(64)"`
	DoctorID  uint      `gorm:"not null;type:int(64)"`
	CreateAt  time.Time `gorm:"autoCreateTime;type:timestamp"`
	UpdateAt  time.Time `gorm:"autoUpdateTime;type:timestamp"`
}
