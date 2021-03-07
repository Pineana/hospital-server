package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func (r *repository) InsertUser(username string, password string, userType int64, parentID int64, doctorID int64) (err error) {
	sqlStr := `insert into user (username,password,user_type,patient_id,doctor_id)value (?,?,?,?,?)`
	_, err = r.DB.Exec(sqlStr, username, password, userType, parentID, doctorID)
	if err != nil {
		fmt.Println("InsertUser", err)
		return err
	}
	return
}

func (r *repository) DeleteUserByID(id int64) (err error) {
	sqlStr := `delete from user where id = ?`
	_, err = r.DB.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("DeleteUserByID", err)
		return err
	}
	return
}

func (r *repository) QueryUserByName(username string) (user *User, err error) {
	user = &User{}
	sqlStr := `select id,username,password,user_type,patient_id,doctor_id from user where username = ?`
	err = r.DB.Get(user, sqlStr, username)
	if err != nil {
		fmt.Println("QueryUserByName ", err)
		return nil, err
	}
	return
}

func (r *repository) QueryUserByID(id int64) (user *User, err error) {
	user = &User{}
	sqlStr := `select id,username,password,user_type,patient_id,doctor_id from user where id = ?`
	err = r.DB.Get(user, sqlStr, id)
	if err != nil {
		fmt.Println("QueryUserByID", err)
		return nil, err
	}
	return
}

func (r *repository) InsertPatient(patient Patient) (err error) {
	sqlStr := `insert into patient (name,birthday,is_married,family_history,past_medical_history)value (?,?,?,?,?,?)`
	_, err = r.DB.Exec(sqlStr, patient.Name, patient.Birthday, patient.IsMarried, patient.FamilyHistory, patient.PastMedicalHistory)
	if err != nil {
		fmt.Println("InsertPatient", err)
		return err
	}
	return
}

func (r *repository) UpdatePatientByID(patient Patient) (err error) {
	sqlStr := `update patient set name=?,birthday=?,is_married=?,family_history=?,past_medical_history=? where id = ?`
	_, err = r.DB.Exec(sqlStr, patient.Name, patient.Birthday, patient.IsMarried, patient.FamilyHistory, patient.PastMedicalHistory, patient.ID)
	if err != nil {
		fmt.Println("UpdatePatientByID", err)
		return err
	}
	return
}

func (r *repository) DeletePatientByID(id int64) (err error) {
	sqlStr := `delete from patient where id = ?`
	_, err = r.DB.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("DeletePatientByID", err)
		return err
	}
	return
}

func (r *repository) QueryPatientByID(id int64) (patient *Patient, err error) {
	patient = &Patient{}
	sqlStr := `select id,name,birthday,is_married,family_history,past_medical_history from patient where id = ?`
	err = r.DB.Get(patient, sqlStr, id)
	if err != nil {
		fmt.Println("QueryPatientByID", err)
		return nil, err
	}
	return
}

func (r *repository) QueryPatientByName(name string) (patient *Patient, err error) {
	patient = &Patient{}
	sqlStr := `select id,name,birthday,is_married,family_history,past_medical_history from patient where name = ?`
	err = r.DB.Get(patient, sqlStr, name)
	if err != nil {
		fmt.Println("QueryPatientByName", err)
		return nil, err
	}
	return
}

func (r *repository) InsertDoctor(doctor Doctor) (err error) {
	sqlStr := `insert into patient (name,birthday,is_married,family_history,past_medical_history)value (?,?,?,?,?,?)`
	_, err = r.DB.Exec(sqlStr, doctor.Name, doctor.Birthday, patient.IsMarried, patient.FamilyHistory, patient.PastMedicalHistory)
	if err != nil {
		fmt.Println("InsertPatient", err)
		return err
	}
	return
}

func (r *repository) UpdateDoctorByID(doctor Doctor) (err error) {
	sqlStr := `update patient set name=?,birthday=?,is_married=?,family_history=?,past_medical_history=? where id = ?`
	_, err = r.DB.Exec(sqlStr, patient.Name, patient.Birthday, patient.IsMarried, patient.FamilyHistory, patient.PastMedicalHistory, patient.ID)
	if err != nil {
		fmt.Println("UpdatePatientByID", err)
		return err
	}
	return
}

func (r *repository) DeletePatientByID(id int64) (err error) {
	sqlStr := `delete from patient where id = ?`
	_, err = r.DB.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("DeletePatientByID", err)
		return err
	}
	return
}

func (r *repository) QueryPatientByID(id int64) (patient *Patient, err error) {
	patient = &Patient{}
	sqlStr := `select id,name,birthday,is_married,family_history,past_medical_history from patient where id = ?`
	err = r.DB.Get(patient, sqlStr, id)
	if err != nil {
		fmt.Println("QueryPatientByID", err)
		return nil, err
	}
	return
}

func (r *repository) QueryPatientByName(name string) (patient *Patient, err error) {
	patient = &Patient{}
	sqlStr := `select id,name,birthday,is_married,family_history,past_medical_history from patient where name = ?`
	err = r.DB.Get(patient, sqlStr, name)
	if err != nil {
		fmt.Println("QueryPatientByName", err)
		return nil, err
	}
	return
}
