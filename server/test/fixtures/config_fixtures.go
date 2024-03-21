package fixtures

import "os"

func CreateEnvs() {
	os.Setenv("ZINC_USER", "admin")
	os.Setenv("ZINC_PASSWORD", "password")
}

func RemoveEnvs() {
	os.Unsetenv("ZINC_USER")
	os.Unsetenv("ZINC_PASSWORD")
}
