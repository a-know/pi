package pi

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var pixelTests = []struct {
	name     string
	input    []string
	exitCode int
}{
	{
		name:     "post pixel - not specify id",
		input:    []string{"pixel", "post", "--username", "c-know", "--date", "20190101", "--quantity", "1"},
		exitCode: 1,
	},
	{
		name:     "post pixel - not specify date",
		input:    []string{"pixel", "post", "--username", "c-know", "--graph-id", "test-id", "--quantity", "1"},
		exitCode: 1,
	},
	{
		name:     "post pixel - not specify quantity",
		input:    []string{"pixel", "post", "--username", "c-know", "--graph-id", "test-id", "--date", "20190101"},
		exitCode: 1,
	},
	{
		name:     "post pixel - not specify username",
		input:    []string{"pixel", "post", "--graph-id", "test-id", "--date", "20190101", "--quantity", "1"},
		exitCode: 1,
	},
	{
		name:     "get pixel - not specify id",
		input:    []string{"pixel", "get", "--username", "c-know", "--date", "20190101"},
		exitCode: 1,
	},
	{
		name:     "get pixel - not specify date",
		input:    []string{"pixel", "get", "--username", "c-know", "--graph-id", "test-id"},
		exitCode: 1,
	},
	{
		name:     "get pixel - not specify username",
		input:    []string{"pixel", "get", "--graph-id", "test-id", "--date", "20190101"},
		exitCode: 1,
	},
	{
		name:     "update pixel - not specify id",
		input:    []string{"pixel", "update", "--username", "c-know", "--date", "20190101", "--quantity", "1"},
		exitCode: 1,
	},
	{
		name:     "update pixel - not specify date",
		input:    []string{"pixel", "update", "--username", "c-know", "--graph-id", "test-id", "--quantity", "1"},
		exitCode: 1,
	},
	{
		name:     "update pixel - not specify quantity",
		input:    []string{"pixel", "update", "--username", "c-know", "--graph-id", "test-id", "--date", "20190101"},
		exitCode: 1,
	},
	{
		name:     "update pixel - not specify username",
		input:    []string{"pixel", "update", "--graph-id", "test-id", "--date", "20190101", "--quantity", "1"},
		exitCode: 1,
	},
	{
		name:     "increment pixel - not specify id",
		input:    []string{"pixel", "increment", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "increment pixel - not specify username",
		input:    []string{"pixel", "increment", "--graph-id", "test-id"},
		exitCode: 1,
	},
	{
		name:     "decrement pixel - not specify id",
		input:    []string{"pixel", "decrement", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "decrement pixel - not specify username",
		input:    []string{"pixel", "decrement", "--graph-id", "test-id"},
		exitCode: 1,
	},
	{
		name:     "delete pixel - not specify id",
		input:    []string{"pixel", "delete", "--username", "c-know", "--date", "20190101"},
		exitCode: 1,
	},
	{
		name:     "delete pixel - not specify date",
		input:    []string{"pixel", "delete", "--username", "c-know", "--graph-id", "test-id"},
		exitCode: 1,
	},
	{
		name:     "delete pixel - not specify username",
		input:    []string{"pixel", "delete", "--graph-id", "test-id", "--date", "20190101"},
		exitCode: 1,
	},
}

func TestPixel(t *testing.T) {
	for _, tt := range pixelTests {
		exitCode := (&CLI{
			ErrStream: ioutil.Discard,
			OutStream: ioutil.Discard,
		}).Run(tt.input)
		if exitCode != tt.exitCode {
			t.Errorf("%s(exitCode): out=%d want=%d", tt.name, exitCode, tt.exitCode)
		}
	}
}

func TestGeneratePostPixelRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	testID := "test-id"
	testDate := "20190101"
	testQuantity := "5"
	testOptionalData := "testOptionalData"
	cmd := &postPixelCommand{
		Username:     testUsername,
		ID:           testID,
		Date:         testDate,
		Quantity:     testQuantity,
		OptionalData: testOptionalData,
	}

	// run
	req, err := generatePostPixelRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "POST" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/v1/users/%s/graphs/%s", afterAPIBaseEnv, testUsername, testID) {
		t.Errorf("Unexpected request path. %s", req.URL.String())
	}
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		t.Errorf("Failed to read request body. %s", err)
	}
	if string(b) != fmt.Sprintf(`{"date":"%s","quantity":"%s","optionalData":"%s"}`, testDate, testQuantity, testOptionalData) {
		t.Errorf("Unexpected request body. %s", string(b))
	}
}

func TestGenerateGetPixelRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	testID := "test-id"
	testDate := "20190101"
	cmd := &getPixelCommand{
		Username: testUsername,
		ID:       testID,
		Date:     testDate,
	}

	// run
	req, err := generateGetPixelRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "GET" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/v1/users/%s/graphs/%s/%s", afterAPIBaseEnv, testUsername, testID, testDate) {
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

func TestGenerateUpdatePixelRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	testID := "test-id"
	testDate := "20190101"
	testQuantity := "5"
	testOptionalData := "testOptionalData"
	cmd := &updatePixelCommand{
		Username:     testUsername,
		ID:           testID,
		Date:         testDate,
		Quantity:     testQuantity,
		OptionalData: testOptionalData,
	}

	// run
	req, err := generateUpdatePixelRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "PUT" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/v1/users/%s/graphs/%s/%s", afterAPIBaseEnv, testUsername, testID, testDate) {
		t.Errorf("Unexpected request path. %s", req.URL.String())
	}
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		t.Errorf("Failed to read request body. %s", err)
	}
	if string(b) != fmt.Sprintf(`{"quantity":"%s","optionalData":"%s"}`, testQuantity, testOptionalData) {
		t.Errorf("Unexpected request body. %s", string(b))
	}
}

func TestGenerateIncrementPixelRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	testID := "test-id"
	cmd := &incrementPixelCommand{
		Username: testUsername,
		ID:       testID,
	}

	// run
	req, err := generateIncrementPixelRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "PUT" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/v1/users/%s/graphs/%s/increment", afterAPIBaseEnv, testUsername, testID) {
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

func TestGenerateDeletePixelRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	testID := "test-id"
	testDate := "20190101"
	cmd := &deletePixelCommand{
		Username: testUsername,
		ID:       testID,
		Date:     testDate,
	}

	// run
	req, err := generateDeletePixelRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "DELETE" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/v1/users/%s/graphs/%s/%s", afterAPIBaseEnv, testUsername, testID, testDate) {
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
