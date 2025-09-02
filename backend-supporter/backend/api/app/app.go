package app

import "backend-supporter/backend/api/app/setting/model"

func init() {
	Apps = []model.App{
		{
			Name: "MyApp",
			Settings: []model.Setting{
				{Name: "Setting1", Active: true, Value: "Value1"},
				{Name: "Setting2", Active: false, Value: "Value2"},
			},
			Terminal: &model.Terminal{
				Log:        []string{"Log1", "Log2"},
				MaxLines:   100,
				AutoScroll: true,
			},
			Hide: false,
		},
		{
			Name: "MySecondApp",
			Settings: []model.Setting{
				{Name: "SettingA", Active: true, Value: "ValueA"},
				{Name: "SettingB", Active: false, Value: "ValueB"},
			},
			Terminal: &model.Terminal{
				Log:        []string{"LogA", "LogB"},
				MaxLines:   50,
				AutoScroll: false,
			},
			Hide: false,
		},
	}
}

var Apps []model.App
