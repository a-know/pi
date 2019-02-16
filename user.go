package pi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type userCommand struct {
	Create createCommand `description:"create User" command:"create" subcommands-optional:"true"`
	Update updateCommand `description:"update User Token" command:"update" subcommands-optional:"true"`
}

type createCommand struct {
	Token               string `short:"t" long:"token" description:"A token string used to authenticate as a user to be created." required:"true"`
	Username            string `short:"u" long:"username" description:"User name to be created." required:"true"`
	AgreeTermsOfService string `long:"agree-terms-of-service" description:"Specify yes or no whether you agree to the terms of service." choice:"yes" choice:"no" required:"true"`
	NotMinor            string `long:"not-minor" description:"Specify yes or no as to whether you are not a minor or if you are a minor and you have the parental consent of using pixela." choice:"yes" choice:"no" required:"true"`
}

type createParams struct {
	Token      string `json:"token"`
	Username   string `json:"username"`
	AgreeTerms string `json:"agreeTermsOfService"`
	NotMinor   string `json:"notMinor"`
}

type updateCommand struct {
	Username string `short:"u" long:"username" description:"User name to be updated." required:"true"`
	NewToken string `short:"t" long:"new-token" description:"A new authentication token for update." required:"true"`
}

type updateParams struct {
	NewToken string `json:"newToken"`
}

// func (b *userCommand) Execute(args []string) error {
// 	fmt.Println("pi user running.")
// 	return nil
// }

func (cC *createCommand) Execute(args []string) error {
	apibase := os.Getenv("PIXELA_API_BASE")
	if apibase == "" {
		apibase = "pixe.la"
	}

	paramStruct := &createParams{
		Token:      cC.Token,
		Username:   cC.Username,
		AgreeTerms: cC.AgreeTermsOfService,
		NotMinor:   cC.NotMinor,
	}
	params, err := json.Marshal(paramStruct)
	if err != nil {
		return fmt.Errorf("Failed to marshal options to json : %s", err)
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("https://%s/v1/users", apibase),
		bytes.NewBuffer(params),
	)
	if err != nil {
		return fmt.Errorf("Failed to create api request : %s", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to request api : %s", err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Failed to get response body : %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return fmt.Errorf("%s", string(b))
	}

	fmt.Println(string(b))

	return nil
}

func (uC *updateCommand) Execute(args []string) error {
	apibase := os.Getenv("PIXELA_API_BASE")
	if apibase == "" {
		apibase = "pixe.la"
	}

	token := os.Getenv("PIXELA_USER_TOKEN")
	if token == "" {
		return fmt.Errorf("Token is not set. Please specify your token by PIXELA_USER_TOKEN environment variable.")
	}

	paramStruct := &updateParams{
		NewToken: uC.NewToken,
	}
	params, err := json.Marshal(paramStruct)
	if err != nil {
		return fmt.Errorf("Failed to marshal options to json : %s", err)
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("https://%s/v1/users/%s", apibase, uC.Username),
		bytes.NewBuffer(params),
	)
	if err != nil {
		return fmt.Errorf("Failed to create api request : %s", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-USER-TOKEN", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to request api : %s", err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Failed to get response body : %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return fmt.Errorf("%s", string(b))
	}

	fmt.Println(string(b))

	return nil
}
