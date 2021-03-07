package http

import (
	"github.com/gin-gonic/gin"
	"hospital-admin/usecase"
	"net/http"
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
	ok, err := h.UseCase.UserRegister(userRegisterReq.Username, userRegisterReq.Password, userRegisterReq.UserType, userRegisterReq.PatientID, userRegisterReq.DoctorID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, getBasicResp(userRegisterResp))
	}
	if ok {
		ctx.JSON(http.StatusOK, getBasicResp(userRegisterResp))
	}
}

func (h *Handle) GetPatientList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}
