package config

import (
	"gvb/config/custom"
	"gvb/config/system"
)

type Config struct {
	System system.System `json:"-"`
	Custom custom.Custom `toml:"custom"`
}
