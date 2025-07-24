package routes

import (
	"NucleusAPI/handlers"
	"fmt"
	"github.com/gorilla/mux"
)

// RegisterCPURoutes 注册CPU相关路由
func RegisterAMDRoutes(router *mux.Router) {
	fmt.Printf("RegisterAMDRoutes start\n")
	amdRouter := router.PathPrefix("/amd").Subrouter()

	amdRouter.HandleFunc("/cpu/{model}", handlers.GetAMDCPUInfo).Methods("GET")
	amdRouter.HandleFunc("/apu/{model}", handlers.GetAMDAPUInfo).Methods("GET")
	fmt.Printf("RegisterAMDRoutes end\n")
}
