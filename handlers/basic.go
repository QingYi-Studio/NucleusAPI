package handlers

import (
	"path/filepath"
	"runtime"
)

// 获取项目根目录的绝对路径
var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Join(filepath.Dir(b), "..")
)
