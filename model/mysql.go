package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func (r *repository) InsertUser(user User) (err error) {
	result := r.DB.Create(&user)
	if result.Error != nil {
		fmt.Println("InsertUser", err)
		return err
	}
	return nil
}

func (r *repository) DeleteUserByID(user User) (err error) {
	result := r.DB.Delete(&user)
	if result.Error != nil {
		fmt.Println("DeleteUserByID", err)
		return err
	}
	return
}

func (r *repository) QueryUserByName(username string) (user *User, err error) {
	result := r.DB.Where("user_name = ?", username).First(&user)
	if result.Error != nil {
		fmt.Println("QueryUserByName", err)
		return nil, err
	}
	return
}

func (r *repository) QueryUserByID(id int64) (user *User, err error) {
	result := r.DB.First(user, id)
	if result.Error != nil {
		fmt.Println("QueryUserByID", err)
		return nil, err
	}
	return
}

func (r *repository) QueryUserList(page int, size int) (users []*User, err error) {
	users = make([]*User, 0)
	result := r.DB.Find(&users).Limit(size).Offset(page * size)
	if result.Error != nil {
		return nil, err
	}
	return
}

func (r *repository) InsertPatient(patient Patient) (err error) {
	result := r.DB.Create(&patient)
	if result.Error != nil {
		fmt.Println("InsertPatient", err)
		return err
	}
	return nil
}

func (r *repository) DeletePatientByID(patient Patient) (err error) {
	result := r.DB.Delete(&patient)
	if result.Error != nil {
		fmt.Println("DeletePatientByID", err)
		return err
	}
	return
}

func (r *repository) UpdatePatientByID(patient Patient) (err error) {
	result := r.DB.Save(&patient)
	if result.Error != nil {
		fmt.Println("UpdatePatientByID", err)
		return err
	}
	return
}

func (r *repository) QueryPatientByName(name string) (patients []*Patient, err error) {
	result := r.DB.Where("name like ?", "%"+name+"%").Find(&patients)
	if result.Error != nil {
		fmt.Println("QueryPatientByName", err)
		return nil, err
	}
	return
}

func (r *repository) QueryPatientList(page int, size int) (patients []*Patient, err error) {
	patients = make([]*Patient, 0)
	result := r.DB.Limit(size).Offset((page - 1) * size).Find(&patients)
	if result.Error != nil {
		return nil, err
	}
	return
}

func (r *repository) CountPatientList(size int) (totalNum int, totalPage int) {
	var count int64
	result := r.DB.Find(&Patient{}).Count(&count)
	if result.Error != nil {
		return 0, 0
	}
	return int(count), int(count)/size + 1
}

func (r *repository) InsertDoctor(doctor Doctor) (err error) {
	result := r.DB.Create(&doctor)
	if result.Error != nil {
		fmt.Println("InsertDoctor", err)
		return err
	}
	return nil
}

func (r *repository) DeleteDoctorByID(doctor Doctor) (err error) {
	result := r.DB.Delete(&doctor)
	if result.Error != nil {
		fmt.Println("DeleteDoctorByID", err)
		return err
	}
	return
}

func (r *repository) UpdateDoctorByID(doctor Doctor) (err error) {
	result := r.DB.Save(&doctor)
	if result.Error != nil {
		fmt.Println("UpdateDoctorByID", err)
		return err
	}
	return
}

func (r *repository) QueryDoctorByName(name string, page, size int) (doctors []*Doctor, err error) {
	result := r.DB.Where("name like ?", "%"+name+"%").Limit(size).Offset((page - 1) * size).Find(&doctors)
	if result.Error != nil {
		fmt.Println("QueryDoctorByName", err)
		return nil, err
	}
	return
}

func (r *repository) CountDoctorByName(name string, size int) (totalNum int, totalPage int) {
	var count int64
	result := r.DB.Where("name like ?", "%"+name+"%").Find(&Doctor{}).Count(&count)
	if result.Error != nil {
		return 0, 0
	}
	return int(count), int(count)/size + 1
}

func (r *repository) QueryDoctorList(page int, size int) (doctors []*Doctor, err error) {
	doctors = make([]*Doctor, 0)
	result := r.DB.Limit(size).Offset((page - 1) * size).Find(&doctors)
	if result.Error != nil {
		return nil, err
	}
	return
}

func (r *repository) CountDoctorList(size int) (totalNum int, totalPage int) {
	var count int64
	result := r.DB.Find(&Doctor{}).Count(&count)
	if result.Error != nil {
		return 0, 0
	}
	return int(count), int(count)/size + 1
}

func (r *repository) InsertRegister(register Register) (err error) {
	result := r.DB.Create(&register)
	if result.Error != nil {
		fmt.Println("InsertRegister", err)
		return err
	}
	return nil
}

func (r *repository) DeleteRegisterByID(register Register) (err error) {
	result := r.DB.Delete(&register)
	if result.Error != nil {
		fmt.Println("DeleteRegisterByID", err)
		return err
	}
	return
}

func (r *repository) UpdateRegisterByID(register Register) (err error) {
	result := r.DB.Save(&register)
	if result.Error != nil {
		fmt.Println("UpdateRegisterByID", err)
		return err
	}
	return
}

func (r *repository) CountRegisterBeforeID(id int) (beforeNum int, err error) {
	var count int64
	err = r.DB.Table("registers").Select("count(*)").Where("id < ?", id).Count(&count).Error
	if err != nil {
		fmt.Println("CountRegisterBeforeID", err)
		return 0, err
	}
	return int(count), nil
}

func (r *repository) QueryRegisterByID(id int) (register *RegisterResult, err error) {
	res := r.DB.Table("registers").Select("registers.id,patients.name as patient_name,patients.year as patient_year,patients.phone as patient_phone,patients.sex as patient_sex,price,CONCAT(doctors.name,'--',doctors.department) as doctor_name").Joins("LEFT JOIN patients ON patients.id = registers.patient_id").Joins("LEFT JOIN doctors ON doctors.id = registers.doctor_id").Where("registers.id = ?", id).Find(&register)
	if res.Error != nil {
		fmt.Println("QueryRegisterByID", err)
		return nil, err
	}
	return
}

func (r *repository) QueryRegisterList(page int, size int) (result []*RegisterResult, err error) {
	result = make([]*RegisterResult, 0)
	res := r.DB.Table("registers").Select("registers.id,patients.name as patient_name,patients.year as patient_year,patients.phone as patient_phone,patients.sex as patient_sex,price,CONCAT(doctors.name,'--',doctors.department) as doctor_name").Joins("LEFT JOIN patients ON patients.id = registers.patient_id").Joins("LEFT JOIN doctors ON doctors.id = registers.doctor_id").Limit(size).Offset((page - 1) * size).Scan(&result)
	if res.Error != nil {
		return nil, err
	}
	return
}

func (r *repository) CountRegisterList(size int) (totalNum int, totalPage int) {
	var count int64
	result := r.DB.Find(&Register{}).Count(&count)
	if result.Error != nil {
		return 0, 0
	}
	return int(count), int(count)/size + 1
}

func (r *repository) InsertDrug(drug Drug) (err error) {
	result := r.DB.Create(&drug)
	if result.Error != nil {
		fmt.Println("InsertDrug", err)
		return err
	}
	return nil
}

func (r *repository) DeleteDrugByID(drug Drug) (err error) {
	result := r.DB.Delete(&drug)
	if result.Error != nil {
		fmt.Println("DeleteDrugByID", err)
		return err
	}
	return
}

func (r *repository) UpdateDrugByID(drug Drug) (err error) {
	result := r.DB.Save(&drug)
	if result.Error != nil {
		fmt.Println("UpdateDrugByID", err)
		return err
	}
	return
}

func (r *repository) QueryDrugByName(name string, page, size int) (drug []*Drug, err error) {
	result := r.DB.Where("name like ?", "%"+name+"%").Limit(size).Offset((page - 1) * size).Find(&drug)
	if result.Error != nil {
		fmt.Println("QueryDrugByName", err)
		return nil, err
	}
	return
}

func (r *repository) CountDrugByName(name string, size int) (totalNum int, totalPage int) {
	var count int64
	result := r.DB.Where("name like ?", "%"+name+"%").Find(&Drug{}).Count(&count)
	if result.Error != nil {
		return 0, 0
	}
	return int(count), int(count)/size + 1
}

func (r *repository) QueryDrugList(page int, size int) (drugs []*Drug, err error) {
	drugs = make([]*Drug, 0)
	result := r.DB.Limit(size).Offset((page - 1) * size).Find(&drugs)
	if result.Error != nil {
		return nil, err
	}
	return
}

func (r *repository) CountDrugList(size int) (totalNum int, totalPage int) {
	var count int64
	result := r.DB.Find(&Drug{}).Count(&count)
	if result.Error != nil {
		return 0, 0
	}
	return int(count), int(count)/size + 1
}

func (r *repository) InsertPrescription(prescription Prescription) (err error) {
	result := r.DB.Create(&prescription)
	if result.Error != nil {
		fmt.Println("InsertPrescription", err)
		return err
	}
	return nil
}

func (r *repository) DeletePrescriptionByID(prescription Prescription) (err error) {
	result := r.DB.Delete(&prescription)
	if result.Error != nil {
		fmt.Println("DeletePrescriptionByID", err)
		return err
	}
	return
}

func (r *repository) UpdatePrescriptionByID(prescription Prescription) (err error) {
	result := r.DB.Save(&prescription)
	if result.Error != nil {
		fmt.Println("UpdatePrescriptionByID", err)
		return err
	}
	return
}

func (r *repository) QueryPrescriptionByID(id int) (prescription *Prescription, err error) {
	result := r.DB.First(prescription, id)
	if result.Error != nil {
		fmt.Println("QueryPrescriptionByID", err)
		return nil, err
	}
	return
}

func (r *repository) QueryPrescriptionList(page int, size int) (prescriptions []*Prescription, err error) {
	prescriptions = make([]*Prescription, 0)
	result := r.DB.Find(&prescriptions).Limit(size).Offset(page * size)
	if result.Error != nil {
		return nil, err
	}
	return
}

func (r *repository) CountPrescriptionList(size int) (totalNum int, totalPage int) {
	var count int64
	result := r.DB.Find(&Prescription{}).Count(&count)
	if result.Error != nil {
		return 0, 0
	}
	return int(count), int(count)/size + 1
}
