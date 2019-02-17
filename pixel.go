package pi

import "fmt"

type pixelCommand struct {
	Post      postPixelCommand      `description:"post a Pixel" command:"post" subcommands-optional:"true"`
	Get       getPixelCommand       `description:"get a Pixel" command:"get" subcommands-optional:"true"`
	Update    updatePixelCommand    `description:"update a Pixel" command:"update" subcommands-optional:"true"`
	Increment incrementPixelCommand `description:"increment a Pixel" command:"increment" subcommands-optional:"true"`
	Decrement decrementPixelCommand `description:"decrement a Pixel" command:"decrement" subcommands-optional:"true"`
	Delete    deletePixelCommand    `description:"delete a Pixel" command:"delete" subcommands-optional:"true"`
}

type postPixelCommand struct {
	Username     string `long:"username" description:"User name of graph owner." required:"true"`
	ID           string `long:"id" description:"ID for identifying the pixelation graph." required:"true"`
	Date         string `long:"date" description:"The date on which the quantity is to be recorded. It is specified in yyyyMMdd format." required:"true"`
	Quantity     string `long:"quantity" description:"Specify the quantity to be registered on the specified date." required:"true"`
	OptionalData string `long:"optional-data" description:"Additional information other than quantity. It is specified as JSON string."`
}
type postPixelParam struct {
	Date         string `json:"date"`
	Quantity     string `json:"quantity"`
	OptionalData string `json:"optionalData"`
}

type getPixelCommand struct {
	Username string `long:"username" description:"User name of graph owner." required:"true"`
	ID       string `long:"id" description:"ID for identifying the pixelation graph." required:"true"`
	Date     string `long:"date" description:"The date on which the quantity is to be recorded. It is specified in yyyyMMdd format." required:"true"`
}

type updatePixelCommand struct {
	Username     string `long:"username" description:"User name of graph owner." required:"true"`
	ID           string `long:"id" description:"ID for identifying the pixelation graph." required:"true"`
	Date         string `long:"date" description:"The date on which the quantity is to be recorded. It is specified in yyyyMMdd format." required:"true"`
	Quantity     string `long:"quantity" description:"Specify the quantity to be registered on the specified date." required:"true"`
	OptionalData string `long:"optional-data" description:"Additional information other than quantity. It is specified as JSON string."`
}
type updatePixelParam struct {
	Quantity     string `json:"quantity"`
	OptionalData string `json:"optionalData"`
}

type incrementPixelCommand struct {
	Username string `long:"username" description:"User name of graph owner." required:"true"`
	ID       string `long:"id" description:"ID for identifying the pixelation graph." required:"true"`
}

type decrementPixelCommand struct {
	Username string `long:"username" description:"User name of graph owner." required:"true"`
	ID       string `long:"id" description:"ID for identifying the pixelation graph." required:"true"`
}

type deletePixelCommand struct {
	Username string `long:"username" description:"User name of graph owner." required:"true"`
	ID       string `long:"id" description:"ID for identifying the pixelation graph." required:"true"`
	Date     string `long:"date" description:"The date on which the quantity is to be recorded. It is specified in yyyyMMdd format." required:"true"`
}

func (pP *postPixelCommand) Execute(args []string) error {
	paramStruct := &postPixelParam{
		Date:         pP.Date,
		Quantity:     pP.Quantity,
		OptionalData: pP.OptionalData,
	}

	req, err := generateRequestWithToken(
		"POST",
		fmt.Sprintf("v1/users/%s/graphs/%s", pP.Username, pP.ID),
		paramStruct,
	)
	if err != nil {
		return fmt.Errorf("Failed to generate create api request : %s", err)
	}

	err = doRequest(req)
	return err
}

func (gP *getPixelCommand) Execute(args []string) error {
	req, err := generateRequestWithToken(
		"GET",
		fmt.Sprintf("v1/users/%s/graphs/%s/%s", gP.Username, gP.ID, gP.Date),
		nil,
	)
	if err != nil {
		return fmt.Errorf("Failed to generate get api request : %s", err)
	}

	err = doRequest(req)
	return err
}

func (uP *updatePixelCommand) Execute(args []string) error {
	paramStruct := &updatePixelParam{
		Quantity:     uP.Quantity,
		OptionalData: uP.OptionalData,
	}

	req, err := generateRequestWithToken(
		"PUT",
		fmt.Sprintf("v1/users/%s/graphs/%s/%s", uP.Username, uP.ID, uP.Date),
		paramStruct,
	)
	if err != nil {
		return fmt.Errorf("Failed to generate update api request : %s", err)
	}

	err = doRequest(req)
	return err
}

func (iP *incrementPixelCommand) Execute(args []string) error {
	req, err := generateRequestWithToken(
		"PUT",
		fmt.Sprintf("v1/users/%s/graphs/%s/increment", iP.Username, iP.ID),
		nil,
	)
	if err != nil {
		return fmt.Errorf("Failed to generate increment api request : %s", err)
	}

	err = doRequest(req)
	return err
}

func (dP *decrementPixelCommand) Execute(args []string) error {
	req, err := generateRequestWithToken(
		"PUT",
		fmt.Sprintf("v1/users/%s/graphs/%s/decrement", dP.Username, dP.ID),
		nil,
	)
	if err != nil {
		return fmt.Errorf("Failed to generate decrement api request : %s", err)
	}

	err = doRequest(req)
	return err
}

func (dP *deletePixelCommand) Execute(args []string) error {
	req, err := generateRequestWithToken(
		"DELETE",
		fmt.Sprintf("v1/users/%s/graphs/%s/%s", dP.Username, dP.ID, dP.Date),
		nil,
	)
	if err != nil {
		return fmt.Errorf("Failed to generate delete api request : %s", err)
	}

	err = doRequest(req)
	return err
}
