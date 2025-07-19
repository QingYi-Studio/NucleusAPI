package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
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

func GetIntelCPUInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	model := vars["model"]
	filePath := filepath.Join(basePath, "data", "cpu", "intel", model+".json")

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "CPU information not found", http.StatusNotFound)
		return
	}

	// 读取CPU信息文件
	data, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Error reading CPU information", http.StatusInternalServerError)
		return
	}

	// 解析JSON数据
	// 这一步好像会导致数据上下乱了，不过不影响，无所谓了
	var cpuInfo map[string]interface{}
	if err := json.Unmarshal(data, &cpuInfo); err != nil {
		http.Error(w, "Error parsing CPU information", http.StatusInternalServerError)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// 返回CPU信息
	json.NewEncoder(w).Encode(cpuInfo)
}

func GetAMDCPUInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	model := vars["model"]
	filePath := filepath.Join(basePath, "data", "cpu", "amd", model+".json")

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "CPU information not found", http.StatusNotFound)
		return
	}

	// 读取CPU信息文件
	data, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Error reading CPU information", http.StatusInternalServerError)
		return
	}

	// 解析JSON数据
	var cpuInfo map[string]interface{}
	if err := json.Unmarshal(data, &cpuInfo); err != nil {
		http.Error(w, "Error parsing CPU information", http.StatusInternalServerError)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// 返回CPU信息
	json.NewEncoder(w).Encode(cpuInfo)
}

func GetAppleCPUInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	model := vars["model"]
	filePath := filepath.Join(basePath, "data", "cpu", "apple", model+".json")

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "CPU information not found", http.StatusNotFound)
		return
	}

	// 读取CPU信息文件
	data, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Error reading CPU information", http.StatusInternalServerError)
		return
	}

	// 解析JSON数据
	var cpuInfo map[string]interface{}
	if err := json.Unmarshal(data, &cpuInfo); err != nil {
		http.Error(w, "Error parsing CPU information", http.StatusInternalServerError)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// 返回CPU信息
	json.NewEncoder(w).Encode(cpuInfo)
}
