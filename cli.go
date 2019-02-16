package pi

import (
	"fmt"
	"io"
	"log"

	flags "github.com/jessevdk/go-flags"
)

const (
	exitCodeOK = iota
	exitCodeParseFlagErr
	exitCodeErr
)

// CLI is struct for command line tool
type CLI struct {
	OutStream, ErrStream io.Writer
}

// Run the ghg
func (cli *CLI) Run(argv []string) int {
	log.SetOutput(cli.ErrStream)
	log.SetFlags(0)
	err := parseArgs(argv)
	if err != nil {
		if ferr, ok := err.(*flags.Error); ok {
			if ferr.Type == flags.ErrHelp {
				return exitCodeOK
			}
			return exitCodeParseFlagErr
		}
		return exitCodeErr
	}
	return exitCodeOK
}

type piOpts struct {
	User  userCommand  `description:"operate User" command:"user" subcommands-optional:"true"`
	Graph graphCommand `description:"operate Graph" command:"graph" subcommands-optional:"true"`
	Ver   verCommand   `description:"display version" command:"version" subcommands-optional:"true"`
}

type verCommand struct{}

func (b *verCommand) Execute(args []string) error {
	fmt.Printf("pi version: %s (rev: %s)\n", version, revision)
	return nil
}

func parseArgs(args []string) error {
	opts := &piOpts{}
	_, err := flags.ParseArgs(opts, args)
	return err
}
