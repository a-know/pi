package pi

import (
	"io/ioutil"
	"testing"
)

var webhookTests = []struct {
	name     string
	input    []string
	exitCode int
}{
	{
		name:     "create webhook - not specify id",
		input:    []string{"webhooks", "create", "--type", "increment", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "create webhook - not specify type",
		input:    []string{"webhooks", "create", "--id", "test-id", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "create webhook - not specify username",
		input:    []string{"webhooks", "create", "--id", "test-id", "--type", "increment"},
		exitCode: 1,
	},
	{
		name:     "create webhook - invalid type",
		input:    []string{"webhooks", "create", "--id", "test-id", "--type", "none", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "get webhooks - not specify username",
		input:    []string{"webhooks", "get"},
		exitCode: 1,
	},
}

func TestWebhook(t *testing.T) {
	for _, tt := range webhookTests {
		exitCode := (&CLI{
			ErrStream: ioutil.Discard,
			OutStream: ioutil.Discard,
		}).Run(tt.input)
		if exitCode != tt.exitCode {
			t.Errorf("%s(exitCode): out=%d want=%d", tt.name, exitCode, tt.exitCode)
		}
	}
}
