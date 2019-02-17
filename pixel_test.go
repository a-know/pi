package pi

import (
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
		input:    []string{"pixel", "post", "--username", "c-know", "--id", "test-id", "--quantity", "1"},
		exitCode: 1,
	},
	{
		name:     "post pixel - not specify quantity",
		input:    []string{"pixel", "post", "--username", "c-know", "--id", "test-id", "--date", "20190101"},
		exitCode: 1,
	},
	{
		name:     "post pixel - not specify username",
		input:    []string{"pixel", "post", "--id", "test-id", "--date", "20190101", "--quantity", "1"},
		exitCode: 1,
	},
	{
		name:     "get pixel - not specify id",
		input:    []string{"pixel", "get", "--username", "c-know", "--date", "20190101"},
		exitCode: 1,
	},
	{
		name:     "get pixel - not specify date",
		input:    []string{"pixel", "get", "--username", "c-know", "--id", "test-id"},
		exitCode: 1,
	},
	{
		name:     "get pixel - not specify username",
		input:    []string{"pixel", "get", "--id", "test-id", "--date", "20190101"},
		exitCode: 1,
	},
	{
		name:     "update pixel - not specify id",
		input:    []string{"pixel", "update", "--username", "c-know", "--date", "20190101", "--quantity", "1"},
		exitCode: 1,
	},
	{
		name:     "update pixel - not specify date",
		input:    []string{"pixel", "update", "--username", "c-know", "--id", "test-id", "--quantity", "1"},
		exitCode: 1,
	},
	{
		name:     "update pixel - not specify quantity",
		input:    []string{"pixel", "update", "--username", "c-know", "--id", "test-id", "--date", "20190101"},
		exitCode: 1,
	},
	{
		name:     "update pixel - not specify username",
		input:    []string{"pixel", "update", "--id", "test-id", "--date", "20190101", "--quantity", "1"},
		exitCode: 1,
	},
	{
		name:     "increment pixel - not specify id",
		input:    []string{"pixel", "increment", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "increment pixel - not specify username",
		input:    []string{"pixel", "increment", "--id", "test-id"},
		exitCode: 1,
	},
	{
		name:     "decrement pixel - not specify id",
		input:    []string{"pixel", "decrement", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "decrement pixel - not specify username",
		input:    []string{"pixel", "decrement", "--id", "test-id"},
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
