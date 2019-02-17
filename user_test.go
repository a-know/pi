package pi

import (
	"io/ioutil"
	"testing"
)

var userTests = []struct {
	name     string
	input    []string
	exitCode int
}{
	{
		name:     "create user - not specify token",
		input:    []string{"users", "create", "--agree-terms-of-service", "no", "--not-minor", "no", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "create user - not specify username",
		input:    []string{"users", "create", "--agree-terms-of-service", "no", "--not-minor", "no", "--token", "thisissecret"},
		exitCode: 1,
	},
	{
		name:     "create user - not specify agree-terms",
		input:    []string{"users", "create", "--not-minor", "no", "--username", "c-know", "--token", "thisissecret"},
		exitCode: 1,
	},
	{
		name:     "create user - not specify minor",
		input:    []string{"users", "create", "--agree-terms-of-service", "no", "--username", "c-know", "--token", "thisissecret"},
		exitCode: 1,
	},
	{
		name:     "create user - minor is invalid",
		input:    []string{"users", "create", "--agree-terms-of-service", "no", "--not-minor", "ok", "--username", "c-know", "--token", "thisissecret"},
		exitCode: 1,
	},
	{
		name:     "create user - agreement is invalid",
		input:    []string{"users", "create", "--agree-terms-of-service", "ok", "--not-minor", "no", "--username", "c-know", "--token", "thisissecret"},
		exitCode: 1,
	},
	{
		name:     "update user - not psecify username",
		input:    []string{"users", "update", "--new-token", "thisissecret"},
		exitCode: 1,
	},
	{
		name:     "update user - not psecify new-token",
		input:    []string{"users", "update", "--username", "c-know"},
		exitCode: 1,
	},
	{
		name:     "delete user - not psecify username",
		input:    []string{"users", "delete"},
		exitCode: 1,
	},
}

func TestUser(t *testing.T) {
	for _, tt := range userTests {
		exitCode := (&CLI{
			ErrStream: ioutil.Discard,
			OutStream: ioutil.Discard,
		}).Run(tt.input)
		if exitCode != tt.exitCode {
			t.Errorf("%s(exitCode): out=%d want=%d", tt.name, exitCode, tt.exitCode)
		}
	}
}
