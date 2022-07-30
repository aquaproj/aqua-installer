package main

import (
	"log"

	"github.com/aquaproj/aqua-installer/pkg/api"
)

var (
	version = ""
	commit  = "" //nolint:gochecknoglobals
	date    = "" //nolint:gochecknoglobals
)

func main() {
	if err := api.Run(&api.LDFlags{
		Version: version,
		Commit:  commit,
		Date:    date,
	}); err != nil {
		log.Fatal(err)
	}
}
