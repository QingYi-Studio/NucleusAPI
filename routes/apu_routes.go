package routes

import (
	"NucleusAPI/handlers"
	"fmt"

	"github.com/gorilla/mux"
)

// RegisterAPURoutes 注册APU相关路由（带限流）
func RegisterAPURoutes(router *mux.Router) {
	fmt.Printf("RegisterAPURoutes start\n")
	apuRouter := router.PathPrefix("/apu").Subrouter()

	apuRouter.HandleFunc("/amd/{model}", handlers.GetAMDAPUInfo).Methods("GET")
	apuRouter.HandleFunc("/{model}", handlers.GetAMDAPUInfo).Methods("GET")
	fmt.Printf("RegisterAPURoutes end\n")
}
