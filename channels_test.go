package pi

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var channelTests = []struct {
	name     string
	input    []string
	exitCode int
}{
	{
		name:     "create channel - not specify username",
		input:    []string{"channels", "create", "--name", "test-name", "--channel-id", "test-id", "--type", "slack", "--detail", `{"url":"https://hooks.slack.com/services/T035DA4QD/B06LMAV40/xxxx","userName":"Pixela Notification","channelName":"pixela-notify"}`},
		exitCode: 1,
	},
	{
		name:     "create channel - not specify id",
		input:    []string{"channels", "create", "--username", "c-know", "--name", "test-name", "--type", "slack", "--detail", `{"url":"https://hooks.slack.com/services/T035DA4QD/B06LMAV40/xxxx","userName":"Pixela Notification","channelName":"pixela-notify"}`},
		exitCode: 1,
	},
	{
		name:     "create channel - not specify name",
		input:    []string{"channels", "create", "--username", "c-know", "--channel-id", "test-id", "--type", "slack", "--detail", `{"url":"https://hooks.slack.com/services/T035DA4QD/B06LMAV40/xxxx","userName":"Pixela Notification","channelName":"pixela-notify"}`},
		exitCode: 1,
	},
	{
		name:     "create channel - not specify type",
		input:    []string{"channels", "create", "--username", "c-know", "--channel-id", "test-id", "--name", "test-name", "--detail", `{"url":"https://hooks.slack.com/services/T035DA4QD/B06LMAV40/xxxx","userName":"Pixela Notification","channelName":"pixela-notify"}`},
		exitCode: 1,
	},
	{
		name:     "create channel - not specify detail",
		input:    []string{"channels", "create", "--username", "c-know", "--channel-id", "test-id", "--name", "test-name", "--type", "slack"},
		exitCode: 1,
	},
	{
		name:     "update channel - not specify username",
		input:    []string{"channels", "update", "--name", "test-name", "--channel-id", "test-id", "--type", "slack", "--detail", `{"url":"https://hooks.slack.com/services/T035DA4QD/B06LMAV40/xxxx","userName":"Pixela Notification","channelName":"pixela-notify"}`},
		exitCode: 1,
	},
	{
		name:     "get channel definition - not specify username",
		input:    []string{"channels", "get"},
		exitCode: 1,
	},
	{
		name:     "delete channel definition - not specify username",
		input:    []string{"channels", "delete", "--channel-id", "test-id"},
		exitCode: 1,
	},
	{
		name:     "delete channel definition - not specify channel-id",
		input:    []string{"channels", "delete", "--username", "c-know"},
		exitCode: 1,
	},
}

func TestChannel(t *testing.T) {
	for _, tt := range channelTests {
		exitCode := (&CLI{
			ErrStream: ioutil.Discard,
			OutStream: ioutil.Discard,
		}).Run(tt.input)
		if exitCode != tt.exitCode {
			t.Errorf("%s(exitCode): out=%d want=%d", tt.name, exitCode, tt.exitCode)
		}
	}
}

func TestGenerateCreateChannelRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	testID := "test-id"
	testName := "test-name"
	testType := "slack"
	testDetail := `{"url":"https://hooks.slack.com/services/T035DA4QD/B06LMAV40/xxxx","userName":"Pixela Notification","channelName":"pixela-notify"}`
	cmd := &createChannelCommand{
		Username: testUsername,
		ID:       testID,
		Name:     testName,
		Type:     testType,
		Detail:   testDetail,
	}

	// run
	req, err := generateCreateChannelRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "POST" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/v1/users/%s/channels", afterAPIBaseEnv, testUsername) {
		t.Errorf("Unexpected request path. %s", req.URL.String())
	}
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		t.Errorf("Failed to read request body. %s", err)
	}
	if string(b) != fmt.Sprintf(`{"id":"%s","name":"%s","type":"%s","detail":%s}`, testID, testName, testType, testDetail) {
		t.Errorf("Unexpected request body. %s", string(b))
	}
}

func TestGenerateUpdateChannelRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	testID := "test-id"
	testName := "test-name"
	testType := "slack"
	testDetail := `{"url":"https://hooks.slack.com/services/T035DA4QD/B06LMAV40/xxxx","userName":"Pixela Notification","channelName":"pixela-notify"}`
	cmd := &updateChannelCommand{
		Username: testUsername,
		ID:       testID,
		Name:     testName,
		Type:     testType,
		Detail:   testDetail,
	}

	// run
	req, err := generateUpdateChannelRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "PUT" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/v1/users/%s/channels/%s", afterAPIBaseEnv, testUsername, testID) {
		t.Errorf("Unexpected request path. %s", req.URL.String())
	}
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		t.Errorf("Failed to read request body. %s", err)
	}
	if string(b) != fmt.Sprintf(`{"id":"%s","name":"%s","type":"%s","detail":%s}`, testID, testName, testType, testDetail) {
		t.Errorf("Unexpected request body. %s", string(b))
	}
}

func TestGenerateUpdateChannelWithSomeParamsRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	testID := "test-id"
	testName := "test-name"
	cmd := &updateChannelCommand{
		Username: testUsername,
		ID:       testID,
		Name:     testName,
	}

	// run
	req, err := generateUpdateChannelRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "PUT" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/v1/users/%s/channels/%s", afterAPIBaseEnv, testUsername, testID) {
		t.Errorf("Unexpected request path. %s", req.URL.String())
	}
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		t.Errorf("Failed to read request body. %s", err)
	}
	if string(b) != fmt.Sprintf(`{"id":"%s","name":"%s"}`, testID, testName) {
		t.Errorf("Unexpected request body. %s", string(b))
	}
}

func TestGenerateGetChannelRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	cmd := &getChannelsCommand{
		Username: testUsername,
	}

	// run
	req, err := generateGetChannelsRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "GET" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/v1/users/%s/channels", afterAPIBaseEnv, testUsername) {
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

func TestGenerateDeleteChannelRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	testID := "test-id"
	cmd := &deleteChannelCommand{
		Username: testUsername,
		ID:       testID,
	}

	// run
	req, err := generateDeleteChannelRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "DELETE" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/v1/users/%s/channels/%s", afterAPIBaseEnv, testUsername, testID) {
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
