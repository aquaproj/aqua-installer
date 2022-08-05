package cli

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/aquaproj/aqua-installer/pkg/api"
)

const helpMessage = `aqua-installer - Install aqua

https://github.com/aquaproj/aqua-installer

Usage:
	$ aqua-installer [--aqua-version latest] [-o <install path>] [-os <OS>] [-arch <ARCH>]

Options:
	--help          show this help message
	--version       show aqua-installer version
	--aqua-version  aqua version. The default value is "latest"
	-o              File Path where aqua is installed. The default value is ${AQUA_ROOT_DIR:-${XDG_DATA_HOME:-$HOME/.local/share}/aquaproj-aqua}/bin
	-os             OS (e.g. linux, darwin, windows). By default, Go's runtime.GOOS. You can change by the environment variable AQUA_GOOS
	-arch           CPU Architecture (amd64 or arm64). By default, Go's runtime.GOARCH. You can change by the environment variable AQUA_GOARCH
`

func Run(ldflags *api.LDFlags) error {
	ctx := context.Background()

	param := &api.Param{}
	flag.StringVar(&param.AquaVersion, "aqua-version", "latest", "aqua version")
	flag.StringVar(&param.Dest, "o", "", "File Path where aqua is installed. The default value is ${AQUA_ROOT_DIR:-${XDG_DATA_HOME:-$HOME/.local/share}/aquaproj-aqua}/bin")
	flag.StringVar(&param.OS, "os", "", "OS (e.g. linux, darwin, windows). By default, Go's runtime.GOOS. You can change by the environment variable AQUA_GOOS")
	flag.StringVar(&param.Arch, "arch", "", "CPU Architecture (amd64 or arm64). By default, Go's runtime.GOARCH. You can change by the environment variable AQUA_GOARCH")
	flag.BoolVar(&param.Help, "help", false, "show this help")
	flag.BoolVar(&param.Version, "version", false, "show aqua-docker version")
	flag.Parse()

	if param.Help {
		fmt.Fprint(os.Stderr, helpMessage)
		return nil
	}

	if param.Version {
		fmt.Fprintf(os.Stderr, "%s (%s)", ldflags.Version, ldflags.Commit)
		return nil
	}

	if param.OS == "" {
		param.OS = os.Getenv("AQUA_GOOS")
		if param.OS == "" {
			param.OS = runtime.GOOS
		}
	}
	if param.Arch == "" {
		param.Arch = os.Getenv("AQUA_GOARCH")
		if param.Arch == "" {
			param.Arch = runtime.GOARCH
		}
	}

	if param.Dest == "" {
		param.Dest = filepath.Join(api.GetRootDir(), "bin", "aqua")
		if param.OS == "windows" {
			param.Dest += ".exe"
		}
	}

	log.Printf("[INFO] Installing aqua %s to %s", param.AquaVersion, param.Dest)
	if err := api.Install(ctx, param); err != nil {
		return fmt.Errorf("install aqua: %w", err)
	}
	return nil
}
