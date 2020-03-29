package pi

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var notificationsTests = []struct {
	name     string
	input    []string
	exitCode int
}{
	{
		name:     "get notifications - not specify graph id",
		input:    []string{"ntf", "get", "--username", "c-know"},
		exitCode: 1,
	},
}

func TestNotifications(t *testing.T) {
	for _, tt := range notificationsTests {
		exitCode := (&CLI{
			ErrStream: ioutil.Discard,
			OutStream: ioutil.Discard,
		}).Run(tt.input)
		if exitCode != tt.exitCode {
			t.Errorf("%s(exitCode): out=%d want=%d", tt.name, exitCode, tt.exitCode)
		}
	}
}

func TestGenerateGetNotificationsRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	testGraphID := "test-graph"
	cmd := &getNotificationsCommand{
		Username: testUsername,
		GraphID:  testGraphID,
	}

	// run
	req, err := generateGetNotificationsRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "GET" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/v1/users/%s/graphs/%s/notifications", afterAPIBaseEnv, testUsername, testGraphID) {
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
}
