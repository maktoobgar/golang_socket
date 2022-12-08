package configs

import (
	_ "embed"

	"github.com/maktoobgar/golang_socket/pkg/config"
)

var (
	//go:embed config.yml
	embededConfig []byte

	CFG = &Config{}
)

type Config struct {
	Host         string   `yaml:"host"`
	Port         int      `yaml:"port"`
	AllowOrigins []string `yaml:"allow_origins"`
	Debug        bool     `yaml:"debug"`
}

func init() {
	config.ParseBytes(embededConfig, CFG)
	config.ReadLocalConfigs(CFG)
}
