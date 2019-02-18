package pi

import "fmt"

type webhooksCommand struct {
	Create createWebhookCommand `description:"create a Webhook" command:"create" subcommands-optional:"true"`
	Get    getWebhooksCommand   `description:"get registered Webhooks" command:"get" subcommands-optional:"true"`
	Unvoke invokeWebhookCommand `description:"invoke Webhook" command:"invoke" subcommands-optional:"true"`
}

type createWebhookCommand struct {
	Username string `long:"username" description:"User name of graph owner." required:"true"`
	ID       string `long:"id" description:"ID for identifying the pixelation graph." required:"true"`
	Type     string `long:"type" description:"Specify the behavior when this Webhook is invoked." choice:"increment" choice:"decrement" required:"true"`
}
type createWebhookParam struct {
	ID   string `json:"graphID"`
	Type string `json:"type"`
}

type getWebhooksCommand struct {
	Username string `long:"username" description:"User name of graph owner." required:"true"`
}

type invokeWebhookCommand struct {
	Username    string `long:"username" description:"User name of graph owner." required:"true"`
	WebhookHash string `long:"webhookHash" description:"Specify webhookHash of registered webhook." required:"true"`
}

func (cW *createWebhookCommand) Execute(args []string) error {
	paramStruct := &createWebhookParam{
		ID:   cW.ID,
		Type: cW.Type,
	}

	req, err := generateRequestWithToken(
		"POST",
		fmt.Sprintf("v1/users/%s/webhooks", cW.Username),
		paramStruct,
	)
	if err != nil {
		return fmt.Errorf("Failed to generate create api request : %s", err)
	}

	err = doRequest(req)
	return err
}

func (gW *getWebhooksCommand) Execute(args []string) error {
	req, err := generateRequestWithToken(
		"GET",
		fmt.Sprintf("v1/users/%s/webhooks", gW.Username),
		nil,
	)
	if err != nil {
		return fmt.Errorf("Failed to generate get api request : %s", err)
	}

	err = doRequest(req)
	return err
}

func (iW *invokeWebhookCommand) Execute(args []string) error {
	req, err := generateRequestWithToken(
		"POST",
		fmt.Sprintf("v1/users/%s/webhooks/%s", iW.Username, iW.WebhookHash),
		nil,
	)
	if err != nil {
		return fmt.Errorf("Failed to generate invoke api request : %s", err)
	}

	err = doRequest(req)
	return err
}
