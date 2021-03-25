package model

import (
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

type Repository interface {
	InsertUser(user User) (err error)
	DeleteUserByID(user User) (err error)
	QueryUserByName(username string) (user *User, err error)
	QueryUserByID(id int64) (user *User, err error)
	QueryUserList(page int, size int) (users []*User, err error)
	InsertPatient(patient Patient) (err error)
	DeletePatientByID(patient Patient) (err error)
	QueryPatientByName(name string) (user *Patient, err error)
	UpdatePatientByID(patient Patient) (err error)
	QueryPatientList(page int, size int) (patients []*Patient, err error)
	CountPatientList(size int) (totalNum int, totalPage int)

	InsertDoctor(doctor Doctor) (err error)
	DeleteDoctorByID(doctor Doctor) (err error)
	QueryDoctorByName(name string) (user *Doctor, err error)
	UpdateDoctorByID(doctor Doctor) (err error)
	QueryDoctorList(page int, size int) (doctors []*Doctor, err error)
	CountDoctorList(size int) (totalNum int, totalPage int)

	InsertRegister(register Register) (err error)
	DeleteRegisterByID(register Register) (err error)
	QueryRegisterByID(id int64) (register *Register, err error)
	UpdateRegisterByID(register Register) (err error)
	QueryRegisterList(page int, size int) (registers []*Register, err error)
	CountRegisterList(size int) (totalNum int, totalPage int)

	InsertDrug(drug Drug) (err error)
	DeleteDrugByID(drug Drug) (err error)
	QueryDrugByName(name string) (drug *Drug, err error)
	UpdateDrugByID(drug Drug) (err error)
	QueryDrugList(page int, size int) (drugs []*Drug, err error)
	CountDrugList(size int) (totalNum int, totalPage int)

	InsertPrescription(prescription Prescription) (err error)
	DeletePrescriptionByID(prescription Prescription) (err error)
	QueryPrescriptionByID(id int) (prescription *Prescription, err error)
	UpdatePrescriptionByID(prescription Prescription) (err error)
	QueryPrescriptionList(page int, size int) (prescriptions []*Prescription, err error)
	CountPrescriptionList(size int) (totalNum int, totalPage int)
}

type repository struct {
	DB *gorm.DB
}
