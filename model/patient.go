package model

import "time"

type Patient struct {
	ID int64 `db:"id"`
	Name string `db:"name"`
	Birthday time.Time `db:"birthday"`
	IsMarried bool `db:"is_married"`
	FamilyHistory string `db:"family_history"`
	PastMedicalHistory string `db:"past_medical_history"`
	CreateAt time.Time `db:"create_at"`
	UpdateAt	time.Time `db:"update_at"`
}