package pi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type channelsCommand struct {
	Create createChannelCommand `description:"create Channel" command:"create" subcommands-optional:"true"`
	Get    getChannelsCommand   `description:"get Channel Definitions" command:"get" subcommands-optional:"true"`
}

type createChannelCommand struct {
	Username string `short:"u" long:"username" description:"User name of channel owner."`
	ID       string `short:"i" long:"channel-id" description:"ID for identifying the channel." required:"true"`
	Name     string `short:"n" long:"name" description:"The name of the channel." required:"true"`
	Type     string `short:"t" long:"type" description:"The type for notification." required:"true"`
	Detail   string `short:"d" long:"detail" description:"Object that specifies the details of the type. It is specified as JSON string." required:"true"`
}

type createChannelParam struct {
	ID     string          `json:"id"`
	Name   string          `json:"name"`
	Type   string          `json:"type"`
	Detail json.RawMessage `json:"detail"`
}

type getChannelsCommand struct {
	Username string `short:"u" long:"username" description:"User name of channel owner."`
}

func (cC *createChannelCommand) Execute(args []string) error {
	req, err := generateCreateChannelRequest(cC)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateCreateChannelRequest(cC *createChannelCommand) (*http.Request, error) {
	username, err := getUsername(cC.Username)
	if err != nil {
		return nil, err
	}

	paramStruct := &createChannelParam{
		ID:     cC.ID,
		Name:   cC.Name,
		Type:   cC.Type,
		Detail: json.RawMessage(cC.Detail),
	}

	req, err := generateRequestWithToken(
		"POST",
		fmt.Sprintf("v1/users/%s/channels", username),
		paramStruct,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate create api request : %s", err)
	}

	return req, nil
}

func (gC *getChannelsCommand) Execute(args []string) error {
	req, err := generateGetChannelsRequest(gC)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateGetChannelsRequest(gC *getChannelsCommand) (*http.Request, error) {
	username, err := getUsername(gC.Username)
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
