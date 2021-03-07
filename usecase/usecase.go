package usecase

import (
	"github.com/jmoiron/sqlx"
	"hospital-admin/model"
	"hospital-admin/utils"
)

func NewUseCase(db *sqlx.DB) UseCase {
	return &userCase{
		Repo: model.NewRepository(db),
	}
}

type userCase struct {
	Repo model.Repository
}

type UseCase interface {
	UserLogin(username string, password string) (ok bool, err error)
	UserRegister(username string, password string, userType int64, patientID int64, doctorID int64) (ok bool, err error)
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

func (u *userCase) UserRegister(username string, password string, userType int64, patientID int64, doctorID int64) (ok bool, err error) {
	encodePass := utils.Base64Encode(password)
	err = u.Repo.InsertUser(username, encodePass, userType, patientID, doctorID)
	if err != nil {
		return false, err
	}
	return true, nil
}
