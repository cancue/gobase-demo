package config

import (
	_ "embed" // embed.
	"github.com/cancue/gobase/config"
)

//go:embed local.yaml
var local []byte

//go:embed prd.yaml
var prd []byte

func init() {
	yamlFiles := map[string][]byte{
		"local": local,
		"prd":   prd,
	}
	config.SetYAMLs(yamlFiles)
}
