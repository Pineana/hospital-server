package usecase

import (
	"gorm.io/gorm"
	"hospital-admin/model"
	"hospital-admin/utils"
)

func NewUseCase(db *gorm.DB) UseCase {
	return &userCase{
		Repo: model.NewRepository(db),
	}
}

type userCase struct {
	Repo model.Repository
}

type UseCase interface {
	UserLogin(username string, password string) (ok bool, err error)
	UserRegister(user model.User) (ok bool, err error)
	// 获取患者列表
	GetPatientList(page int, size int) (patient []*model.Patient, totalNum int, totalPage int, err error)
	// 插入患者
	AddPatient(data model.Patient) (err error)
	// 删除患者
	DeletePatient(data model.Patient) (err error)
	// 编辑患者
	EditPatient(data model.Patient) (err error)
	// 查询患者
	QueryPatient(name string) (patients []*model.Patient, err error)

	// 插入医生
	AddDoctor(data model.Doctor) (err error)
	// 删除医生
	DeleteDoctor(data model.Doctor) (err error)
	// 编辑医生
	EditDoctor(data model.Doctor) (err error)
	// 查询医生
	QueryDoctor(name string, page, size int) (doctors []*model.Doctor, totalNum int, totalPage int, err error)
	// 获取医生列表
	GetDoctorList(page int, size int) (doctors []*model.Doctor, totalNum int, totalPage int, err error)

	// 插入药品
	AddDrug(data model.Drug) (err error)
	// 删除药品
	DeleteDrug(data model.Drug) (err error)
	// 编辑药品
	EditDrug(data model.Drug) (err error)
	// 查询药品
	QueryDrug(name string, page, size int) (drugs []*model.Drug, totalNum, totalPage int, err error)
	// 获取药品列表
	GetDrugList(page int, size int) (drugs []*model.Drug, totalNum int, totalPage int, err error)

	// 插入挂号
	AddRegister(data model.Register) (err error)
	// 删除挂号
	DeleteRegister(data model.Register) (err error)
	// 获取挂号列表
	GetRegisterList(page int, size int) (result []*model.RegisterResult, totalNum int, totalPage int, err error)

	// 插入药品
	AddPrescription(data model.Prescription) (err error)
	// 删除药品
	DeletePrescription(data model.Prescription) (err error)
	// 编辑药品
	EditPrescription(data model.Prescription) (err error)
	// 查询药品
	QueryPrescription(id int) (prescription *model.Prescription, err error)
	// 获取药品列表
	GetPrescriptionList(page int, size int) (prescriptions []*model.Prescription, totalNum int, totalPage int, err error)
}

func (u *userCase) UserLogin(username string, password string) (ok bool, err error) {
	user, err := u.Repo.QueryUserByName(username)
	if err != nil {
		return false, err
	}
	plainPass := utils.Base64Decode(user.Password)
	if password == plainPass {
		return true, nil
	}
	return false, nil
}

func (u *userCase) UserRegister(user model.User) (ok bool, err error) {
	encodePass := utils.Base64Encode(user.Password)
	user.Password = encodePass
	err = u.Repo.InsertUser(user)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *userCase) GetPatientList(page int, size int) (patient []*model.Patient, totalNum int, totalPage int, err error) {
	totalNum, totalPage = u.Repo.CountPatientList(size)
	patient, err = u.Repo.QueryPatientList(page, size)
	return
}

func (u *userCase) AddPatient(data model.Patient) (err error) {
	return u.Repo.InsertPatient(data)
}

func (u *userCase) DeletePatient(data model.Patient) (err error) {
	return u.Repo.DeletePatientByID(data)
}

func (u *userCase) EditPatient(data model.Patient) (err error) {
	return u.Repo.UpdatePatientByID(data)
}

func (u *userCase) QueryPatient(name string) (patients []*model.Patient, err error) {
	return u.Repo.QueryPatientByName(name)
}

func (u *userCase) GetDoctorList(page int, size int) (doctors []*model.Doctor, totalNum int, totalPage int, err error) {
	totalNum, totalPage = u.Repo.CountDoctorList(size)
	doctors, err = u.Repo.QueryDoctorList(page, size)
	return
}

func (u *userCase) AddDoctor(data model.Doctor) (err error) {
	return u.Repo.InsertDoctor(data)
}

func (u *userCase) DeleteDoctor(data model.Doctor) (err error) {
	return u.Repo.DeleteDoctorByID(data)
}

func (u *userCase) EditDoctor(data model.Doctor) (err error) {
	return u.Repo.UpdateDoctorByID(data)
}

func (u *userCase) QueryDoctor(name string, page, size int) (doctors []*model.Doctor, totalNum int, totalPage int, err error) {
	totalNum, totalPage = u.Repo.CountDoctorByName(name, size)
	doctors, err = u.Repo.QueryDoctorByName(name, page, size)
	return
}

func (u *userCase) GetDrugList(page int, size int) (drugs []*model.Drug, totalNum int, totalPage int, err error) {
	totalNum, totalPage = u.Repo.CountDrugList(size)
	drugs, err = u.Repo.QueryDrugList(page, size)
	return
}

func (u *userCase) AddDrug(data model.Drug) (err error) {
	return u.Repo.InsertDrug(data)
}

func (u *userCase) DeleteDrug(data model.Drug) (err error) {
	return u.Repo.DeleteDrugByID(data)
}

func (u *userCase) EditDrug(data model.Drug) (err error) {
	return u.Repo.UpdateDrugByID(data)
}

func (u *userCase) QueryDrug(name string, page, size int) (drugs []*model.Drug, totalNum, totalPage int, err error) {
	totalNum, totalPage = u.Repo.CountRegisterList(size)
	drugs, err = u.Repo.QueryDrugByName(name, page, size)
	return
}

func (u *userCase) GetRegisterList(page int, size int) (result []*model.RegisterResult, totalNum int, totalPage int, err error) {
	totalNum, totalPage = u.Repo.CountRegisterList(size)
	result, err = u.Repo.QueryRegisterList(page, size)
	return
}

func (u *userCase) AddRegister(data model.Register) (err error) {
	return u.Repo.InsertRegister(data)
}

func (u *userCase) DeleteRegister(data model.Register) (err error) {
	return u.Repo.DeleteRegisterByID(data)
}

func (u *userCase) GetPrescriptionList(page int, size int) (prescriptions []*model.Prescription, totalNum int, totalPage int, err error) {
	totalNum, totalPage = u.Repo.CountPrescriptionList(size)
	prescriptions, err = u.Repo.QueryPrescriptionList(page, size)
	return
}

func (u *userCase) AddPrescription(data model.Prescription) (err error) {
	return u.Repo.InsertPrescription(data)
}

func (u *userCase) DeletePrescription(data model.Prescription) (err error) {
	return u.Repo.DeletePrescriptionByID(data)
}

func (u *userCase) EditPrescription(data model.Prescription) (err error) {
	return u.Repo.UpdatePrescriptionByID(data)
}

func (u *userCase) QueryPrescription(id int) (prescription *model.Prescription, err error) {
	return u.Repo.QueryPrescriptionByID(id)
}
