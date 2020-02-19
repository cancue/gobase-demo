package config

import (
	"path"
	"path/filepath"
	"runtime"

	"github.com/cancue/gobase/util"
)

func GetStageYAML() (stage string, yaml map[string]interface{}) {
	stage = util.GetEnv("STAGE", "local")
	_, b, _, _ := runtime.Caller(0)
	path := filepath.Join(path.Dir(b), stage+".yaml")
	yaml = util.ReadYAMLFile(path)

	return
}
