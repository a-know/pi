package pi

import (
	"fmt"
	"net/http"
)

type channelsCommand struct {
	Get getChannelsCommand `description:"get Channel Definitions" command:"get" subcommands-optional:"true"`
}

type getChannelsCommand struct {
	Username string `short:"u" long:"username" description:"User name of channel owner."`
}

func (gG *getChannelsCommand) Execute(args []string) error {
	req, err := generateGetChannelsRequest(gG)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateGetChannelsRequest(gG *getChannelsCommand) (*http.Request, error) {
	username, err := getUsername(gG.Username)
	if err != nil {
		return nil, err
	}

	req, err := generateRequestWithToken(
		"GET",
		fmt.Sprintf("v1/users/%s/channels", username),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate get api request : %s", err)
	}

	return req, nil
}
