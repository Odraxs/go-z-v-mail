package app

type ZincsearchCredentials struct {
	User     string
	Password string
}

func LoadZincsearchCredentials() ZincsearchCredentials {
	cfg := ZincsearchCredentials{}

	// Emulation of config loading since the doc requested to no use any other external library
	// so I can't use something like `godotenv` to get the user credentials for a .env file
	cfg.User = "admin"
	cfg.Password = "password"

	return cfg
}
