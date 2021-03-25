package model

type Doctor struct {
	Id           int    `gorm:"column:id;primary_key" json:"id"`
	Name         string `gorm:"column:name;NOT NULL" json:"name"`
	Year         int    `gorm:"column:year;NOT NULL" json:"year"`
	DepartmentId int    `gorm:"column:department_id;NOT NULL" json:"department_id"`
	Phone        string `gorm:"column:phone;NOT NULL" json:"phone"`
	Sex          int    `gorm:"column:sex;NOT NULL" json:"sex"`
}
