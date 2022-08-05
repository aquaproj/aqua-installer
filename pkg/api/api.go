package api

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

const (
	dirPermission  os.FileMode = 0o755
	filePermission os.FileMode = 0o755
)

func Install(ctx context.Context, param *Param) error {
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
	if err := Unarchive(f, resp.Body, param.OS == "windows"); err != nil {
		return fmt.Errorf("downloand and unarchive aqua: %w", err)
	}
	return nil
}
