# aqua-installer

[![Build Status](https://github.com/aquaproj/aqua-installer/workflows/test/badge.svg)](https://github.com/aquaproj/aqua-installer/actions)
[![GitHub last commit](https://img.shields.io/github/last-commit/aquaproj/aqua-installer.svg)](https://github.com/aquaproj/aqua-installer)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/aquaproj/aqua-installer/main/LICENSE)

Install [aqua](https://github.com/aquaproj/aqua) quickly

* [Shell Script](#shell-script)
* [Go](#go)
* [GitHub Actions](#github-actions)

## Shell Script

```console
$ curl -sSfL https://raw.githubusercontent.com/aquaproj/aqua-installer/v1.1.2/aqua-installer | bash
```

You can pass the following parameters.

* `-v [aqua version]`: aqua version
* `-i [aqua install path]`: aqua's install path
  * default (linux, macOS): `${AQUA_ROOT_DIR:-${XDG_DATA_HOME:-$HOME/.local/share}/aquaproj-aqua}/bin/aqua`
  * default (windows): `${AQUA_ROOT_DIR:-$HOME/AppData/Local/aquaproj-aqua}/bin/aqua`

e.g.

```console
$ curl -sSfL https://raw.githubusercontent.com/aquaproj/aqua-installer/v1.1.2/aqua-installer | bash -s -- -v v1.25.0 -i bin/aqua
```

If the version isn't specified, the latest version would be installed.

### Verify the signature with Cosign

From aqua-installer v1.1.3, [aqua-installer](aqua-installer) will support the verification with [Cosign](https://docs.sigstore.dev/cosign/overview).
You have to install [Cosign](https://docs.sigstore.dev/cosign/installation).

```sh
curl -sSfL -O https://github.com/aquaproj/aqua-installer/releases/download/v1.1.3-3/aqua-installer
COSIGN_EXPERIMENTAL=1 cosign verify-blob \
  --signature https://github.com/aquaproj/aqua-installer/releases/download/v1.1.3-3/aqua-installer.sig \
  --certificate https://github.com/aquaproj/aqua-installer/releases/download/v1.1.3-3/aqua-installer.pem \
  aqua-installer
```

After the verification, you can run `aqua-installer` safely.

```sh
chmod a+x aqua-installer
./aqua-installer
```

## Go

```console
$ go run github.com/aquaproj/aqua-installer@latest -help
aqua-installer - Install aqua

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
```

e.g.

```console
$ go run github.com/aquaproj/aqua-installer@latest
```

## GitHub Actions

e.g.

```yaml
- uses: aquaproj/aqua-installer@v1.1.2
  with:
    aqua_version: v1.25.0
```

```yaml
- uses: aquaproj/aqua-installer@v1.1.2
  with:
    aqua_version: v1.25.0
    install_path: /tmp/bin/aqua
    working_directory: foo
    aqua_opts: ""
  env:
    AQUA_CONFIG: aqua-config.yaml
    AQUA_LOG_LEVEL: debug
```

### Inputs

Please see [action.yaml](action.yaml) too.

#### Required Inputs

name | description
--- | --- 
aqua_version | Installed aqua version

#### Optional Inputs

name | default | description
--- | --- | ---
install_path | `${AQUA_ROOT_DIR:-${XDG_DATA_HOME:-$HOME/.local/share}/aquaproj-aqua}/bin/aqua` | aqua's install path
enable_aqua_install | `"true"` | if this is `"false"`, executing `aqua i` and updating `GITHUB_PATH` are skipped
aqua_opts | `-l` | `aqua i`'s option. If you want to specify global options, please use environment variables
working_directory | `""` | working directory

### Outputs

Please see [action.yaml](action.yaml) too.

Nothing.

## License

[MIT](LICENSE)
