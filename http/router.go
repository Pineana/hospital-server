package http

import "github.com/gin-gonic/gin"

func BuildRouter(engine *gin.Engine, handle *Handle) {
	SetBaseRouter(engine, handle)
}

func SetBaseRouter(engine *gin.Engine, handle *Handle) {
	r := engine.Group("/base")
	r.POST("/login", handle.UserLogin)
	r.POST("/register", handle.UserRegister)

	r.GET("/query-patient", handle.QueryPatient)
	r.POST("/add-patient", handle.AddPatient)
	r.PUT("/edit-patient", handle.EditPatient)
	r.DELETE("/delete-patient", handle.DeletePatient)
	r.GET("/get-patient-list", handle.GetPatientList)

	r.GET("/query-doctor", handle.QueryDoctor)
	r.POST("/add-doctor", handle.AddDoctor)
	r.PUT("/edit-doctor", handle.EditDoctor)
	r.DELETE("/delete-doctor", handle.DeleteDoctor)
	r.GET("/get-doctor-list", handle.GetDoctorList)

	r.GET("/get-register-list", handle.GetRegisterList)
	r.GET("/query-register", handle.AddRegister)
	r.POST("/add-register", handle.AddRegister)
	r.DELETE("/delete-register", handle.DeleteRegister)

	r.GET("/query-drug", handle.QueryDrug)
	r.POST("/add-drug", handle.AddDrug)
	r.PUT("/edit-drug", handle.EditDrug)
	r.DELETE("/delete-drug", handle.DeleteDrug)
	r.GET("/get-drug-list", handle.GetDrugList)

	r.GET("/query-prescription", handle.QueryPrescription)
	r.POST("/add-prescription", handle.AddPrescription)
	r.PUT("/edit-prescription", handle.EditPrescription)
	r.DELETE("/delete-prescription", handle.DeletePrescription)
	r.GET("/get-prescription-list", handle.GetPrescriptionList)

}
