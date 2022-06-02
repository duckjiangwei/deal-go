package config

import (
	"hk591_go/pkg/config"
)

func init() {
	config.Add("domain", func() map[string]interface{} {

		httpPrx := "https://"
		sign := ""
		appEnv := config.Env("APP_ENV", "dev")

		if appEnv == "dev" {
			sign = ".dev"
		} else if appEnv == "debug" {
			sign = ".debug"
		}
		return map[string]interface{}{
			"base_url": httpPrx + "www" + sign + ".591.com.hk",
			"p1_url":   httpPrx + "p1" + sign + ".591.com.hk",
		}
	})
}
