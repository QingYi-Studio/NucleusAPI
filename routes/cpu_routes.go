package routes

import (
	"ComputerInfoAPI/handlers"

	"github.com/gorilla/mux"
)

// RegisterCPURoutes 注册CPU相关路由
func RegisterCPURoutes(router *mux.Router) {
	cpuRouter := router.PathPrefix("/cpu").Subrouter()

	cpuRouter.HandleFunc("/intel/{model}", handlers.GetIntelCPUInfo).Methods("GET")
	cpuRouter.HandleFunc("/amd/{model}", handlers.GetAMDCPUInfo).Methods("GET")
	cpuRouter.HandleFunc("/apple/{model}", handlers.GetAppleCPUInfo).Methods("GET")
}
