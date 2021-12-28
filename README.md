# aqua-installer

[![Build Status](https://github.com/aquaproj/aqua-installer/workflows/test/badge.svg)](https://github.com/aquaproj/aqua-installer/actions)
[![GitHub last commit](https://img.shields.io/github/last-commit/aquaproj/aqua-installer.svg)](https://github.com/aquaproj/aqua-installer)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/aquaproj/aqua-installer/main/LICENSE)

Install [aqua](https://github.com/aquaproj/aqua) quickly

* [Shell](#shell)
* [GitHub Actions](#github-actions)

## Shell

```
$ curl -sSfL https://raw.githubusercontent.com/aquaproj/aqua-installer/v0.4.0/aqua-installer | bash
```

You can pass the following parameters.

* `-v [aqua version]`: aqua version
* `-i [aqua install path]`: aqua's install path (default: `/usr/local/bin/aqua`)

e.g.

```
$ curl -sSfL https://raw.githubusercontent.com/aquaproj/aqua-installer/v0.4.0/aqua-installer | bash -s -- -v v0.8.7 -i bin/aqua
```

If the version isn't specified, the latest version would be installed.

## GitHub Actions

e.g.

```yaml
- uses: aquaproj/aqua-installer@v0.4.0
  with:
    version: v0.8.7 # renovate: depName=aquaproj/aqua
    install_path: /tmp/bin/aqua
```

### Inputs

- version: (Required) installed aqua's version (e.g. `v0.1.0`)
- install_path: (Default: `/usr/local/bin/aqua`) aqua's install path

### Outputs

Nothing

## License

[MIT](LICENSE)
