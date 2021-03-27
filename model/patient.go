package model

type Patient struct {
	Id                 int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"` // 主键
	Name               string `gorm:"column:name;NOT NULL" json:"name"`               // 姓名
	Year               string `gorm:"column:year;NOT NULL" json:"year"`
	Sex                string `gorm:"column:sex;NOT NULL" json:"sex"` // 性别
	Phone              string `gorm:"column:phone;NOT NULL" json:"phone"`
	PastMedicalHistory string `gorm:"column:past_medical_history" json:"pastMedicalHistory"` // 既往病史
}
