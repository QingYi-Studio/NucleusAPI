package routes

import "github.com/gorilla/mux"

// NewRouter 创建并配置主路由器
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// 注册各类路由
	RegisterBaseRoutes(router)
	RegisterCPURoutes(router)
	RegisterAPURoutes(router)

	return router
}
