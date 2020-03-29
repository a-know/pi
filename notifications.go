package pi

import (
	"fmt"
	"net/http"
)

type notificationsCommand struct {
	Get getNotificationsCommand `description:"get Notifications" command:"get" subcommands-optional:"true"`
}

type getNotificationsCommand struct {
	Username string `short:"u" long:"username" description:"User name of owner."`
	GraphID  string `short:"g" long:"graph-id" description:"ID for identifying the graph." required:"true"`
}

func (gN *getNotificationsCommand) Execute(args []string) error {
	req, err := generateGetNotificationsRequest(gN)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateGetNotificationsRequest(gN *getNotificationsCommand) (*http.Request, error) {
	username, err := getUsername(gN.Username)
	if err != nil {
		return nil, err
	}

	req, err := generateRequestWithToken(
		"GET",
		fmt.Sprintf("v1/users/%s/graphs/%s/notifications", username, gN.GraphID),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate get api request : %s", err)
	}

	return req, nil
}
