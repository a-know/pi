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
		name:     "get channel definition - not specify username",
		input:    []string{"channels", "get"},
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
