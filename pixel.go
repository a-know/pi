package pi

import "fmt"

type pixelCommand struct {
	Post postPixelCommand `description:"post a Pixel" command:"post" subcommands-optional:"true"`
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
