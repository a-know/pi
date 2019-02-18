package pi

import (
	"fmt"
	"os"
)

type graphsCommand struct {
	Create createGraphCommand `description:"create Graph" command:"create" subcommands-optional:"true"`
	Get    getGraphsCommand   `description:"get Graph Definitions" command:"get" subcommands-optional:"true"`
	SVG    graphSVGCommand    `description:"get SVG Graph URL" command:"svg" subcommands-optional:"true"`
	Update updateGraphCommand `description:"update Graph Definition" command:"update" subcommands-optional:"true"`
	Detail graphDetailCommand `description:"get Graph detail URL" command:"detail" subcommands-optional:"true"`
	Delete deleteGraphCommand `description:"delete Graph" command:"delete" subcommands-optional:"true"`
}

type createGraphCommand struct {
	Username       string `long:"username" description:"User name of graph owner." required:"true"`
	ID             string `long:"id" description:"ID for identifying the pixelation graph." required:"true"`
	Name           string `long:"name" description:"The name of the pixelation graph." required:"true"`
	Unit           string `long:"unit" description:"A unit of the quantity recorded in the pixelation graph. Ex) commit, kilogram, calory." required:"true"`
	Type           string `long:"type" description:"The type of quantity to be handled in the graph. Only int or float are supported." choice:"int" choice:"float" required:"true"`
	Color          string `long:"color" description:"The display color of the pixel in the pixelation graph." choice:"shibafu" choice:"momiji" choice:"sora" choice:"ichou" choice:"ajisai" choice:"kuro" required:"true"`
	Timezone       string `long:"timezone" description:"The timezone for handling this graph"`
	SelfSufficient string `long:"self-sufficient" description:"If SVG graph with this field 'increment' or 'decrement' is referenced, Pixel of this graph itself will be incremented or decremented." choice:"increment" choice:"decrement" choice:"none"`
}
type createGraphParam struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Unit           string `json:"unit"`
	Type           string `json:"type"`
	Color          string `json:"color"`
	Timezone       string `json:"timezone"`
	SelfSufficient string `json:"selfSufficient"`
}

type getGraphsCommand struct {
	Username string `long:"username" description:"User name of graph owner." required:"true"`
}
type getGraphParam struct{}

type graphSVGCommand struct {
	Username string `long:"username" description:"User name of graph owner." required:"true"`
	ID       string `long:"id" description:"ID for identifying the pixelation graph." required:"true"`
	Date     string `long:"date" description:"If you specify it in yyyyMMdd format, will create a pixelation graph dating back to the past with that day as the start date."`
	Mode     string `long:"mode" description:"Specify the graph display mode." choice:"short"`
}

type updateGraphCommand struct {
	Username       string   `long:"username" description:"User name of graph owner." required:"true"`
	ID             string   `long:"id" description:"ID for identifying the pixelation graph." required:"true"`
	Name           string   `long:"name" description:"The name of the pixelation graph."`
	Unit           string   `long:"unit" description:"A unit of the quantity recorded in the pixelation graph. Ex) commit, kilogram, calory."`
	Color          string   `long:"color" description:"The display color of the pixel in the pixelation graph." choice:"shibafu" choice:"momiji" choice:"sora" choice:"ichou" choice:"ajisai" choice:"kuro"`
	Timezone       string   `long:"timezone" description:"The timezone for handling this graph"`
	PurgeCacheURLs []string `long:"purge-cache-urls" description:"he URL to send the purge request to purge the cache when the graph is updated."`
	SelfSufficient string   `long:"self-sufficient" description:"If SVG graph with this field 'increment' or 'decrement' is referenced, Pixel of this graph itself will be incremented or decremented." choice:"increment" choice:"decrement" choice:"none"`
}
type updateGraphParam struct {
	Name           string   `json:"name"`
	Unit           string   `json:"unit"`
	Color          string   `json:"color"`
	Timezone       string   `json:"timezone"`
	PurgeCacheURLs []string `json:"purgeCacheURLs"`
	SelfSufficient string   `json:"selfSufficient"`
}

type graphDetailCommand struct {
	Username string `long:"username" description:"User name of graph owner." required:"true"`
	ID       string `long:"id" description:"ID for identifying the pixelation graph." required:"true"`
}

type deleteGraphCommand struct {
	Username string `long:"username" description:"User name of graph owner." required:"true"`
	ID       string `long:"id" description:"ID for identifying the pixelation graph." required:"true"`
}

func (cG *createGraphCommand) Execute(args []string) error {
	paramStruct := &createGraphParam{
		ID:             cG.ID,
		Name:           cG.Name,
		Unit:           cG.Unit,
		Type:           cG.Type,
		Color:          cG.Color,
		Timezone:       cG.Timezone,
		SelfSufficient: cG.SelfSufficient,
	}

	req, err := generateRequestWithToken(
		"POST",
		fmt.Sprintf("v1/users/%s/graphs", cG.Username),
		paramStruct,
	)
	if err != nil {
		return fmt.Errorf("Failed to generate create api request : %s", err)
	}

	err = doRequest(req)
	return err
}

func (gG *getGraphsCommand) Execute(args []string) error {
	req, err := generateRequestWithToken(
		"GET",
		fmt.Sprintf("v1/users/%s/graphs", gG.Username),
		nil,
	)
	if err != nil {
		return fmt.Errorf("Failed to generate get api request : %s", err)
	}

	err = doRequest(req)
	return err
}

func (gS *graphSVGCommand) Execute(args []string) error {
	apibase := os.Getenv("PIXELA_API_BASE")
	if apibase == "" {
		apibase = "pixe.la"
	}
	url := fmt.Sprintf("https://%s/v1/users/%s/graphs/%s", apibase, gS.Username, gS.ID)

	if gS.Date != "" {
		url = fmt.Sprintf("%s?date=%s", url, gS.Date)
		if gS.Mode != "" {
			url = fmt.Sprintf("%s&mode=%s", url, gS.Mode)
		}
	} else if gS.Mode != "" {
		url = fmt.Sprintf("%s?mode=%s", url, gS.Mode)
	}
	fmt.Print(url)

	return nil
}

func (uG *updateGraphCommand) Execute(args []string) error {
	if len(uG.PurgeCacheURLs) > 5 {
		return fmt.Errorf("You can only specify up to five URLs for PurgeCacheURLs param.")
	}

	paramStruct := &updateGraphParam{
		Name:           uG.Name,
		Unit:           uG.Unit,
		Color:          uG.Color,
		Timezone:       uG.Timezone,
		PurgeCacheURLs: uG.PurgeCacheURLs,
		SelfSufficient: uG.SelfSufficient,
	}

	req, err := generateRequestWithToken(
		"PUT",
		fmt.Sprintf("v1/users/%s/graphs/%s", uG.Username, uG.ID),
		paramStruct,
	)
	if err != nil {
		return fmt.Errorf("Failed to generate update api request : %s", err)
	}

	err = doRequest(req)
	return err
}

func (gD *graphDetailCommand) Execute(args []string) error {
	apibase := os.Getenv("PIXELA_API_BASE")
	if apibase == "" {
		apibase = "pixe.la"
	}
	fmt.Printf("https://%s/v1/users/%s/graphs/%s.html", apibase, gD.Username, gD.ID)
	return nil
}

func (dG *deleteGraphCommand) Execute(args []string) error {
	req, err := generateRequestWithToken(
		"DELETE",
		fmt.Sprintf("v1/users/%s/graphs/%s", dG.Username, dG.ID),
		nil,
	)
	if err != nil {
		return fmt.Errorf("Failed to generate delete api request : %s", err)
	}

	err = doRequest(req)
	return err
}