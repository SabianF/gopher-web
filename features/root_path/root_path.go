package root_path

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)

	root = filepath.Join(filepath.Dir(b), "../..")
)

func GetRootPath() string {
	return root
}
