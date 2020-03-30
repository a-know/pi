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
	{
		name:     "post notification - not specify notification id",
		input:    []string{"ntf", "create", "--channel-id", "test-id", "--username", "c-know", "--name", "test-name", "--graph-id", "hoge10", "--target", "quantity", "--condition", ">", "--threshold", "5"},
		exitCode: 1,
	},
	{
		name:     "post notification - not specify channel id",
		input:    []string{"ntf", "create", "--notification-id", "ntf-id", "--username", "c-know", "--name", "test-name", "--graph-id", "hoge10", "--target", "quantity", "--condition", ">", "--threshold", "5"},
		exitCode: 1,
	},
	{
		name:     "post notification - not specify name",
		input:    []string{"ntf", "create", "--notification-id", "ntf-id", "--username", "c-know", "--channel-id", "test-id", "--graph-id", "hoge10", "--target", "quantity", "--condition", ">", "--threshold", "5"},
		exitCode: 1,
	},
	{
		name:     "post notification - not specify graph id",
		input:    []string{"ntf", "create", "--notification-id", "ntf-id", "--username", "c-know", "--channel-id", "test-id", "--name", "test-name", "--target", "quantity", "--condition", ">", "--threshold", "5"},
		exitCode: 1,
	},
	{
		name:     "post notification - not specify target",
		input:    []string{"ntf", "create", "--notification-id", "ntf-id", "--username", "c-know", "--channel-id", "test-id", "--name", "test-name", "--graph-id", "hoge10", "--condition", ">", "--threshold", "5"},
		exitCode: 1,
	},
	{
		name:     "post notification - not specify condition",
		input:    []string{"ntf", "create", "--notification-id", "ntf-id", "--username", "c-know", "--channel-id", "test-id", "--name", "test-name", "--graph-id", "hoge10", "--target", "quantity", "--threshold", "5"},
		exitCode: 1,
	},
	{
		name:     "post notification - not specify threshold",
		input:    []string{"ntf", "create", "--notification-id", "ntf-id", "--username", "c-know", "--channel-id", "test-id", "--name", "test-name", "--graph-id", "hoge10", "--target", "quantity", "--condition", ">"},
		exitCode: 1,
	},
	{
		name:     "put notification - not specify graph id",
		input:    []string{"ntf", "create", "--notification-id", "ntf-id", "--username", "c-know", "--channel-id", "test-id", "--name", "test-name", "--target", "quantity", "--condition", ">", "--threshold", "5"},
		exitCode: 1,
	},
	{
		name:     "delete notification - not specify graph id",
		input:    []string{"ntf", "delete", "--notification-id", "ntf-id", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "delete notification - not specify notification id",
		input:    []string{"ntf", "delete", "--graph-id", "hoge10", "--username", "c-know"},
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

func TestGeneratePutNotificationRequest(t *testing.T) {
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
	cmd := &putNotificationCommand{
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
	req, err := generatePutNotificationRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "PUT" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/v1/users/%s/graphs/%s/notifications/%s", afterAPIBaseEnv, testUsername, testGraphID, testID) {
		t.Errorf("Unexpected request path. %s", req.URL.String())
	}
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		t.Errorf("Failed to read request body. %s", err)
	}
	if string(b) != fmt.Sprintf(`{"name":"%s","target":"%s","condition":"%s","threshold":"%s","channelID":"%s"}`, testName, testTarget, testCondition, testThreshold, testChannelID) {
		t.Errorf("Unexpected request body. %s", string(b))
	}
}

func TestGeneratePutNotificationWithSomeParamRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	testGraphID := "test-graph"
	testID := "test-id"
	testName := "test-notification"
	testCondition := ">"
	testThreshold := "5"
	testChannelID := "channel-id"
	cmd := &putNotificationCommand{
		Username: testUsername,
		GraphID:  testGraphID,
		ID:       testID,
		Name:     testName,
		// Not specify
		// Target:    testTarget,
		Condition: testCondition,
		Threshold: testThreshold,
		ChannelID: testChannelID,
	}

	// run
	req, err := generatePutNotificationRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "PUT" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/v1/users/%s/graphs/%s/notifications/%s", afterAPIBaseEnv, testUsername, testGraphID, testID) {
		t.Errorf("Unexpected request path. %s", req.URL.String())
	}
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		t.Errorf("Failed to read request body. %s", err)
	}
	if string(b) != fmt.Sprintf(`{"name":"%s","condition":"%s","threshold":"%s","channelID":"%s"}`, testName, testCondition, testThreshold, testChannelID) {
		t.Errorf("Unexpected request body. %s", string(b))
	}
}

func TestGenerateDeleteNotificationRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	testGraphID := "test-graph"
	testID := "test-id"
	cmd := &deleteNotificationCommand{
		Username: testUsername,
		GraphID:  testGraphID,
		ID:       testID,
	}

	// run
	req, err := generateDeleteNotificationRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "DELETE" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/v1/users/%s/graphs/%s/notifications/%s", afterAPIBaseEnv, testUsername, testGraphID, testID) {
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
