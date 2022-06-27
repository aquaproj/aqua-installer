# clivm-installer

[![Build Status](https://github.com/clivm/clivm-installer/workflows/test/badge.svg)](https://github.com/clivm/clivm-installer/actions)
[![GitHub last commit](https://img.shields.io/github/last-commit/clivm/clivm-installer.svg)](https://github.com/clivm/clivm-installer)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/clivm/clivm-installer/main/LICENSE)

Install [clivm](https://github.com/clivm/clivm) quickly

* [Shell Script](#shell-script)
* [GitHub Actions](#github-actions)

## Shell Script

```console
$ curl -sSfL https://raw.githubusercontent.com/clivm/clivm-installer/v1.0.0/clivm-installer | bash
```

You can pass the following parameters.

* `-v [clivm version]`: clivm version
* `-i [clivm install path]`: clivm's install path
  * default (linux, macOS): `${AQUA_ROOT_DIR:-${XDG_DATA_HOME:-$HOME/.local/share}/clivm}/bin/clivm`
  * default (windows): `${AQUA_ROOT_DIR:-$HOME/AppData/Local/clivm}/bin/clivm`

e.g.

```console
$ curl -sSfL https://raw.githubusercontent.com/clivm/clivm-installer/v1.0.0/clivm-installer | bash -s -- -v v1.14.0 -i bin/clivm
```

If the version isn't specified, the latest version would be installed.

## GitHub Actions

e.g.

```yaml
- uses: clivm/clivm-installer@v1.0.0
  with:
    clivm_version: v1.14.0
```

```yaml
- uses: clivm/clivm-installer@v1.0.0
  with:
    clivm_version: v1.14.0
    install_path: /tmp/bin/clivm
    working_directory: foo
    clivm_opts: ""
  env:
    AQUA_CONFIG: clivm-config.yaml
    AQUA_LOG_LEVEL: debug
```

### Inputs

Please see [action.yaml](action.yaml) too.

#### Required Inputs

name | description
--- | --- 
clivm_version | Installed clivm version

#### Optional Inputs

name | default | description
--- | --- | ---
install_path | ${AQUA_ROOT_DIR:-${XDG_DATA_HOME:-$HOME/.local/share}/clivm}/bin/clivm | clivm's install path
enable_clivm_install | `"true"` | if this is `"false"`, executing `clivm i` and updating `GITHUB_PATH` are skipped
clivm_opts | `-l` | `clivm i`'s option. If you want to specify global options, please use environment variables
working_directory | `""` | working directory

### Outputs

Please see [action.yaml](action.yaml) too.

Nothing.

## License

[MIT](LICENSE)
