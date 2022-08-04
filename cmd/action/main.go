package main

import (
	"log"

	"github.com/aquaproj/aqua-installer/pkg/action"
	"github.com/aquaproj/aqua-installer/pkg/api"
)

var (
	version = ""
	commit  = "" //nolint:gochecknoglobals
	date    = "" //nolint:gochecknoglobals
)

func main() {
	if err := action.Run(&api.LDFlags{
		Version: version,
		Commit:  commit,
		Date:    date,
	}); err != nil {
		log.Fatal(err)
	}
}
