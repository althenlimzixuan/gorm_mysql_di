package config

var ConfigHeader = "GORM_MYSQL"

type GormMySqlConfig struct {
	Url     string `envconfig:"URL"`
	LogMode bool   `envconfig:"LOGMODE"`
}
