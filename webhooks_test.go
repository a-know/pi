package pi

import (
	"fmt"
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
	{
		name:     "invoke webhook - not specify username",
		input:    []string{"webhooks", "invoke", "--webhookHash", "hash"},
		exitCode: 1,
	},
	{
		name:     "invoke webhook - not specify WebhookHash",
		input:    []string{"webhooks", "invoke", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "delete webhook - not specify username",
		input:    []string{"webhooks", "delete", "--webhookHash", "hash"},
		exitCode: 1,
	},
	{
		name:     "delete webhook - not specify WebhookHash",
		input:    []string{"webhooks", "delete", "--username", "c-know"},
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

func TestGenerateCreateWebhookRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	testID := "test-id"
	testType := "increment"
	cmd := &createWebhookCommand{
		Username: testUsername,
		ID:       testID,
		Type:     testType,
	}

	// run
	req, err := generateCreateWebhookRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "POST" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/v1/users/%s/webhooks", afterAPIBaseEnv, testUsername) {
		t.Errorf("Unexpected request path. %s", req.URL.String())
	}
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		t.Errorf("Failed to read request body. %s", err)
	}
	if string(b) != fmt.Sprintf(`{"graphID":"%s","type":"%s"}`, testID, testType) {
		t.Errorf("Unexpected request body. %s", string(b))
	}
}

func TestGenerateGetWebhooksRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	cmd := &getWebhooksCommand{
		Username: testUsername,
	}

	// run
	req, err := generateGetWebhooksRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "GET" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/v1/users/%s/webhooks", afterAPIBaseEnv, testUsername) {
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

func TestGenerateInvokeWebhookRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	testWebhookHash := "webhookhashstring"
	cmd := &invokeWebhookCommand{
		Username:    testUsername,
		WebhookHash: testWebhookHash,
	}

	// run
	req, err := generateInvokeWebhookRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "POST" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/v1/users/%s/webhooks/%s", afterAPIBaseEnv, testUsername, testWebhookHash) {
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

func TestGenerateDeleteWebhookRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	testWebhookHash := "webhookhashstring"
	cmd := &deleteWebhookCommand{
		Username:    testUsername,
		WebhookHash: testWebhookHash,
	}

	// run
	req, err := generateDeleteWebhookRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "DELETE" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/v1/users/%s/webhooks/%s", afterAPIBaseEnv, testUsername, testWebhookHash) {
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
