package pi

import (
	"fmt"
	"net/http"
)

type notificationsCommand struct {
	Post postNotificationCommand `description:"post Notification setting" command:"create" subcommands-optional:"true"`
	Get  getNotificationsCommand `description:"get Notifications" command:"get" subcommands-optional:"true"`
}

type getNotificationsCommand struct {
	Username string `short:"u" long:"username" description:"User name of owner."`
	GraphID  string `short:"g" long:"graph-id" description:"ID for identifying the graph." required:"true"`
}

type postNotificationCommand struct {
	Username  string `short:"u" long:"username" description:"User name of graph owner."`
	GraphID   string `short:"g" long:"graph-id" description:"ID for identifying the graph." required:"true"`
	ID        string `short:"i" long:"notifiation-id" description:"ID for identifying the notification setting." required:"true"`
	Name      string `short:"n" long:"name" description:"It is the name of the notification settings." required:"true"`
	Target    string `short:"t" long:"target" description:"Specify the target to be notified." required:"true"`
	Condition string `short:"d" long:"condition" description:"Specify the condition used to judge whether to notify or not." required:"true"`
	Threshold string `short:"s" long:"threshold" description:"Specify the threshold value for deciding whether to notify or not. The number must match the graph type(int or float)." required:"true"`
	ChannelID string `short:"c" long:"channel-id" description:"Specify the ID of the channel to be notified." required:"true"`
}

type postNotificationParam struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Target    string `json:"target"`
	Condition string `json:"condition"`
	Threshold string `json:"threshold"`
	ChannelID string `json:"channelID"`
}

type putNotificationCommand struct {
	Username  string `short:"u" long:"username" description:"User name of graph owner."`
	GraphID   string `short:"g" long:"graph-id" description:"ID for identifying the graph." required:"true"`
	ID        string `short:"i" long:"notifiation-id" description:"ID for identifying the notification setting." required:"true"`
	Name      string `short:"n" long:"name" description:"It is the name of the notification settings."`
	Target    string `short:"t" long:"target" description:"Specify the target to be notified."`
	Condition string `short:"d" long:"condition" description:"Specify the condition used to judge whether to notify or not."`
	Threshold string `short:"s" long:"threshold" description:"Specify the threshold value for deciding whether to notify or not. The number must match the graph type(int or float)."`
	ChannelID string `short:"c" long:"channel-id" description:"Specify the ID of the channel to be notified."`
}

type putNotificationParam struct {
	Name      string `json:"name,omitempty"`
	Target    string `json:"target,omitempty"`
	Condition string `json:"condition,omitempty"`
	Threshold string `json:"threshold,omitempty"`
	ChannelID string `json:"channelID,omitempty"`
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

func (pN *postNotificationCommand) Execute(args []string) error {
	req, err := generatePostNotificationRequest(pN)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generatePostNotificationRequest(pN *postNotificationCommand) (*http.Request, error) {
	username, err := getUsername(pN.Username)
	if err != nil {
		return nil, err
	}

	paramStruct := &postNotificationParam{
		ID:        pN.ID,
		Name:      pN.Name,
		Target:    pN.Target,
		Condition: pN.Condition,
		Threshold: pN.Threshold,
		ChannelID: pN.ChannelID,
	}

	req, err := generateRequestWithToken(
		"POST",
		fmt.Sprintf("v1/users/%s/graphs/%s/notifications", username, pN.GraphID),
		paramStruct,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate create api request : %s", err)
	}

	return req, nil
}

func (pN *putNotificationCommand) Execute(args []string) error {
	req, err := generatePutNotificationRequest(pN)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generatePutNotificationRequest(pN *putNotificationCommand) (*http.Request, error) {
	username, err := getUsername(pN.Username)
	if err != nil {
		return nil, err
	}

	paramStruct := &putNotificationParam{
		Name:      pN.Name,
		Target:    pN.Target,
		Condition: pN.Condition,
		Threshold: pN.Threshold,
		ChannelID: pN.ChannelID,
	}

	req, err := generateRequestWithToken(
		"PUT",
		fmt.Sprintf("v1/users/%s/graphs/%s/notifications/%s", username, pN.GraphID, pN.ID),
		paramStruct,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate create api request : %s", err)
	}

	return req, nil
}
