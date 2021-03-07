package http

func getBasicResp(data interface{}) BasicResp {
	return BasicResp{
		ErrCode: 1000,
		Msg:     "",
		Data:    data,
	}
}

type BasicResp struct {
	ErrCode int64       `json:"err_code"`
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
	UserType  int64  `json:"user_type"`
	PatientID int64  `json:"patient_id"`
	DoctorID  int64  `json:"doctor_id"`
}

type UserRegisterResp struct {
}

type GetPatientListReq struct {
	PageNum  int64 `json:"page_num"`
	PageSize int64 `json:"page_size"`
}

type GetPatientListResp struct {
}
