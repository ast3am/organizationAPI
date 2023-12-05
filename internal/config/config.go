package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"gitlab.com/ast3am77/test-go/internal/models"
	"gitlab.com/ast3am77/test-go/pkg/logging"
)

var cfg *models.Config

func GetConfig(path string) *models.Config {
	logger := logging.GetLogger("")
	logger.DebugMsg("read config")
	cfg = &models.Config{}
	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		logger.FatalMsg("", err)
	}
	return cfg
}
