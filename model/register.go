package model

import "time"

type Register struct {
	ID        uint
	PatientID uint
	Price     float64
	DoctorID  uint
	CreateAt  time.Time
	UpdateAt  time.Time
}
