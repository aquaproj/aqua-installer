package api

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

type LDFlags struct {
	Version string
	Commit  string
	Date    string
}

type Param struct {
	AquaVersion string
	Dest        string
	OS          string
	Arch        string
	Help        bool
	Version     bool
}

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

const (
	dirPermission  os.FileMode = 0o755
	filePermission os.FileMode = 0o755
)

func Run(ldflags *LDFlags) error {
	ctx := context.Background()

	param := &Param{}
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
		param.Dest = filepath.Join(getRootDir(), "bin", "aqua")
		if param.OS == "windows" {
			param.Dest += ".exe"
		}
	}

	log.Printf("[INFO] Installing aqua %s to %s", param.AquaVersion, param.Dest)
	if err := installAqua(ctx, param); err != nil {
		return fmt.Errorf("install aqua: %w", err)
	}
	return nil
}

func installAqua(ctx context.Context, param *Param) error {
	if err := os.MkdirAll(filepath.Dir(param.Dest), dirPermission); err != nil {
		return fmt.Errorf("create a directory where aqua is installed: %w", err)
	}
	u := ""
	if param.AquaVersion == "latest" {
		u = fmt.Sprintf("https://github.com/aquaproj/aqua/releases/latest/download/aqua_%s_%s.tar.gz", param.OS, param.Arch)
	} else {
		u = fmt.Sprintf("https://github.com/aquaproj/aqua/releases/download/%s/aqua_%s_%s.tar.gz", param.AquaVersion, param.OS, param.Arch)
	}
	log.Printf("Downloading %s", u)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return fmt.Errorf("create a HTTP request: %w", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("send a HTTP request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= http.StatusBadRequest {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read a response body: %w", err)
		}
		return fmt.Errorf("download aqua but status code >= 400: status_code=%d, response_body=%s", resp.StatusCode, string(b))
	}
	f, err := os.OpenFile(param.Dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, filePermission)
	if err != nil {
		return fmt.Errorf("create a file %s: %w", param.Dest, err)
	}
	// f, err := os.Create(param.Dest)
	// if err != nil {
	// 	return fmt.Errorf("create a file %s: %w", param.Dest, err)
	// }
	defer f.Close()
	// if err := os.Chmod(param.Dest, filePermission); err != nil {
	// 	return fmt.Errorf("change a file permission %s: %w", param.Dest, err)
	// }
	if err := unarchive(f, resp.Body, param.OS == "windows"); err != nil {
		return fmt.Errorf("downloand and unarchive aqua: %w", err)
	}
	return nil
}
