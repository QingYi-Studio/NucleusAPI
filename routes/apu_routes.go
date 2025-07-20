package routes

import (
	"NucleusAPI/handlers"

	"github.com/gorilla/mux"
)

// RegisterAPURoutes 注册APU相关路由（带限流）
func RegisterAPURoutes(router *mux.Router) {
	apuRouter := router.PathPrefix("/apu").Subrouter()

	apuRouter.HandleFunc("/amd/{model}", handlers.GetAMDAPUInfo).Methods("GET")
	apuRouter.HandleFunc("/{model}", handlers.GetAMDAPUInfo).Methods("GET")
}
