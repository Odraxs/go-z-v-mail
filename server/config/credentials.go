package config

import (
	"log"
	"os"
)

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

	// Configure the env vars in the shell, since the doc requested to not use any other external library
	// so I can't use something like `godotenv` to get the user credentials for a .env file
	credentials.User = os.Getenv("ZINC_USER")
	credentials.Password = os.Getenv("ZINC_PASSWORD")

	zincsearchCredentials = credentials

	return *credentials
}

func GetZincsearchCredentials() *ZincsearchCredentials {
	return zincsearchCredentials
}
