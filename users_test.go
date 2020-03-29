package pi

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var userTests = []struct {
	name     string
	input    []string
	exitCode int
}{
	{
		name:     "create user - not specify token",
		input:    []string{"users", "create", "--agree-terms-of-service", "no", "--not-minor", "no", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "create user - not specify username",
		input:    []string{"users", "create", "--agree-terms-of-service", "no", "--not-minor", "no", "--token", "thisissecret"},
		exitCode: 1,
	},
	{
		name:     "create user - not specify agree-terms",
		input:    []string{"users", "create", "--not-minor", "no", "--username", "c-know", "--token", "thisissecret"},
		exitCode: 1,
	},
	{
		name:     "create user - not specify minor",
		input:    []string{"users", "create", "--agree-terms-of-service", "no", "--username", "c-know", "--token", "thisissecret"},
		exitCode: 1,
	},
	{
		name:     "create user - minor is invalid",
		input:    []string{"users", "create", "--agree-terms-of-service", "no", "--not-minor", "ok", "--username", "c-know", "--token", "thisissecret"},
		exitCode: 1,
	},
	{
		name:     "create user - agreement is invalid",
		input:    []string{"users", "create", "--agree-terms-of-service", "ok", "--not-minor", "no", "--username", "c-know", "--token", "thisissecret"},
		exitCode: 1,
	},
	{
		name:     "update user - not specify username",
		input:    []string{"users", "update", "--new-token", "thisissecret"},
		exitCode: 1,
	},
	{
		name:     "update user - not specify new-token",
		input:    []string{"users", "update", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "delete user - not specify username",
		input:    []string{"users", "delete"},
		exitCode: 1,
	},
}

func TestUser(t *testing.T) {
	for _, tt := range userTests {
		exitCode := (&CLI{
			ErrStream: ioutil.Discard,
			OutStream: ioutil.Discard,
		}).Run(tt.input)
		if exitCode != tt.exitCode {
			t.Errorf("%s(exitCode): out=%d want=%d", tt.name, exitCode, tt.exitCode)
		}
	}
}

func TestGenerateCreateUserRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testToken := "thisissecret"
	testUsername := "c-know"
	testAgreeTermsOfService := "false"
	testNotMinor := "true"
	testThanksCode := "ThisIsThanksCode"
	cmd := &createUserCommand{
		Token:               testToken,
		Username:            testUsername,
		AgreeTermsOfService: testAgreeTermsOfService,
		NotMinor:            testNotMinor,
		ThanksCode:          testThanksCode,
	}

	// run
	req, err := generateCreateUserRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "POST" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/%s", afterAPIBaseEnv, "v1/users") {
		t.Errorf("Unexpected request path. %s", req.URL.String())
	}
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		t.Errorf("Failed to read request body. %s", err)
	}
	if string(b) != fmt.Sprintf(`{"token":"%s","username":"%s","agreeTermsOfService":"%s","notMinor":"%s","thanksCode":"%s"}`, testToken, testUsername, testAgreeTermsOfService, testNotMinor, testThanksCode) {
		t.Errorf("Unexpected request body. %s", string(b))
	}
}

func TestGenerateUpdateUserRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, afterTokenEnv := prepare()

	testUsername := "c-know"
	testNewToken := "thisissecret"
	testThanksCode := "ThisIsThanksCode"
	cmd := &updateUserCommand{
		Username:   testUsername,
		NewToken:   testNewToken,
		ThanksCode: testThanksCode,
	}

	// run
	req, err := generateUpdateUserRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "PUT" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/%s/%s", afterAPIBaseEnv, "v1/users", testUsername) {
		t.Errorf("Unexpected request path. %s", req.URL.String())
	}
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		t.Errorf("Failed to read request body. %s", err)
	}
	if string(b) != fmt.Sprintf(`{"newToken":"%s","thanksCode":"%s"}`, testNewToken, testThanksCode) {
		t.Errorf("Unexpected request body. %s", string(b))
	}
	if req.Header.Get("X-USER-TOKEN") != afterTokenEnv {
		t.Errorf("Unexpected request header. %s", req.Header.Get("X-USER-TOKEN"))
	}
}

func TestGenerateUpdateUserRequestWithSomeParamas(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, afterTokenEnv := prepare()

	testUsername := "c-know"
	testThanksCode := "ThisIsThanksCode"
	cmd := &updateUserCommand{
		Username: testUsername,
		// Not specify newToken
		// NewToken:   testNewToken,
		ThanksCode: testThanksCode,
	}

	// run
	req, err := generateUpdateUserRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "PUT" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/%s/%s", afterAPIBaseEnv, "v1/users", testUsername) {
		t.Errorf("Unexpected request path. %s", req.URL.String())
	}
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		t.Errorf("Failed to read request body. %s", err)
	}
	if string(b) != fmt.Sprintf(`{"thanksCode":"%s"}`, testThanksCode) {
		t.Errorf("Unexpected request body. %s", string(b))
	}
	if req.Header.Get("X-USER-TOKEN") != afterTokenEnv {
		t.Errorf("Unexpected request header. %s", req.Header.Get("X-USER-TOKEN"))
	}
}

func TestGenerateDeleteUserRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, afterTokenEnv := prepare()

	testUsername := "c-know"
	cmd := &deleteUserCommand{
		Username: testUsername,
	}

	// run
	req, err := generateDeleteUserRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "DELETE" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/%s/%s", afterAPIBaseEnv, "v1/users", testUsername) {
		t.Errorf("Unexpected request path. %s", req.URL.String())
	}
	if req.Body != nil {
		b, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			t.Errorf("Failed to read request body. %s", err)
		}
		t.Errorf("Unexpected request body. %s", string(b))
	}
	if req.Header.Get("X-USER-TOKEN") != afterTokenEnv {
		t.Errorf("Unexpected request header. %s", req.Header.Get("X-USER-TOKEN"))
	}
}
