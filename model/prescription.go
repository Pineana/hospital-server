package model

import "time"

type Prescription struct {
	ID int64 `db:"id"`
	PatientID int64 `db:"patient_id"`
	DrugID	int64  `db:"drug_id"`
	CreateAt time.Time `db:"create_at"`
	UpdateAt	time.Time `db:"update_at"`
}