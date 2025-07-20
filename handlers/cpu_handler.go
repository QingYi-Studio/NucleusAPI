package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path/filepath"
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

	// 定义需要重定向的APU型号集合
	apuModels := map[string]bool{
		"A4-3300":       true,
		"A4-3400":       true,
		"A4-3420":       true,
		"A4-5300":       true,
		"A4-6300":       true,
		"A6-3600":       true,
		"A6-3650":       true,
		"A6-3670K":      true,
		"A6-5400K":      true,
		"A6-6400K":      true,
		"A6-7400K":      true,
		"A6-7470K":      true,
		"A6-7480":       true,
		"A6-9500":       true,
		"A6-9500E":      true,
		"A8-3800":       true,
		"A8-3850":       true,
		"A8-3870K":      true,
		"A8-5500":       true,
		"A8-5600K":      true,
		"A8-6600K":      true,
		"A8-7600":       true,
		"A8-7650K":      true,
		"A8-7670K":      true,
		"A8-7680":       true,
		"A8-9600":       true,
		"A10-5700":      true,
		"A10-5800K":     true,
		"A10-6800K":     true,
		"A10-7700K":     true,
		"A10-7800":      true,
		"A10-7850K":     true,
		"A10-7860K":     true,
		"A10-7870K":     true,
		"A10-7890K":     true,
		"A10-9700":      true,
		"A10-9700E":     true,
		"A12-9800":      true,
		"A12-9800E":     true,
		"AthlonX4-631":  true,
		"AthlonX4-638":  true,
		"AthlonX4-641":  true,
		"AthlonX4-651":  true,
		"AthlonX4-740":  true,
		"AthlonX4-750":  true,
		"AthlonX4-750K": true,
		"AthlonX4-760K": true,
		"AthlonX4-840":  true,
		"AthlonX4-845":  true,
		"AthlonX4-860K": true,
		"AthlonX4-870K": true,
		"AthlonX4-880K": true,
		"AthlonX4-950":  true,
	}

	// 检查当前型号是否需要重定向
	if apuModels[model] {
		// 构建新的URL路径并重定向
		newPath := "/apu/amd/" + model
		http.Redirect(w, r, newPath, http.StatusMovedPermanently)
		return
	}

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
