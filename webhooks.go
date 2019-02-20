package pi

import (
	"fmt"
	"net/http"
)

type webhooksCommand struct {
	Create createWebhookCommand `description:"create a Webhook" command:"create" subcommands-optional:"true"`
	Get    getWebhooksCommand   `description:"get registered Webhooks" command:"get" subcommands-optional:"true"`
	Unvoke invokeWebhookCommand `description:"invoke Webhook" command:"invoke" subcommands-optional:"true"`
	Delete deleteWebhookCommand `description:"delete a Webhook" command:"delete" subcommands-optional:"true"`
}

type createWebhookCommand struct {
	Username string `long:"username" description:"User name of graph owner."`
	ID       string `long:"graph-id" description:"ID for identifying the pixelation graph." required:"true"`
	Type     string `long:"type" description:"Specify the behavior when this Webhook is invoked." choice:"increment" choice:"decrement" required:"true"`
}
type createWebhookParam struct {
	ID   string `json:"graphID"`
	Type string `json:"type"`
}

type getWebhooksCommand struct {
	Username string `long:"username" description:"User name of graph owner."`
}

type invokeWebhookCommand struct {
	Username    string `long:"username" description:"User name of graph owner."`
	WebhookHash string `long:"webhookHash" description:"Specify webhookHash of registered webhook." required:"true"`
}

type deleteWebhookCommand struct {
	Username    string `long:"username" description:"User name of graph owner."`
	WebhookHash string `long:"webhookHash" description:"Specify webhookHash of registered webhook." required:"true"`
}

func (cW *createWebhookCommand) Execute(args []string) error {
	req, err := generateCreateWebhookRequest(cW)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateCreateWebhookRequest(cW *createWebhookCommand) (*http.Request, error) {
	username, err := getUsername(cW.Username)
	if err != nil {
		return nil, err
	}

	paramStruct := &createWebhookParam{
		ID:   cW.ID,
		Type: cW.Type,
	}

	req, err := generateRequestWithToken(
		"POST",
		fmt.Sprintf("v1/users/%s/webhooks", username),
		paramStruct,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate create api request : %s", err)
	}

	return req, nil
}

func (gW *getWebhooksCommand) Execute(args []string) error {
	req, err := generateGetWebhooksRequest(gW)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateGetWebhooksRequest(gW *getWebhooksCommand) (*http.Request, error) {
	username, err := getUsername(gW.Username)
	if err != nil {
		return nil, err
	}

	req, err := generateRequestWithToken(
		"GET",
		fmt.Sprintf("v1/users/%s/webhooks", username),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate get api request : %s", err)
	}

	return req, nil
}

func (iW *invokeWebhookCommand) Execute(args []string) error {
	req, err := generateInvokeWebhookRequest(iW)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateInvokeWebhookRequest(iW *invokeWebhookCommand) (*http.Request, error) {
	username, err := getUsername(iW.Username)
	if err != nil {
		return nil, err
	}

	req, err := generateRequestWithToken(
		"POST",
		fmt.Sprintf("v1/users/%s/webhooks/%s", username, iW.WebhookHash),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate invoke api request : %s", err)
	}

	return req, nil
}

func (dW *deleteWebhookCommand) Execute(args []string) error {
	req, err := generateDeleteWebhookRequest(dW)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateDeleteWebhookRequest(dW *deleteWebhookCommand) (*http.Request, error) {
	username, err := getUsername(dW.Username)
	if err != nil {
		return nil, err
	}

	req, err := generateRequestWithToken(
		"DELETE",
		fmt.Sprintf("v1/users/%s/webhooks/%s", username, dW.WebhookHash),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate delete api request : %s", err)
	}

	return req, nil
}
