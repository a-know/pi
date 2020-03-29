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

func TestGeneratePostNotificationRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	testGraphID := "test-graph"
	testID := "test-id"
	testName := "test-notification"
	testTarget := "quantity"
	testCondition := ">"
	testThreshold := "5"
	testChannelID := "channel-id"
	cmd := &postNotificationCommand{
		Username:  testUsername,
		GraphID:   testGraphID,
		ID:        testID,
		Name:      testName,
		Target:    testTarget,
		Condition: testCondition,
		Threshold: testThreshold,
		ChannelID: testChannelID,
	}

	// run
	req, err := generatePostNotificationRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "POST" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/v1/users/%s/graphs/%s/notifications", afterAPIBaseEnv, testUsername, testGraphID) {
		t.Errorf("Unexpected request path. %s", req.URL.String())
	}
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		t.Errorf("Failed to read request body. %s", err)
	}
	if string(b) != fmt.Sprintf(`{"id":"%s","name":"%s","target":"%s","condition":"%s","threshold":"%s","channelID":"%s"}`, testID, testName, testTarget, testCondition, testThreshold, testChannelID) {
		t.Errorf("Unexpected request body. %s", string(b))
	}
}
