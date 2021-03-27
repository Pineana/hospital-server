package model

import "time"

type Prescription struct {
	ID        int
	PatientID int
	DrugID    int
	CreateAt  time.Time
	UpdateAt  time.Time
}
