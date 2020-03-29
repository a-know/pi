package pi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type channelsCommand struct {
	Create createChannelCommand `description:"create Channel" command:"create" subcommands-optional:"true"`
	Update updateChannelCommand `description:"update Channel Definition" command:"update" subcommands-optional:"true"`
	Get    getChannelsCommand   `description:"get Channel Definitions" command:"get" subcommands-optional:"true"`
	Delete deleteChannelCommand `description:"delete Channel" command:"delete" subcommands-optional:"true"`
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

type updateChannelCommand struct {
	Username string `short:"u" long:"username" description:"User name of channel owner."`
	ID       string `short:"i" long:"channel-id" description:"ID for identifying the channel." required:"true"`
	Name     string `short:"n" long:"name" description:"The name of the channel."`
	Type     string `short:"t" long:"type" description:"The type for notification."`
	Detail   string `short:"d" long:"detail" description:"Object that specifies the details of the type. It is specified as JSON string."`
}

type updateChannelParam struct {
	ID     string          `json:"id"`
	Name   string          `json:"name,omitempty"`
	Type   string          `json:"type,omitempty"`
	Detail json.RawMessage `json:"detail,omitempty"`
}

type getChannelsCommand struct {
	Username string `short:"u" long:"username" description:"User name of channel owner."`
}

type deleteChannelCommand struct {
	Username string `short:"u" long:"username" description:"User name of channel owner."`
	ID       string `short:"i" long:"channel-id" description:"ID for identifying the channel." required:"true"`
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

func (uC *updateChannelCommand) Execute(args []string) error {
	req, err := generateUpdateChannelRequest(uC)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateUpdateChannelRequest(uC *updateChannelCommand) (*http.Request, error) {
	username, err := getUsername(uC.Username)
	if err != nil {
		return nil, err
	}

	paramStruct := &updateChannelParam{
		ID:     uC.ID,
		Name:   uC.Name,
		Type:   uC.Type,
		Detail: json.RawMessage(uC.Detail),
	}

	req, err := generateRequestWithToken(
		"PUT",
		fmt.Sprintf("v1/users/%s/channels/%s", username, uC.ID),
		paramStruct,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate update api request : %s", err)
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

func (dC *deleteChannelCommand) Execute(args []string) error {
	req, err := generateDeleteChannelRequest(dC)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateDeleteChannelRequest(dC *deleteChannelCommand) (*http.Request, error) {
	username, err := getUsername(dC.Username)
	if err != nil {
		return nil, err
	}

	req, err := generateRequestWithToken(
		"DELETE",
		fmt.Sprintf("v1/users/%s/channels/%s", username, dC.ID),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate delete api request : %s", err)
	}
	return req, nil
}
