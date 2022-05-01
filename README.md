# aqua-installer

[![Build Status](https://github.com/aquaproj/aqua-installer/workflows/test/badge.svg)](https://github.com/aquaproj/aqua-installer/actions)
[![GitHub last commit](https://img.shields.io/github/last-commit/aquaproj/aqua-installer.svg)](https://github.com/aquaproj/aqua-installer)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/aquaproj/aqua-installer/main/LICENSE)

Install [aqua](https://github.com/aquaproj/aqua) quickly

* [Shell Script](#shell-script)
* [GitHub Actions](#github-actions)

## Shell Script

```console
$ curl -sSfL https://raw.githubusercontent.com/aquaproj/aqua-installer/v1.0.0/aqua-installer | bash
```

You can pass the following parameters.

* `-v [aqua version]`: aqua version
* `-i [aqua install path]`: aqua's install path (default: `${AQUA_ROOT_DIR:-${XDG_DATA_HOME:-$HOME/.local/share}/aquaproj-aqua}/bin/aqua`)

e.g.

```console
$ curl -sSfL https://raw.githubusercontent.com/aquaproj/aqua-installer/v1.0.0/aqua-installer | bash -s -- -v v1.5.0 -i bin/aqua
```

If the version isn't specified, the latest version would be installed.

## GitHub Actions

e.g.

```yaml
- uses: aquaproj/aqua-installer@v1.0.0
  with:
    aqua_version: v1.5.0
```

```yaml
- uses: aquaproj/aqua-installer@v1.0.0
  with:
    aqua_version: v1.5.0
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
install_path | ${AQUA_ROOT_DIR:-${XDG_DATA_HOME:-$HOME/.local/share}/aquaproj-aqua}/bin/aqua | aqua's install path
enable_aqua_install | `"true"` | if this is `"false"`, executing `aqua i` and updating `GITHUB_PATH` are skipped
aqua_opts | `-l` | `aqua i`'s option. If you want to specify global options, please use environment variables
working_directory | `""` | working directory

### Outputs

Please see [action.yaml](action.yaml) too.

Nothing.

## License

[MIT](LICENSE)
