package model

import "time"

type Doctor struct {
	ID int64 `db:"id"`
	Name string `db:"name"`
	Year int64	`db:"year"`
	DepartmentID int64 `db:"department_id"`
	Phone int64 `db:"phone"`
	Sex bool `db:"sex"`
	CreateAt time.Time `db:"create_at"`
	UpdateAt	time.Time `db:"update_at"`
}