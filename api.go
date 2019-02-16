package pi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func generateRequest(method string, path string, paramStruct interface{}) (*http.Request, error) {
	apibase := os.Getenv("PIXELA_API_BASE")
	if apibase == "" {
		apibase = "pixe.la"
	}

	var reqBody io.Reader
	if paramStruct == nil {
		reqBody = nil
	} else {
		params, err := json.Marshal(paramStruct)
		if err != nil {
			return nil, fmt.Errorf("Failed to marshal options to json : %s", err)
		}
		reqBody = bytes.NewBuffer(params)
	}

	req, err := http.NewRequest(
		method,
		fmt.Sprintf("https://%s/%s", apibase, path),
		reqBody,
	)

	if err == nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, err
}

func generateRequestWithToken(method string, path string, paramStruct interface{}) (*http.Request, error) {
	req, err := generateRequest(method, path, paramStruct)

	token := os.Getenv("PIXELA_USER_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("Token is not set. Please specify your token by PIXELA_USER_TOKEN environment variable.")
	}

	req.Header.Set("X-USER-TOKEN", token)

	return req, err
}

func doRequest(req *http.Request) error {

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to request api : %s", err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Failed to get response body : %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return fmt.Errorf("%s", string(b))
	}

	fmt.Println(string(b))

	return nil
}
