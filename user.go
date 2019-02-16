package pi

import (
	"fmt"
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
	paramStruct := &createParams{
		Token:      cC.Token,
		Username:   cC.Username,
		AgreeTerms: cC.AgreeTermsOfService,
		NotMinor:   cC.NotMinor,
	}

	req, err := generateRequest(
		"POST",
		"/v1/users",
		paramStruct,
	)
	if err != nil {
		return fmt.Errorf("Failed to generate create api request : %s", err)
	}

	err = doRequest(req)
	return err
}

func (uC *updateCommand) Execute(args []string) error {
	paramStruct := &updateParams{
		NewToken: uC.NewToken,
	}

	req, err := generateRequestWithToken(
		"PUT",
		fmt.Sprintf("/v1/users/%s", uC.Username),
		paramStruct,
	)
	if err != nil {
		return fmt.Errorf("Failed to generate update api request : %s", err)
	}

	err = doRequest(req)
	return err
}