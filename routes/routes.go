package routes

import (
	"ComputerInfoAPI/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// 添加根路径处理
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Computer Info API!"))
		w.Write([]byte("This API is provided free of charge by QingYi Studio."))
	})

	// Intel CPU信息路由
	intelCpuRouter := router.PathPrefix("/cpu").Subrouter()
	intelCpuRouter.HandleFunc("/intel/{model}", handlers.GetIntelCPUInfo).Methods("GET")

	// Intel CPU信息路由
	amdCpuRouter := router.PathPrefix("/cpu").Subrouter()
	amdCpuRouter.HandleFunc("/amd/{model}", handlers.GetAMDCPUInfo).Methods("GET")

	// NVIDIA GPU信息路由

	return router
}
