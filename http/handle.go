package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hospital-admin/model"
	"hospital-admin/usecase"
	"net/http"
	"strconv"
)

func NewHandle(u usecase.UseCase) *Handle {
	return &Handle{UseCase: u}
}

type Handle struct {
	UseCase usecase.UseCase
}

// 用户登录
func (h *Handle) UserLogin(ctx *gin.Context) {
	userLoginReq := &UserLoginReq{}
	userLoginResp := &UserLoginResp{}
	err := ctx.Bind(userLoginReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, getBasicResp(userLoginResp))
	}
	ok, err := h.UseCase.UserLogin(userLoginReq.Username, userLoginReq.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, getBasicResp(userLoginResp))
	}
	if ok {
		userLoginResp.IsAllow = true
		ctx.JSON(http.StatusOK, getBasicResp(userLoginResp))
	}
}

// 用户注册
func (h *Handle) UserRegister(ctx *gin.Context) {
	userRegisterReq := &UserRegisterReq{}
	userRegisterResp := &UserRegisterResp{}
	err := ctx.Bind(userRegisterReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, getBasicResp(userRegisterResp))
	}
	user := model.User{
		UserName:  userRegisterReq.Username,
		Password:  userRegisterReq.Password,
		UserType:  userRegisterReq.UserType,
		PatientID: 0,
		DoctorID:  0,
	}
	ok, err := h.UseCase.UserRegister(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, getBasicResp(userRegisterResp))
	}
	if ok {
		ctx.JSON(http.StatusOK, getBasicResp(userRegisterResp))
	}
}

func (h *Handle) GetPatientList(ctx *gin.Context) {
	patientList, totalNum, totalPage, err := h.UseCase.GetPatientList(1, 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	data := make(map[string]interface{}, 0)
	data["totalNum"] = totalNum
	data["totalPage"] = totalPage
	data["patientList"] = patientList
	ctx.JSON(http.StatusOK, getBasicResp(data))
}

func (h *Handle) AddPatient(ctx *gin.Context) {
	patient := model.Patient{}
	err := ctx.Bind(&patient)
	if err != nil {
		fmt.Println(err)
	}
	err = h.UseCase.AddPatient(patient)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (h *Handle) DeletePatient(ctx *gin.Context) {
	patient := model.Patient{}
	err := ctx.Bind(&patient)
	if err != nil {
		fmt.Println(err)
	}
	err = h.UseCase.DeletePatient(patient)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	ctx.JSON(http.StatusOK, nil)
}

func (h *Handle) EditPatient(ctx *gin.Context) {
	patient := model.Patient{}
	err := ctx.Bind(&patient)
	if err != nil {
		fmt.Println(err)
	}
	err = h.UseCase.EditPatient(patient)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	ctx.JSON(http.StatusOK, nil)
}

func (h *Handle) QueryPatient(ctx *gin.Context) {
	name := ctx.Query("name")
	res, err := h.UseCase.QueryPatient(name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	ctx.JSON(http.StatusOK, getBasicResp(res))
}

func (h *Handle) GetDoctorList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))

	patientList, totalNum, totalPage, err := h.UseCase.GetDoctorList(page, size)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	data := make(map[string]interface{}, 0)
	data["totalNum"] = totalNum
	data["totalPage"] = totalPage
	data["patientList"] = patientList
	ctx.JSON(http.StatusOK, getBasicResp(data))
}

func (h *Handle) AddDoctor(ctx *gin.Context) {
	doctor := model.Doctor{}
	err := ctx.Bind(&doctor)
	if err != nil {
		fmt.Println(err)
	}
	err = h.UseCase.AddDoctor(doctor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (h *Handle) DeleteDoctor(ctx *gin.Context) {
	doctor := model.Doctor{}
	err := ctx.Bind(&doctor)
	if err != nil {
		fmt.Println(err)
	}
	err = h.UseCase.DeleteDoctor(doctor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	ctx.JSON(http.StatusOK, nil)
}

func (h *Handle) EditDoctor(ctx *gin.Context) {
	doctor := model.Doctor{}
	err := ctx.Bind(&doctor)
	if err != nil {
		fmt.Println(err)
	}
	err = h.UseCase.EditDoctor(doctor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	ctx.JSON(http.StatusOK, nil)
}

func (h *Handle) QueryDoctor(ctx *gin.Context) {
	name := ctx.Query("name")
	res, err := h.UseCase.QueryDoctor(name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	ctx.JSON(http.StatusOK, getBasicResp(res))
}

func (h *Handle) GetDrugList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))

	patientList, totalNum, totalPage, err := h.UseCase.GetDrugList(page, size)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	data := make(map[string]interface{}, 0)
	data["totalNum"] = totalNum
	data["totalPage"] = totalPage
	data["patientList"] = patientList
	ctx.JSON(http.StatusOK, getBasicResp(data))
}

func (h *Handle) AddDrug(ctx *gin.Context) {
	doctor := model.Drug{}
	err := ctx.Bind(&doctor)
	if err != nil {
		fmt.Println(err)
	}
	err = h.UseCase.AddDrug(doctor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (h *Handle) DeleteDrug(ctx *gin.Context) {
	doctor := model.Drug{}
	err := ctx.Bind(&doctor)
	if err != nil {
		fmt.Println(err)
	}
	err = h.UseCase.DeleteDrug(doctor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	ctx.JSON(http.StatusOK, nil)
}

func (h *Handle) EditDrug(ctx *gin.Context) {
	doctor := model.Drug{}
	err := ctx.Bind(&doctor)
	if err != nil {
		fmt.Println(err)
	}
	err = h.UseCase.EditDrug(doctor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	ctx.JSON(http.StatusOK, nil)
}

func (h *Handle) QueryDrug(ctx *gin.Context) {
	name := ctx.Query("name")
	res, err := h.UseCase.QueryDrug(name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	ctx.JSON(http.StatusOK, getBasicResp(res))
}

func (h *Handle) GetRegisterList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))

	patientList, totalNum, totalPage, err := h.UseCase.GetRegisterList(page, size)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	data := make(map[string]interface{}, 0)
	data["totalNum"] = totalNum
	data["totalPage"] = totalPage
	data["patientList"] = patientList
	ctx.JSON(http.StatusOK, getBasicResp(data))
}

func (h *Handle) AddRegister(ctx *gin.Context) {
	doctor := model.Register{}
	err := ctx.Bind(&doctor)
	if err != nil {
		fmt.Println(err)
	}
	err = h.UseCase.AddRegister(doctor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (h *Handle) DeleteRegister(ctx *gin.Context) {
	doctor := model.Register{}
	err := ctx.Bind(&doctor)
	if err != nil {
		fmt.Println(err)
	}
	err = h.UseCase.DeleteRegister(doctor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	ctx.JSON(http.StatusOK, nil)
}

func (h *Handle) GetPrescriptionList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))

	patientList, totalNum, totalPage, err := h.UseCase.GetPrescriptionList(page, size)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	data := make(map[string]interface{}, 0)
	data["totalNum"] = totalNum
	data["totalPage"] = totalPage
	data["patientList"] = patientList
	ctx.JSON(http.StatusOK, getBasicResp(data))
}

func (h *Handle) AddPrescription(ctx *gin.Context) {
	doctor := model.Prescription{}
	err := ctx.Bind(&doctor)
	if err != nil {
		fmt.Println(err)
	}
	err = h.UseCase.AddPrescription(doctor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (h *Handle) DeletePrescription(ctx *gin.Context) {
	doctor := model.Prescription{}
	err := ctx.Bind(&doctor)
	if err != nil {
		fmt.Println(err)
	}
	err = h.UseCase.DeletePrescription(doctor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	ctx.JSON(http.StatusOK, nil)
}

func (h *Handle) EditPrescription(ctx *gin.Context) {
	doctor := model.Prescription{}
	err := ctx.Bind(&doctor)
	if err != nil {
		fmt.Println(err)
	}
	err = h.UseCase.EditPrescription(doctor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	ctx.JSON(http.StatusOK, nil)
}

func (h *Handle) QueryPrescription(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	res, err := h.UseCase.QueryPrescription(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	ctx.JSON(http.StatusOK, getBasicResp(res))
}
