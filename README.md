# aqua-installer

[![Build Status](https://github.com/aquaproj/aqua-installer/workflows/test/badge.svg)](https://github.com/aquaproj/aqua-installer/actions)
[![GitHub last commit](https://img.shields.io/github/last-commit/aquaproj/aqua-installer.svg)](https://github.com/aquaproj/aqua-installer)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/aquaproj/aqua-installer/main/LICENSE)

Install [aqua](https://github.com/aquaproj/aqua) quickly

* [Shell Script](#shell-script)
* [Go](#go)
* [GitHub Actions](#github-actions)

## Shell Script

You can install aqua by the following one liner.

```console
$ curl -sSfL https://raw.githubusercontent.com/aquaproj/aqua-installer/v2.0.2/aqua-installer | bash
```

But the one liner is a bit dangerous because aqua-installer may be tampered.
We recommend verifying aqua-installer's checksum before running it.

```sh
curl -sSfL -O https://raw.githubusercontent.com/aquaproj/aqua-installer/v2.0.2/aqua-installer
echo "acbb573997d664fcb8df20a8a5140dba80a4fd21f3d9e606e478e435a8945208  aqua-installer" | sha256sum -c
chmod +x aqua-installer
./aqua-installer
```

aqua-installer installs aqua to the following path.

* linux, macOS: `${AQUA_ROOT_DIR:-${XDG_DATA_HOME:-$HOME/.local/share}/aquaproj-aqua}/bin/aqua`
* windows: `${AQUA_ROOT_DIR:-$HOME/AppData/Local/aquaproj-aqua}/bin/aqua`

:warning: From aqua-installer v2, aqua-installer doesn't support specifying the install path.

You can pass the following parameters.

* `-v [aqua version]`: aqua version

e.g.

```console
$ curl -sSfL https://raw.githubusercontent.com/aquaproj/aqua-installer/v2.0.2/aqua-installer | bash -s -- -v v1.34.1
```

If the version isn't specified, the latest version would be installed.

## GitHub Actions

e.g.

```yaml
- uses: aquaproj/aqua-installer@v2.0.2
  with:
    aqua_version: v1.34.1
```

```yaml
- uses: aquaproj/aqua-installer@v2.0.2
  with:
    aqua_version: v1.34.1
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

:warning: From aqua-installer v2, aqua-installer doesn't support specifying the install path.

name | default | description
--- | --- | ---
enable_aqua_install | `"true"` | if this is `"false"`, executing `aqua i` and updating `GITHUB_PATH` are skipped
aqua_opts | `-l` | `aqua i`'s option. If you want to specify global options, please use environment variables
working_directory | `""` | working directory

### Outputs

Please see [action.yaml](action.yaml) too.

Nothing.

## How to update

https://aquaproj.github.io/docs/tutorial-extras/update-aqua

## License

[MIT](LICENSE)
