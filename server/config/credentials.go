package config

import "log"

type ZincsearchCredentials struct {
	User     string
	Password string
}

var zincsearchCredentials *ZincsearchCredentials

func LoadZincsearchCredentials() ZincsearchCredentials {
	if zincsearchCredentials != nil {
		log.Println("Credentials already loaded")
		return *zincsearchCredentials
	}

	log.Println("Loading zincsearch credentials...")
	credentials := &ZincsearchCredentials{}
	// Emulation of config loading since the doc requested to no use any other external library
	// so I can't use something like `godotenv` to get the user credentials for a .env file
	credentials.User = "admin"
	credentials.Password = "password"

	zincsearchCredentials = credentials

	return *credentials
}

func GetZincsearchCredentials() *ZincsearchCredentials {
	return zincsearchCredentials
}
