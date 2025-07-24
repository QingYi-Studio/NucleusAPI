package routes

import (
	"NucleusAPI/handlers"
	"fmt"

	"github.com/gorilla/mux"
)

// RegisterCPURoutes 注册CPU相关路由（带限流）
func RegisterCPURoutes(router *mux.Router) {
	fmt.Printf("RegisterCPURoutes start\n")
	cpuRouter := router.PathPrefix("/cpu").Subrouter()

	cpuRouter.HandleFunc("/intel/{model}", handlers.GetIntelCPUInfo).Methods("GET")
	cpuRouter.HandleFunc("/amd/{model}", handlers.GetAMDCPUInfo).Methods("GET")
	cpuRouter.HandleFunc("/apple/{model}", handlers.GetAppleCPUInfo).Methods("GET")
	fmt.Printf("RegisterCPURoutes end\n")
}
