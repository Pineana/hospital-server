package model

import "time"

type User struct {
	ID        int64     `db:"id"`
	UserName  string    `db:"username"`
	Password  string    `db:"password"`
	UserType  int64     `db:"user_type"`
	PatientID int64     `db:"patient_id"`
	DoctorID  int64     `db:"doctor_id"`
	CreateAt  time.Time `db:"create_at"`
	UpdateAt  time.Time `db:"update_at"`
}
