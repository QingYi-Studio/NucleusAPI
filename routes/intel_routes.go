package routes

import (
	"NucleusAPI/handlers"
	"github.com/gorilla/mux"
)

// RegisterCPURoutes 注册CPU相关路由（带限流）
func RegisterIntelRoutes(router *mux.Router) {
	cpuRouter := router.PathPrefix("/intel").Subrouter()

	cpuRouter.HandleFunc("/cpu/{model}", handlers.GetIntelCPUInfo).Methods("GET")
}
