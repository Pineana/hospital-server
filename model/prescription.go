package model

import "time"

type Prescription struct {
	ID        uint
	PatientID uint
	DrugID    uint
	CreateAt  time.Time
	UpdateAt  time.Time
}
