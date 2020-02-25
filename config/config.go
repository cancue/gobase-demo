package config

import (
	"path"
	"path/filepath"
	"runtime"

	"github.com/cancue/gobase/util"
)

// GetStageYAML reads yaml file and returns the contents.
//
// The name of the yaml file is determined by an environment variable whose key is 'STAGE'.
// If there is no environment variable, 'local' is selected by default. (e.g. 'local.yaml')
func GetStageYAML() (stage string, yaml map[string]interface{}) {
	stage = util.GetEnv("STAGE", "local")
	_, b, _, _ := runtime.Caller(0)
	path := filepath.Join(path.Dir(b), stage+".yaml")
	yaml = util.ReadYAMLFile(path)

	return
}
