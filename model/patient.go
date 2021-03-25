package model

type Patient struct {
	Id                 int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"` // 主键
	Name               string `gorm:"column:name;NOT NULL" json:"name"`               // 姓名
	Year               int    `gorm:"column:year;NOT NULL" json:"year"`
	IsMarried          bool   `gorm:"column:is_married;NOT NULL" json:"is_married"`            // 婚姻状况
	FamilyHistory      string `gorm:"column:family_history" json:"family_history"`             // 家族病史
	PastMedicalHistory string `gorm:"column:past_medical_history" json:"past_medical_history"` // 既往病史
}
