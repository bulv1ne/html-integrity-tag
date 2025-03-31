package flagparse

import (
	"flag"
	"fmt"
	"html-integrity-tag/internal"
)

type CliArguments struct {
	Url     string
	ShaSize int
}

func Parse() CliArguments {
	var cliArgs CliArguments

	flag.IntVar(&cliArgs.ShaSize, "sha", 384, "SHA size (256 or 384)")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		internal.HandleError(fmt.Errorf("you must provide a URL as the first argument"))
	}
	cliArgs.Url = args[0]

	if cliArgs.ShaSize != 256 && cliArgs.ShaSize != 384 {
		internal.HandleError(fmt.Errorf("unsupported SHA size: %d, only sha256 and sha384 are supported", cliArgs.ShaSize))
	}
	return cliArgs
}
