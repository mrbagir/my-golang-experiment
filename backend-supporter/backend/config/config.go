package config

import "backend-supporter/backend/api/app/setting/model"

type AppConfig struct {
	Apps []model.App `yaml:"apps"`
}
