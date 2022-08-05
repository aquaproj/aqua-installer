package action

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/aquaproj/aqua-installer/pkg/api"
	"github.com/mattn/go-shellwords"
	xhr "github.com/rocketlaunchr/gopherjs-xhr"
	githubactions "github.com/sethvargo/go-githubactions"
)

var errAquaVersionIsRequired = errors.New("aqua_version is required")

const osWindows = "windows"

func getGOOS() string {
	runnerOS := os.Getenv("RUNNER_OS")
	fmt.Fprintln(os.Stderr, "[DEBUG] RUNNER_OS: ", runnerOS)
	switch runnerOS {
	case "Linux":
		return "linux"
	case "Windows":
		return osWindows
	case "macOS":
		return "darwin"
	}
	return ""
}

func getGOARCH() string {
	runnerArch := os.Getenv("RUNNER_ARCH")
	fmt.Fprintln(os.Stderr, "[DEBUG] RUNNER_ARCH: ", runnerArch)
	switch runnerArch {
	case "X64":
		return "amd64"
	case "ARM64":
		return "arm64"
	}
	return ""
}

func Run(ldflags *api.LDFlags) error { //nolint:funlen,cyclop
	ctx := context.Background()
	aquaVersion := githubactions.GetInput("aqua_version")
	if aquaVersion == "" {
		return errAquaVersionIsRequired
	}
	param := &api.Param{
		Dest:        githubactions.GetInput("install_path"),
		AquaVersion: aquaVersion,
	}

	enableAquaInstall, err := strconv.ParseBool(githubactions.GetInput("enable_aqua_install"))
	if err != nil {
		return fmt.Errorf("parse enable_aqua_install as bool: %w", err)
	}

	var aquaOpts []string
	if aquaOptsStr := githubactions.GetInput("aqua_opts"); aquaOptsStr != "" {
		a, err := shellwords.Parse(aquaOptsStr)
		if err != nil {
			return fmt.Errorf("parse aqua_opts as shell arguments: %w", err)
		}
		aquaOpts = a
	}

	if param.OS == "" {
		param.OS = os.Getenv("AQUA_GOOS")
		if param.OS == "" {
			param.OS = getGOOS()
		}
	}
	if param.Arch == "" {
		param.Arch = os.Getenv("AQUA_GOARCH")
		if param.Arch == "" {
			param.Arch = getGOARCH()
		}
	}

	binDir := filepath.Join(api.GetRootDir(), "bin")

	if param.Dest == "" {
		param.Dest = filepath.Join(binDir, "aqua")
		if param.OS == osWindows {
			param.Dest += ".exe"
		}
	}

	log.Printf("[INFO] Installing aqua %s to %s", param.AquaVersion, param.Dest)
	if err := install(ctx, param); err != nil {
		return fmt.Errorf("install aqua: %w", err)
	}

	if !enableAquaInstall {
		return nil
	}

	githubactions.SetEnv("PATH", binDir)
	if param.OS == "windows" {
		githubactions.SetEnv("PATH", filepath.Join(api.GetRootDir(), "bat"))
	}

	if err := aquaI(ctx, githubactions.GetInput("working_directory"), aquaOpts); err != nil {
		return fmt.Errorf("run aqua i: %w", err)
	}

	return nil
}

func aquaI(ctx context.Context, workingDir string, opts []string) error {
	fmt.Fprintln(os.Stderr, "+ aqua i "+strings.Join(opts, " "))
	args := append([]string{"i"}, opts...)
	cmd := exec.CommandContext(ctx, "aqua", args...)
	cmd.Dir = workingDir
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("execute a command: aqua i: %w", err)
	}
	return nil
}

const (
	dirPermission  os.FileMode = 0o755
	filePermission os.FileMode = 0o755
)

func install(ctx context.Context, param *api.Param) error {
	if err := os.MkdirAll(filepath.Dir(param.Dest), dirPermission); err != nil {
		return fmt.Errorf("create a directory where aqua is installed: %w", err)
	}
	var u string
	if param.AquaVersion == "latest" {
		u = fmt.Sprintf("https://github.com/aquaproj/aqua/releases/latest/download/aqua_%s_%s.tar.gz", param.OS, param.Arch)
	} else {
		u = fmt.Sprintf("https://github.com/aquaproj/aqua/releases/download/%s/aqua_%s_%s.tar.gz", param.AquaVersion, param.OS, param.Arch)
	}
	log.Printf("Downloading %s", u)
	req := xhr.NewRequest("GET", u)
	req.ResponseType = xhr.Blob
	if err := req.Send(ctx, nil); err != nil {
		return fmt.Errorf("send a HTTP request: %w", err)
	}
	body := req.ResponseBytes()
	if !req.IsStatus2xx() {
		return fmt.Errorf("download aqua but status code != 2xx: status_code=%d, response_body=%s", req.Status, string(body))
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
	if err := api.Unarchive(f, bytes.NewBuffer(body), param.OS == "windows"); err != nil {
		return fmt.Errorf("downloand and unarchive aqua: %w", err)
	}
	return nil
}
