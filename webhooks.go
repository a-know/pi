package pi

import "fmt"

type webhooksCommand struct {
	Create createWebhookCommand `description:"create a Webhook" command:"create" subcommands-optional:"true"`
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
