package model

import "github.com/jmoiron/sqlx"

func NewRepository(db *sqlx.DB) Repository {
	return &repository{DB: db}
}

type Repository interface {
	InsertUser(username string, password string, userType int64, parentID int64, doctorID int64) (err error)
	DeleteUserByID(id int64) (err error)
	QueryUserByID(id int64) (user *User, err error)
	QueryUserByName(username string) (user *User, err error)

	InsertPatient(patient Patient) (err error)
	UpdatePatientByID(patient Patient) (err error)
	DeletePatientByID(id int64) (err error)
	QueryPatientByID(id int64) (patient *Patient, err error)
	QueryPatientByName(name string) (patient *Patient, err error)

	InsertDoctor(doctor Doctor) (err error)
	UpdateDoctorByID(doctor Doctor) (err error)
	DeleteDoctorByID(id int64) (err error)
	QueryDoctorByID(id int64) (doctor *Doctor, err error)
	QueryDoctorByName(name string) (doctor *Doctor, err error)

	InsertDrug(drug Drugs) (err error)
	UpdateDrugByID(drug Drugs) (err error)
	DeleteDrugByID(id int64) (err error)
	QueryDrugByID(id int64) (drug *Drugs, err error)
	QueryDrugByName(name string) (drug *Drugs, err error)

	InsertPrescription(prescription Prescription) (err error)
	UpdatePrescriptionByID(prescription Prescription) (err error)
	DeletePrescriptionByID(id int64) (err error)
	QueryPrescriptionByID(id int64) (prescription *Prescription, err error)
	QueryPrescriptionByName(name string) (prescription *Prescription, err error)

	InsertRegister(register Register) (err error)
	UpdateRegister(register Register) (err error)
	DeleteRegisterByID(id int64) (err error)
	QueryRegisterByID(id int64) (register *Register, err error)
	QueryRegisterByName(name string) (register *Register, err error)
}

type repository struct {
	DB *sqlx.DB
}
