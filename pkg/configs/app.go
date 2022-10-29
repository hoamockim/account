package configs

import "fmt"

type appConfig struct {
	RunMode string `json:"run_mode" env:"run_mode"`
	Host    string `json:"host" env:"HOST"`
	Port    int    `json:"port" env:"PORT"`
	MySql   MySQL  `json:"db"`
	Mongo   Mongo  `json:"mongo"`
	Redis   Redis  `json:"redis"`
	Jwt     JwtKey `json:"jwt"`
	Cert    Cert   `json:"cert"`
}

type configValidate interface {
	isValid() bool
}

func InitConfig() {
	//TODO: init config from env
}

func (cfg *appConfig) isValid() bool {
	return cfg.MySql.isValid()
}

func AppURL() string {
	return fmt.Sprintf("%s:%d", app.Host, app.Port)
}
