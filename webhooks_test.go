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
		name:     "create graph - not specify id",
		input:    []string{"webhooks", "create", "--type", "increment", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "create graph - not specify type",
		input:    []string{"webhooks", "create", "--id", "test-id", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "create graph - not specify username",
		input:    []string{"webhooks", "create", "--id", "test-id", "--type", "increment"},
		exitCode: 1,
	},
	{
		name:     "create graph - invalid type",
		input:    []string{"webhooks", "create", "--id", "test-id", "--type", "none", "--username", "c-know"},
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
