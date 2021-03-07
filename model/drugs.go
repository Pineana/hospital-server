package model

import "time"

type Drugs struct {
	ID int64 `db:"id"`
	Name string `db:"name"`
	Stock int64 `db:"stock"`
	Price int64 `db:"price"`
	Type int64	`db:"type"`
	Description string `db:"description"`
	CreateAt time.Time `db:"create_at"`
	UpdateAt	time.Time `db:"update_at"`
}