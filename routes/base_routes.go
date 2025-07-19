package routes

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gorilla/mux"
)

// 获取项目根目录的绝对路径
var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Join(filepath.Dir(b), "..")
)

// RegisterBaseRoutes 注册基础路由
func RegisterBaseRoutes(router *mux.Router) {
	// 根路径：返回 HTML 页面
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveHTML(w, filepath.Join(basePath, "pages", "index.html"))
	})

	// 文档路径：返回 HTML 页面
	router.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		serveHTML(w, filepath.Join(basePath, "pages", "docs.html"))
	})
}

// serveHTML 辅助函数：提供HTML文件
func serveHTML(w http.ResponseWriter, filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "file not found: "+err.Error(), http.StatusNotFound)
		return
	}
	defer f.Close()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.Copy(w, f)
}
