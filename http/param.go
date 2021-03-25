package http

import "hospital-admin/model"

func getBasicResp(data interface{}) BasicResp {
	return BasicResp{
		ErrCode: 1000,
		Msg:     "",
		Data:    data,
	}
}

type BasicResp struct {
	ErrCode uint        `json:"err_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type UserLoginReq struct {
	Username string `json:"username""`
	Password string `json:"password"`
}

type UserLoginResp struct {
	IsAllow bool `json:"is_allow"`
}

type UserRegisterReq struct {
	Username  string `json:"username""`
	Password  string `json:"password"`
	UserType  uint8  `json:"user_type"`
	PatientID uint   `json:"patient_id"`
	DoctorID  uint   `json:"doctor_id"`
}

type UserRegisterResp struct {
}

type GetUserListReq struct {
	PageNum  int `json:"page_num"`
	PageSize int `json:"page_size"`
}

type GetUserListResp struct {
	PageNum   int           `json:"page_num"`
	PageSize  int           `json:"page_size"`
	TotalNum  int           `json:"total_num"`
	TotalSize int           `json:"total_size"`
	UserList  []*model.User `json:"user_list"`
}

type GetPatientListReq struct {
	PageNum  int `json:"page_num"`
	PageSize int `json:"page_size"`
}

type GetPatientListResp struct {
	PageNum     int              `json:"page_num"`
	PageSize    int              `json:"page_size"`
	TotalNum    int              `json:"total_num"`
	TotalSize   int              `json:"total_size"`
	PatientList []*model.Patient `json:"patient_list"`
}
