package pi

import "os"

func prepare() (string, string, string, string) {
	beforeAPIBaseEnv := os.Getenv("PIXELA_API_BASE")
	afterAPIBaseEnv := "pixela.example.com"
	os.Setenv("PIXELA_API_BASE", afterAPIBaseEnv)
	beforeTokenEnv := os.Getenv("PIXELA_USER_TOKEN")
	afterTokenEnv := "thisissecret"
	os.Setenv("PIXELA_USER_TOKEN", afterTokenEnv)
	return beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, afterTokenEnv
}

func cleanup(beforeAPIBaseEnv string, beforeTokenEnv string) {
	os.Setenv("PIXELA_API_BASE", beforeAPIBaseEnv)
	os.Setenv("PIXELA_USER_TOKEN", beforeTokenEnv)
}
