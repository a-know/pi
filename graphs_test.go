package pi

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var graphTests = []struct {
	name     string
	input    []string
	exitCode int
}{
	{
		name:     "create graph - not specify id",
		input:    []string{"graphs", "create", "--name", "test-name", "--type", "int", "--unit", "commits", "--color", "shibafu", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "create graph - not specify name",
		input:    []string{"graphs", "create", "--id", "test-id", "--type", "int", "--unit", "commits", "--color", "shibafu", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "create graph - not specify type",
		input:    []string{"graphs", "create", "--id", "test-id", "--name", "test-name", "--unit", "commits", "--color", "shibafu", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "create graph - not specify unit",
		input:    []string{"graphs", "create", "--id", "test-id", "--name", "test-name", "--type", "int", "--color", "shibafu", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "create graph - not specify color",
		input:    []string{"graphs", "create", "--id", "test-id", "--name", "test-name", "--type", "int", "--unit", "commits", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "create graph - not specify username",
		input:    []string{"graphs", "create", "--id", "test-id", "--name", "test-name", "--type", "int", "--unit", "commits", "--color", "shibafu"},
		exitCode: 1,
	},
	{
		name:     "create graph - color is invalid",
		input:    []string{"graphs", "create", "--id", "test-id", "--name", "test-name", "--type", "int", "--unit", "commits", "--color", "rainbow", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "create graph - type is invalid",
		input:    []string{"graphs", "create", "--id", "test-id", "--name", "test-name", "--type", "string", "--unit", "commits", "--color", "shibafu", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "create graph - self-sufficient is invalid",
		input:    []string{"graphs", "create", "--id", "test-id", "--name", "test-name", "--type", "int", "--unit", "commits", "--color", "shibafu", "--username", "c-know", "--self-sufficient", "yes"},
		exitCode: 1,
	},
	{
		name:     "get graph definition - not psecify username",
		input:    []string{"graphs", "get"},
		exitCode: 1,
	},
	{
		name:     "get svg graph url - not psecify username",
		input:    []string{"graphs", "svg", "--id", "test-id"},
		exitCode: 1,
	},
	{
		name:     "get svg graph url - not psecify id",
		input:    []string{"graphs", "svg", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "get svg graph url - invalid mode",
		input:    []string{"graphs", "svg", "--id", "test-id", "--username", "c-know", "--mode", "long"},
		exitCode: 1,
	},
	{
		name:     "update graph - not specify id",
		input:    []string{"graphs", "update", "--name", "test-name", "--unit", "commits", "--color", "shibafu", "--username", "c-know", "--purge-cache-urls", "http://example.com/a", "--purge-cache-urls", "http://example.com/b"},
		exitCode: 1,
	},
	{
		name:     "update graph - not specify username",
		input:    []string{"graphs", "update", "--name", "test-name", "--unit", "commits", "--color", "shibafu", "--id", "test-id", "--purge-cache-urls", "http://example.com/a", "--purge-cache-urls", "http://example.com/b"},
		exitCode: 1,
	},
	{
		name:     "update graph - invalid color name",
		input:    []string{"graphs", "update", "--name", "test-name", "--unit", "commits", "--color", "rainbow", "--username", "c-know", "--id", "test-id", "--purge-cache-urls", "http://example.com/a", "--purge-cache-urls", "http://example.com/b"},
		exitCode: 1,
	},
	{
		name:     "update graph - invalid self-sufficient",
		input:    []string{"graphs", "update", "--name", "test-name", "--unit", "commits", "--color", "shibafu", "--username", "c-know", "--id", "test-id", "--self-sufficient", "ok", "--purge-cache-urls", "http://example.com/a", "--purge-cache-urls", "http://example.com/b"},
		exitCode: 1,
	},
	{
		name:     "update graph - purge cache urls limit over",
		input:    []string{"graphs", "update", "--name", "test-name", "--unit", "commits", "--color", "shibafu", "--username", "c-know", "--id", "test-id", "--purge-cache-urls", "http://example.com/a", "--purge-cache-urls", "http://example.com/b", "--purge-cache-urls", "http://example.com/c", "--purge-cache-urls", "http://example.com/d", "--purge-cache-urls", "http://example.com/e", "--purge-cache-urls", "http://example.com/f"},
		exitCode: 2,
	},
	{
		name:     "get graph detail url - not psecify username",
		input:    []string{"graphs", "detail", "--id", "test-id"},
		exitCode: 1,
	},
	{
		name:     "get graph detail url - not psecify id",
		input:    []string{"graphs", "detail", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "delete graph - not psecify username",
		input:    []string{"graphs", "delete", "--id", "test-id"},
		exitCode: 1,
	},
	{
		name:     "delete graph - not psecify id",
		input:    []string{"graphs", "delete", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "get graph pixels - not psecify username",
		input:    []string{"graphs", "pixels", "--id", "test-id"},
		exitCode: 1,
	},
	{
		name:     "get graph pixels - not psecify id",
		input:    []string{"graphs", "pixels", "--username", "c-know"},
		exitCode: 1,
	},
}

func TestGraph(t *testing.T) {
	for _, tt := range graphTests {
		exitCode := (&CLI{
			ErrStream: ioutil.Discard,
			OutStream: ioutil.Discard,
		}).Run(tt.input)
		if exitCode != tt.exitCode {
			t.Errorf("%s(exitCode): out=%d want=%d", tt.name, exitCode, tt.exitCode)
		}
	}
}

func TestGenerateCreateGraphRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	testID := "test-id"
	testName := "test-name"
	testUnit := "commits"
	testType := "int"
	testColor := "ajisai"
	testTimezone := "asia/Tokyo"
	testSelfSufficient := "none"
	cmd := &createGraphCommand{
		Username:       testUsername,
		ID:             testID,
		Name:           testName,
		Unit:           testUnit,
		Type:           testType,
		Color:          testColor,
		Timezone:       testTimezone,
		SelfSufficient: testSelfSufficient,
	}

	// run
	req, err := generateCreateGraphRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "POST" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/v1/users/%s/graphs", afterAPIBaseEnv, testUsername) {
		t.Errorf("Unexpected request path. %s", req.URL.String())
	}
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		t.Errorf("Failed to read request body. %s", err)
	}
	if string(b) != fmt.Sprintf(`{"id":"%s","name":"%s","unit":"%s","type":"%s","color":"%s","timezone":"%s","selfSufficient":"%s"}`, testID, testName, testUnit, testType, testColor, testTimezone, testSelfSufficient) {
		t.Errorf("Unexpected request body. %s", string(b))
	}
}

func TestGenerateGetGraphRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	cmd := &getGraphsCommand{
		Username: testUsername,
	}

	// run
	req, err := generateGetGraphsRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "GET" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("https://%s/v1/users/%s/graphs", afterAPIBaseEnv, testUsername) {
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

func TestGenerateSVGUrlNoParam(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	testID := "test-id"
	testDate := ""
	testMode := ""
	cmd := &graphSVGCommand{
		Username: testUsername,
		ID:       testID,
		Date:     testDate,
		Mode:     testMode,
	}

	// run
	url := generateSVGUrl(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if url != fmt.Sprintf("https://%s/v1/users/%s/graphs/%s", afterAPIBaseEnv, testUsername, testID) {
		t.Errorf("Unexpected url. %s", url)
	}
}

func TestGenerateSVGUrlDateSpecified(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	testID := "test-id"
	testDate := "20190101"
	testMode := ""
	cmd := &graphSVGCommand{
		Username: testUsername,
		ID:       testID,
		Date:     testDate,
		Mode:     testMode,
	}

	// run
	url := generateSVGUrl(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if url != fmt.Sprintf("https://%s/v1/users/%s/graphs/%s?date=%s", afterAPIBaseEnv, testUsername, testID, testDate) {
		t.Errorf("Unexpected url. %s", url)
	}
}

func TestGenerateSVGUrlModeSpecified(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	testID := "test-id"
	testDate := ""
	testMode := "short"
	cmd := &graphSVGCommand{
		Username: testUsername,
		ID:       testID,
		Date:     testDate,
		Mode:     testMode,
	}

	// run
	url := generateSVGUrl(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if url != fmt.Sprintf("https://%s/v1/users/%s/graphs/%s?mode=%s", afterAPIBaseEnv, testUsername, testID, testMode) {
		t.Errorf("Unexpected url. %s", url)
	}
}

func TestGenerateSVGUrlBothParamSpecified(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	testID := "test-id"
	testDate := "20190101"
	testMode := "short"
	cmd := &graphSVGCommand{
		Username: testUsername,
		ID:       testID,
		Date:     testDate,
		Mode:     testMode,
	}

	// run
	url := generateSVGUrl(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if url != fmt.Sprintf("https://%s/v1/users/%s/graphs/%s?date=%s&mode=%s", afterAPIBaseEnv, testUsername, testID, testDate, testMode) {
		t.Errorf("Unexpected url. %s", url)
	}
}

func TestGenerateUpdateGraphRequest(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, afterAPIBaseEnv, _ := prepare()

	testUsername := "c-know"
	testID := "test-id"
	testName := "test-name"
	testUnit := "commits"
	testColor := "ajisai"
	testTimezone := "asia/Tokyo"
	testSelfSufficient := "none"
	testPurgeCacheURLs := []string{"http://test.example.com/"}
	cmd := &updateGraphCommand{
		Username:       testUsername,
		ID:             testID,
		Name:           testName,
		Unit:           testUnit,
		Color:          testColor,
		Timezone:       testTimezone,
		PurgeCacheURLs: testPurgeCacheURLs,
		SelfSufficient: testSelfSufficient,
	}

	// run
	req, err := generateUpdateGraphRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "PUT" {
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
	if string(b) != fmt.Sprintf(`{"name":"%s","unit":"%s","color":"%s","timezone":"%s","purgeCacheURLs":["%s"],"selfSufficient":"%s"}`, testName, testUnit, testColor, testTimezone, testPurgeCacheURLs[0], testSelfSufficient) {
		t.Errorf("Unexpected request body. %s", string(b))
	}
}

func TestGenerateUpdateGraphRequestOver5URLsSpecified(t *testing.T) {
	// prepare
	beforeAPIBaseEnv, beforeTokenEnv, _, _ := prepare()

	testUsername := "c-know"
	testID := "test-id"
	testName := "test-name"
	testUnit := "commits"
	testColor := "ajisai"
	testTimezone := "asia/Tokyo"
	testSelfSufficient := "none"
	testPurgeCacheURLs := []string{"http://test.example.com/1", "http://test.example.com/2", "http://test.example.com/3", "http://test.example.com/4", "http://test.example.com/5", "http://test.example.com/6"}
	cmd := &updateGraphCommand{
		Username:       testUsername,
		ID:             testID,
		Name:           testName,
		Unit:           testUnit,
		Color:          testColor,
		Timezone:       testTimezone,
		PurgeCacheURLs: testPurgeCacheURLs,
		SelfSufficient: testSelfSufficient,
	}

	// run
	_, err := generateUpdateGraphRequest(cmd)

	// cleanup
	cleanup(beforeAPIBaseEnv, beforeTokenEnv)

	// assertion
	if err == nil {
		t.Errorf("Error should have occurs.")
	}
}
