package pi

import (
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
