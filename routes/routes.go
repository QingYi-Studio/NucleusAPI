package routes

import (
	"NucleusAPI/middleware" // 添加这个导入
	"github.com/gorilla/mux"
)

// NewRouter 创建并配置主路由器
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// 应用 CORS 中间件
	router.Use(middleware.EnableCORS)

	// 注册各类路由
	RegisterBaseRoutes(router)  // 基础路由
	RegisterCPURoutes(router)   // CPU路由
	RegisterAPURoutes(router)   // APU路由
	RegisterIntelRoutes(router) // Intel路由
	RegisterAMDRoutes(router)   // AMD路由

	return router
}
