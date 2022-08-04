package main

import (
	"log"

	"github.com/aquaproj/aqua-installer/pkg/api"
	"github.com/aquaproj/aqua-installer/pkg/cli"
)

var (
	version = ""
	commit  = "" //nolint:gochecknoglobals
	date    = "" //nolint:gochecknoglobals
)

func main() {
	if err := cli.Run(&api.LDFlags{
		Version: version,
		Commit:  commit,
		Date:    date,
	}); err != nil {
		log.Fatal(err)
	}
}
