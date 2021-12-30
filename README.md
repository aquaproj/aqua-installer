# aqua-installer

[![Build Status](https://github.com/aquaproj/aqua-installer/workflows/test/badge.svg)](https://github.com/aquaproj/aqua-installer/actions)
[![GitHub last commit](https://img.shields.io/github/last-commit/aquaproj/aqua-installer.svg)](https://github.com/aquaproj/aqua-installer)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/aquaproj/aqua-installer/main/LICENSE)

Install [aqua](https://github.com/aquaproj/aqua) quickly

* [Shell Script](#shell-script)
* [GitHub Actions](#github-actions)

## Shell Script

```console
$ curl -sSfL https://raw.githubusercontent.com/aquaproj/aqua-installer/v0.4.0/aqua-installer | bash
```

You can pass the following parameters.

* `-v [aqua version]`: aqua version
* `-i [aqua install path]`: aqua's install path (default: `/usr/local/bin/aqua`)

e.g.

```console
$ curl -sSfL https://raw.githubusercontent.com/aquaproj/aqua-installer/v0.4.0/aqua-installer | bash -s -- -v v0.8.10 -i bin/aqua
```

If the version isn't specified, the latest version would be installed.

## GitHub Actions

e.g.

```yaml
- uses: aquaproj/aqua-installer@v0.4.0
  with:
    aqua_version: v0.8.7
    install_path: /tmp/bin/aqua
```

### Inputs, Outputs

Please see [action.yaml](action.yaml)

## License

[MIT](LICENSE)
