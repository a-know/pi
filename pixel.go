package pi

import (
	"fmt"
	"net/http"
)

type pixelCommand struct {
	Post      postPixelCommand      `description:"post a Pixel" command:"post" subcommands-optional:"true"`
	Get       getPixelCommand       `description:"get a Pixel" command:"get" subcommands-optional:"true"`
	Update    updatePixelCommand    `description:"update a Pixel" command:"update" subcommands-optional:"true"`
	Increment incrementPixelCommand `description:"increment a Pixel" command:"increment" subcommands-optional:"true"`
	Decrement decrementPixelCommand `description:"decrement a Pixel" command:"decrement" subcommands-optional:"true"`
	Delete    deletePixelCommand    `description:"delete a Pixel" command:"delete" subcommands-optional:"true"`
}

type postPixelCommand struct {
	Username     string `short:"u" long:"username" description:"User name of graph owner."`
	ID           string `short:"g" long:"graph-id" description:"ID for identifying the pixelation graph." required:"true"`
	Date         string `short:"d" long:"date" description:"The date on which the quantity is to be recorded. It is specified in yyyyMMdd format." required:"true"`
	Quantity     string `short:"q" long:"quantity" description:"Specify the quantity to be registered on the specified date." required:"true"`
	OptionalData string `short:"o" long:"optional-data" description:"Additional information other than quantity. It is specified as JSON string."`
}
type postPixelParam struct {
	Date         string `json:"date"`
	Quantity     string `json:"quantity"`
	OptionalData string `json:"optionalData"`
}

type getPixelCommand struct {
	Username string `short:"u" long:"username" description:"User name of graph owner."`
	ID       string `short:"g" long:"graph-id" description:"ID for identifying the pixelation graph." required:"true"`
	Date     string `short:"d" long:"date" description:"The date on which the quantity is to be recorded. It is specified in yyyyMMdd format." required:"true"`
}

type updatePixelCommand struct {
	Username     string `short:"u" long:"username" description:"User name of graph owner."`
	ID           string `short:"g" long:"graph-id" description:"ID for identifying the pixelation graph." required:"true"`
	Date         string `short:"d" long:"date" description:"The date on which the quantity is to be recorded. It is specified in yyyyMMdd format." required:"true"`
	Quantity     string `short:"q" long:"quantity" description:"Specify the quantity to be registered on the specified date." required:"true"`
	OptionalData string `short:"o" long:"optional-data" description:"Additional information other than quantity. It is specified as JSON string."`
}
type updatePixelParam struct {
	Quantity     string `json:"quantity"`
	OptionalData string `json:"optionalData,omitempty"`
}

type incrementPixelCommand struct {
	Username string `short:"u" long:"username" description:"User name of graph owner."`
	ID       string `short:"g" long:"graph-id" description:"ID for identifying the pixelation graph." required:"true"`
}

type decrementPixelCommand struct {
	Username string `short:"u" long:"username" description:"User name of graph owner."`
	ID       string `short:"g" long:"graph-id" description:"ID for identifying the pixelation graph." required:"true"`
}

type deletePixelCommand struct {
	Username string `short:"u" long:"username" description:"User name of graph owner."`
	ID       string `short:"g" long:"graph-id" description:"ID for identifying the pixelation graph." required:"true"`
	Date     string `short:"d" long:"date" description:"The date on which the quantity is to be recorded. It is specified in yyyyMMdd format." required:"true"`
}

func (pP *postPixelCommand) Execute(args []string) error {
	req, err := generatePostPixelRequest(pP)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generatePostPixelRequest(pP *postPixelCommand) (*http.Request, error) {
	username, err := getUsername(pP.Username)
	if err != nil {
		return nil, err
	}

	paramStruct := &postPixelParam{
		Date:         pP.Date,
		Quantity:     pP.Quantity,
		OptionalData: pP.OptionalData,
	}

	req, err := generateRequestWithToken(
		"POST",
		fmt.Sprintf("v1/users/%s/graphs/%s", username, pP.ID),
		paramStruct,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate create api request : %s", err)
	}

	return req, nil
}

func (gP *getPixelCommand) Execute(args []string) error {
	req, err := generateGetPixelRequest(gP)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateGetPixelRequest(gP *getPixelCommand) (*http.Request, error) {
	username, err := getUsername(gP.Username)
	if err != nil {
		return nil, err
	}

	req, err := generateRequestWithToken(
		"GET",
		fmt.Sprintf("v1/users/%s/graphs/%s/%s", username, gP.ID, gP.Date),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate get api request : %s", err)
	}

	return req, nil
}

func (uP *updatePixelCommand) Execute(args []string) error {
	req, err := generateUpdatePixelRequest(uP)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateUpdatePixelRequest(uP *updatePixelCommand) (*http.Request, error) {
	username, err := getUsername(uP.Username)
	if err != nil {
		return nil, err
	}

	paramStruct := &updatePixelParam{
		Quantity:     uP.Quantity,
		OptionalData: uP.OptionalData,
	}

	req, err := generateRequestWithToken(
		"PUT",
		fmt.Sprintf("v1/users/%s/graphs/%s/%s", username, uP.ID, uP.Date),
		paramStruct,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate update api request : %s", err)
	}

	return req, nil
}

func (iP *incrementPixelCommand) Execute(args []string) error {
	req, err := generateIncrementPixelRequest(iP)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateIncrementPixelRequest(iP *incrementPixelCommand) (*http.Request, error) {
	username, err := getUsername(iP.Username)
	if err != nil {
		return nil, err
	}

	req, err := generateRequestWithToken(
		"PUT",
		fmt.Sprintf("v1/users/%s/graphs/%s/increment", username, iP.ID),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate increment api request : %s", err)
	}

	return req, nil
}

func (dP *decrementPixelCommand) Execute(args []string) error {
	req, err := generateDecrementPixelRequest(dP)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateDecrementPixelRequest(dP *decrementPixelCommand) (*http.Request, error) {
	username, err := getUsername(dP.Username)
	if err != nil {
		return nil, err
	}

	req, err := generateRequestWithToken(
		"PUT",
		fmt.Sprintf("v1/users/%s/graphs/%s/decrement", username, dP.ID),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate decrement api request : %s", err)
	}

	return req, nil
}

func (dP *deletePixelCommand) Execute(args []string) error {
	req, err := generateDeletePixelRequest(dP)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateDeletePixelRequest(dP *deletePixelCommand) (*http.Request, error) {
	username, err := getUsername(dP.Username)
	if err != nil {
		return nil, err
	}

	req, err := generateRequestWithToken(
		"DELETE",
		fmt.Sprintf("v1/users/%s/graphs/%s/%s", username, dP.ID, dP.Date),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate delete api request : %s", err)
	}

	return req, nil
}
