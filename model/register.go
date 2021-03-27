package model

type Register struct {
	Id        int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	PatientId int    `gorm:"column:patient_id;NOT NULL" json:"patientId"`
	Price     string `gorm:"column:price;default:1;NOT NULL" json:"price"`
	DoctorId  int    `gorm:"column:doctor_id;default:0;NOT NULL" json:"doctorId"`
}

type RegisterResult struct {
	Id               int    `json:"id"`
	PatientName      string `json:"patientName"`
	PatientYear      string `json:"patientYear"`
	PatientPhone     string `json:"patientPhone"`
	PatientSex       string `json:"patientSex"`
	Price            string `json:"price"`
	DoctorName       string `json:"doctorName"`
	DoctorDepartment string `json:"doctorDepartment"`
}
