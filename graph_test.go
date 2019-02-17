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
		input:    []string{"graph", "create", "--name", "test-name", "--type", "int", "--unit", "commits", "--color", "shibafu", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "create graph - not specify name",
		input:    []string{"graph", "create", "--id", "test-id", "--type", "int", "--unit", "commits", "--color", "shibafu", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "create graph - not specify type",
		input:    []string{"graph", "create", "--id", "test-id", "--name", "test-name", "--unit", "commits", "--color", "shibafu", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "create graph - not specify unit",
		input:    []string{"graph", "create", "--id", "test-id", "--name", "test-name", "--type", "int", "--color", "shibafu", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "create graph - not specify color",
		input:    []string{"graph", "create", "--id", "test-id", "--name", "test-name", "--type", "int", "--unit", "commits", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "create graph - not specify username",
		input:    []string{"graph", "create", "--id", "test-id", "--name", "test-name", "--type", "int", "--unit", "commits", "--color", "shibafu"},
		exitCode: 1,
	},
	{
		name:     "create graph - color is invalid",
		input:    []string{"graph", "create", "--id", "test-id", "--name", "test-name", "--type", "int", "--unit", "commits", "--color", "rainbow", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "create graph - type is invalid",
		input:    []string{"graph", "create", "--id", "test-id", "--name", "test-name", "--type", "string", "--unit", "commits", "--color", "shibafu", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "create graph - self-sufficient is invalid",
		input:    []string{"graph", "create", "--id", "test-id", "--name", "test-name", "--type", "int", "--unit", "commits", "--color", "shibafu", "--username", "c-know", "--self-sufficient", "yes"},
		exitCode: 1,
	},
	{
		name:     "get graph definition - not psecify username",
		input:    []string{"graph", "get"},
		exitCode: 1,
	},
	{
		name:     "get svg graph url - not psecify username",
		input:    []string{"graph", "svg", "--id", "test-id"},
		exitCode: 1,
	},
	{
		name:     "get svg graph url - not psecify id",
		input:    []string{"graph", "svg", "--username", "c-know"},
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
