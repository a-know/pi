package pi

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestGenerateRequestApiBaseEnvExist(t *testing.T) {
	// prepare
	beforeEnv := os.Getenv("PIXELA_API_BASE")
	afterEnv := "pixela.example.com"
	os.Setenv("PIXELA_API_BASE", afterEnv)
	testToken := "thisissecret"
	testUsername := "c-know"
	testAgreement := "false"
	testMinor := "true"
	testParamStruct := &createUserParams{
		Token:      testToken,
		Username:   testUsername,
		AgreeTerms: testAgreement,
		NotMinor:   testMinor,
	}

	// test call
	testMethod := "POST"
	testPath := "v1/users/c-know"
	req, err := generateRequest(testMethod, testPath, testParamStruct)

	// cleanup
	os.Setenv("PIXELA_API_BASE", beforeEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != testMethod {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/%s", afterEnv, testPath) {
		t.Errorf("Unexpected request path. %s", req.URL.String())
	}
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		t.Errorf("Failed to read request body. %s", err)
	}
	if string(b) != fmt.Sprintf("{\"token\":\"%s\",\"username\":\"%s\",\"agreeTermsOfService\":\"%s\",\"notMinor\":\"%s\"}", testToken, testUsername, testAgreement, testMinor) {
		t.Errorf("Unexpected request body. %s", string(b))
	}
	if req.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Unexpected request header. %s", req.Header.Get("Content-Type"))
	}
}

func TestGenerateRequestApiBaseEnvNotExist(t *testing.T) {
	// prepare
	beforeEnv := os.Getenv("PIXELA_API_BASE")
	afterEnv := ""
	os.Setenv("PIXELA_API_BASE", afterEnv)
	testToken := "thisissecret"
	testUsername := "c-know"
	testAgreement := "false"
	testMinor := "true"
	testParamStruct := &createUserParams{
		Token:      testToken,
		Username:   testUsername,
		AgreeTerms: testAgreement,
		NotMinor:   testMinor,
	}

	// test call
	testMethod := "POST"
	testPath := "v1/users/c-know"
	req, err := generateRequest(testMethod, testPath, testParamStruct)

	// cleanup
	os.Setenv("PIXELA_API_BASE", beforeEnv)

	// assertion
	if err != nil {
		t.Errorf("Failed to generate request. %s", err)
	}

	if req.URL.String() != fmt.Sprintf("https://pixe.la/%s", testPath) {
		t.Errorf("Unexpected request path. %s", req.URL.String())
	}
}
func TestGenerateRequestBodyIsNil(t *testing.T) {
	// prepare
	testMethod := "POST"
	testPath := "v1/users/c-know"

	// test call
	req, err := generateRequest(testMethod, testPath, nil)

	// cleanup
	// nop

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
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

func TestGenerateRequestWithToken(t *testing.T) {
	// prepare
	beforeEnv := os.Getenv("PIXELA_USER_TOKEN")
	afterEnv := "thisissecret"
	os.Setenv("PIXELA_USER_TOKEN", afterEnv)

	// test call
	testMethod := "POST"
	testPath := "v1/users/c-know"
	req, err := generateRequestWithToken(testMethod, testPath, nil)

	// cleanup
	os.Setenv("PIXELA_USER_TOKEN", beforeEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Header.Get("X-USER-TOKEN") != afterEnv {
		t.Errorf("Unexpected request header. %s", req.Header.Get("X-USER-TOKEN"))
	}
}

func TestGenerateRequestWithTokenNoToken(t *testing.T) {
	// prepare
	beforeEnv := os.Getenv("PIXELA_USER_TOKEN")
	afterEnv := ""
	os.Setenv("PIXELA_USER_TOKEN", afterEnv)

	// test call
	testMethod := "POST"
	testPath := "v1/users/c-know"
	_, err := generateRequestWithToken(testMethod, testPath, nil)

	// cleanup
	os.Setenv("PIXELA_USER_TOKEN", beforeEnv)

	// assertion
	if err == nil {
		t.Errorf("Error should has occurs.")
	}
}
