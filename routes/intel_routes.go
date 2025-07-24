package routes

import (
	"NucleusAPI/handlers"
	"fmt"
	"github.com/gorilla/mux"
)

// RegisterCPURoutes 注册CPU相关路由
func RegisterIntelRoutes(router *mux.Router) {
	fmt.Printf("RegisterIntelRoutes start\n")
	intelRouter := router.PathPrefix("/intel").Subrouter()

	intelRouter.HandleFunc("/cpu/{model}", handlers.GetIntelCPUInfo).Methods("GET")
	fmt.Printf("RegisterIntelRoutes end\n")
}
