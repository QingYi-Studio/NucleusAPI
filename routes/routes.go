package routes

import (
	"ComputerInfoAPI/handlers"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

// 获取项目根目录的绝对路径
var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Join(filepath.Dir(b), "..")
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// 根路径：返回 HTML 页面
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		filePath := filepath.Join(basePath, "pages", "index.html")

		f, err := os.Open(filePath)
		if err != nil {
			http.Error(w, "index.html not found: "+err.Error(), http.StatusNotFound)
			return
		}
		defer f.Close()

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.Copy(w, f)
	})

	// 文档路径：返回 HTML 页面
	router.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		filePath := filepath.Join(basePath, "pages", "docs.html")

		f, err := os.Open(filePath)
		if err != nil {
			http.Error(w, "docs.html not found: "+err.Error(), http.StatusNotFound)
			return
		}
		defer f.Close()

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.Copy(w, f)
	})

	cpuRouter := router.PathPrefix("/cpu").Subrouter()

	// Intel CPU 信息路由
	cpuRouter.HandleFunc("/intel/{model}", handlers.GetIntelCPUInfo).Methods("GET")

	// AMD CPU 信息路由
	cpuRouter.HandleFunc("/amd/{model}", handlers.GetAMDCPUInfo).Methods("GET")

	// Apple CPU 信息路由
	cpuRouter.HandleFunc("/apple/{model}", handlers.GetAppleCPUInfo).Methods("GET")

	return router
}
