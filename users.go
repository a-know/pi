package pi

import (
	"fmt"
	"net/http"
)

type usersCommand struct {
	Create createUserCommand `description:"create User" command:"create" subcommands-optional:"true"`
	Update updateUserCommand `description:"update User Token" command:"update" subcommands-optional:"true"`
	Delete deleteUserCommand `description:"delete User" command:"delete" subcommands-optional:"true"`
}

type createUserCommand struct {
	Token               string `short:"t" long:"token" description:"A token string used to authenticate as a user to be created." required:"true"`
	Username            string `short:"u" long:"username" description:"User name to be created." required:"true"`
	AgreeTermsOfService string `short:"a" long:"agree-terms-of-service" description:"Specify yes or no whether you agree to the terms of service." choice:"yes" choice:"no" required:"true"`
	NotMinor            string `short:"m" long:"not-minor" description:"Specify yes or no as to whether you are not a minor or if you are a minor and you have the parental consent of using pixela." choice:"yes" choice:"no" required:"true"`
}

type createUserParams struct {
	Token      string `json:"token"`
	Username   string `json:"username"`
	AgreeTerms string `json:"agreeTermsOfService"`
	NotMinor   string `json:"notMinor"`
}

type updateUserCommand struct {
	Username string `short:"u" long:"username" description:"User name to be updated."`
	NewToken string `short:"t" long:"new-token" description:"A new authentication token for update." required:"true"`
}

type updateUserParams struct {
	NewToken string `json:"newToken"`
}

type deleteUserCommand struct {
	Username string `short:"u" long:"username" description:"User name to be deleted."`
}

func (cC *createUserCommand) Execute(args []string) error {

	req, err := generateCreateUserRequest(cC)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateCreateUserRequest(cC *createUserCommand) (*http.Request, error) {
	paramStruct := &createUserParams{
		Token:      cC.Token,
		Username:   cC.Username,
		AgreeTerms: cC.AgreeTermsOfService,
		NotMinor:   cC.NotMinor,
	}

	req, err := generateRequest(
		"POST",
		"v1/users",
		paramStruct,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate create api request : %s", err)
	}
	return req, nil
}

func (uC *updateUserCommand) Execute(args []string) error {
	req, err := generateUpdateUserRequest(uC)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateUpdateUserRequest(uC *updateUserCommand) (*http.Request, error) {
	username, err := getUsername(uC.Username)
	if err != nil {
		return nil, err
	}

	paramStruct := &updateUserParams{
		NewToken: uC.NewToken,
	}

	req, err := generateRequestWithToken(
		"PUT",
		fmt.Sprintf("v1/users/%s", username),
		paramStruct,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate update api request : %s", err)
	}
	return req, nil
}

func (dC *deleteUserCommand) Execute(args []string) error {
	req, err := generateDeleteUserRequest(dC)
	if err != nil {
		return err
	}

	err = doRequest(req)
	return err
}

func generateDeleteUserRequest(dC *deleteUserCommand) (*http.Request, error) {
	username, err := getUsername(dC.Username)
	if err != nil {
		return nil, err
	}

	req, err := generateRequestWithToken(
		"DELETE",
		fmt.Sprintf("v1/users/%s", username),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate delete api request : %s", err)
	}
	return req, nil
}
