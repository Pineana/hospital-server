package model

type Doctor struct {
	Id         int    `gorm:"column:id;primary_key" json:"id"`
	Name       string `gorm:"column:name;NOT NULL" json:"name"`
	Year       string `gorm:"column:year;NOT NULL" json:"year"`
	Department string `gorm:"column:department;NOT NULL" json:"department"`
	Phone      string `gorm:"column:phone;NOT NULL" json:"phone"`
	Sex        string `gorm:"column:sex;NOT NULL" json:"sex"`
}
