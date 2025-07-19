package routes

import (
	"ComputerInfoAPI/handlers"

	"github.com/gorilla/mux"
)

// RegisterAPURoutes 注册APU相关路由
func RegisterAPURoutes(router *mux.Router) {
	apuRouter := router.PathPrefix("/apu").Subrouter()

	apuRouter.HandleFunc("/amd/{model}", handlers.GetAMDCPUInfo).Methods("GET")
	apuRouter.HandleFunc("/{model}", handlers.GetAMDCPUInfo).Methods("GET")
}
