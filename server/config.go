package main

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	SERVER_HOST    string
	SERVER_PORT    string
	DB_DSN         string
	LNBITS_URL     string
	LNBITS_API_KEY string
	LNBITS_USER_ID string
}

func GetConfig(params ...string) Configuration {
	configuration := Configuration{}
	env := "dev"
	if len(params) > 0 {
		env = params[0]
	}
	fileName := fmt.Sprintf("./%s_config.json", env)
	gonfig.GetConf(fileName, &configuration)
	return configuration
}
