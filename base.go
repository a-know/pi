package pi

import (
	"fmt"
	"os"
)

func getUsername(cmdUsername string) (string, error) {
	username := cmdUsername
	if username == "" {
		username = os.Getenv("PIXELA_USER_NAME")
	}
	if username == "" {
		return username, fmt.Errorf("`username` not specified. Please specify username by command line option or `PIXELA_USER_NAME` environment variables")
	}
	return username, nil
}
