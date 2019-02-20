package pi

import (
	"fmt"
	"net/http"
	"os"
)

type graphsCommand struct {
	Create createGraphCommand    `description:"create Graph" command:"create" subcommands-optional:"true"`
	Get    getGraphsCommand      `description:"get Graph Definitions" command:"get" subcommands-optional:"true"`
	SVG    graphSVGCommand       `description:"get SVG Graph URL" command:"svg" subcommands-optional:"true"`
	Update updateGraphCommand    `description:"update Graph Definition" command:"update" subcommands-optional:"true"`
	Detail graphDetailCommand    `description:"get Graph detail URL" command:"detail" subcommands-optional:"true"`
	Delete deleteGraphCommand    `description:"delete Graph" command:"delete" subcommands-optional:"true"`
	Pixels getGraphPixelsCommand `description:"get Graph Pixels" command:"pixels" subcommands-optional:"true"`
}

type createGraphCommand struct {
	Username       string `short:"u" long:"username" description:"User name of graph owner."`
	ID             string `short:"g" long:"graph-id" description:"ID for identifying the pixelation graph." required:"true"`
	Name           string `short:"n" long:"name" description:"The name of the pixelation graph." required:"true"`
	Unit           string `short:"u" long:"unit" description:"A unit of the quantity recorded in the pixelation graph. Ex) commit, kilogram, calory." required:"true"`
	Type           string `short:"t" long:"type" description:"The type of quantity to be handled in the graph. Only int or float are supported." choice:"int" choice:"float" required:"true"`
	Color          string `short:"c" long:"color" description:"The display color of the pixel in the pixelation graph." choice:"shibafu" choice:"momiji" choice:"sora" choice:"ichou" choice:"ajisai" choice:"kuro" required:"true"`
	Timezone       string `short:"z" long:"timezone" description:"The timezone for handling this graph"`
	SelfSufficient string `short:"s" long:"self-sufficient" description:"If SVG graph with this field 'increment' or 'decrement' is referenced, Pixel of this graph itself will be incremented or decremented." choice:"increment" choice:"decrement" choice:"none"`
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
	Username string `short:"u" long:"username" description:"User name of graph owner."`
}
type getGraphParam struct{}

type graphSVGCommand struct {
	Username string `short:"u" long:"username" description:"User name of graph owner."`
	ID       string `short:"g" long:"graph-id" description:"ID for identifying the pixelation graph." required:"true"`
	Date     string `short:"d" long:"date" description:"If you specify it in yyyyMMdd format, will create a pixelation graph dating back to the past with that day as the start date."`
	Mode     string `short:"m" long:"mode" description:"Specify the graph display mode." choice:"short"`
}

type updateGraphCommand struct {
	Username       string   `short:"u" long:"username" description:"User name of graph owner."`
	ID             string   `short:"g" long:"graph-id" description:"ID for identifying the pixelation graph." required:"true"`
	Name           string   `short:"n" long:"name" description:"The name of the pixelation graph."`
	Unit           string   `short:"i" long:"unit" description:"A unit of the quantity recorded in the pixelation graph. Ex) commit, kilogram, calory."`
	Color          string   `short:"c" long:"color" description:"The display color of the pixel in the pixelation graph." choice:"shibafu" choice:"momiji" choice:"sora" choice:"ichou" choice:"ajisai" choice:"kuro"`
	Timezone       string   `short:"z" long:"timezone" description:"The timezone for handling this graph"`
	PurgeCacheURLs []string `short:"p" long:"purge-cache-urls" description:"The URL to send the purge request to purge the cache when the graph is updated. Multiple params can be specified."`
	SelfSufficient string   `short:"s" long:"self-sufficient" description:"If SVG graph with this field 'increment' or 'decrement' is referenced, Pixel of this graph itself will be incremented or decremented." choice:"increment" choice:"decrement" choice:"none"`
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
	Username string `short:"u" long:"username" description:"User name of graph owner."`
	ID       string `short:"g" long:"graph-id" description:"ID for identifying the pixelation graph." required:"true"`
}

type deleteGraphCommand struct {
	Username string `short:"u" long:"username" description:"User name of graph owner."`
	ID       string `short:"g" long:"graph-id" description:"ID for identifying the pixelation graph." required:"true"`
}

type getGraphPixelsCommand struct {
	Username string `short:"u" long:"username" description:"User name of graph owner."`
	ID       string `short:"g" long:"graph-id" description:"ID for identifying the pixelation graph." required:"true"`
	From     string `short:"f" long:"from" description:"Specify the start position of the period."`
	To       string `short:"t" long:"to" description:"Specify the end position of the period."`
}

func (cG *createGraphCommand) Execute(args []string) error {
	req, err := generateCreateGraphRequest(cG)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateCreateGraphRequest(cG *createGraphCommand) (*http.Request, error) {
	username, err := getUsername(cG.Username)
	if err != nil {
		return nil, err
	}

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
		fmt.Sprintf("v1/users/%s/graphs", username),
		paramStruct,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate create api request : %s", err)
	}

	return req, nil
}

func (gG *getGraphsCommand) Execute(args []string) error {
	req, err := generateGetGraphsRequest(gG)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateGetGraphsRequest(gG *getGraphsCommand) (*http.Request, error) {
	username, err := getUsername(gG.Username)
	if err != nil {
		return nil, err
	}

	req, err := generateRequestWithToken(
		"GET",
		fmt.Sprintf("v1/users/%s/graphs", username),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate get api request : %s", err)
	}

	return req, nil
}

func (gS *graphSVGCommand) Execute(args []string) error {
	url, err := generateSVGUrl(gS)
	if err != nil {
		return err
	}

	fmt.Print(url)
	return nil
}

func generateSVGUrl(gS *graphSVGCommand) (string, error) {
	username, err := getUsername(gS.Username)
	if err != nil {
		return username, err
	}

	apibase := os.Getenv("PIXELA_API_BASE")
	if apibase == "" {
		apibase = "pixe.la"
	}
	url := fmt.Sprintf("https://%s/v1/users/%s/graphs/%s", apibase, username, gS.ID)

	if gS.Date != "" {
		url = fmt.Sprintf("%s?date=%s", url, gS.Date)
		if gS.Mode != "" {
			url = fmt.Sprintf("%s&mode=%s", url, gS.Mode)
		}
	} else if gS.Mode != "" {
		url = fmt.Sprintf("%s?mode=%s", url, gS.Mode)
	}
	return url, nil
}

func (uG *updateGraphCommand) Execute(args []string) error {
	req, err := generateUpdateGraphRequest(uG)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateUpdateGraphRequest(uG *updateGraphCommand) (*http.Request, error) {
	username, err := getUsername(uG.Username)
	if err != nil {
		return nil, err
	}

	if len(uG.PurgeCacheURLs) > 5 {
		return nil, fmt.Errorf("you can only specify up to five URLs for PurgeCacheURLs param")
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
		fmt.Sprintf("v1/users/%s/graphs/%s", username, uG.ID),
		paramStruct,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate update api request : %s", err)
	}

	return req, nil
}

func (gD *graphDetailCommand) Execute(args []string) error {
	username, err := getUsername(gD.Username)
	if err != nil {
		return err
	}

	apibase := os.Getenv("PIXELA_API_BASE")
	if apibase == "" {
		apibase = "pixe.la"
	}
	fmt.Printf("https://%s/v1/users/%s/graphs/%s.html", apibase, username, gD.ID)
	return nil
}

func (dG *deleteGraphCommand) Execute(args []string) error {
	req, err := generateDeleteGraphRequest(dG)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateDeleteGraphRequest(dG *deleteGraphCommand) (*http.Request, error) {
	username, err := getUsername(dG.Username)
	if err != nil {
		return nil, err
	}

	req, err := generateRequestWithToken(
		"DELETE",
		fmt.Sprintf("v1/users/%s/graphs/%s", username, dG.ID),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate delete api request : %s", err)
	}
	return req, nil
}

func (gGP *getGraphPixelsCommand) Execute(args []string) error {
	req, err := generateGetGraphPixelsRequest(gGP)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateGetGraphPixelsRequest(gGP *getGraphPixelsCommand) (*http.Request, error) {
	username, err := getUsername(gGP.Username)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("v1/users/%s/graphs/%s/pixels", username, gGP.ID)

	if gGP.From != "" {
		url = fmt.Sprintf("%s?from=%s", url, gGP.From)
		if gGP.To != "" {
			url = fmt.Sprintf("%s&to=%s", url, gGP.To)
		}
	} else if gGP.To != "" {
		url = fmt.Sprintf("%s?to=%s", url, gGP.To)
	}

	req, err := generateRequestWithToken(
		"GET",
		url,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate get api request : %s", err)
	}

	return req, nil
}
