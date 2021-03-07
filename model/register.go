package model

import "time"

type Register struct {
	ID int64 `db:"id"`
	PatientID int64 `db:"patient_id"`
	Price float64 `db:"price"`
	DoctorID int64 `db:"doctor_id"`
	CreateAt time.Time `db:"create_at"`
	UpdateAt	time.Time `db:"update_at"`
}