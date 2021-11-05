# aqua-installer

[![Build Status](https://github.com/suzuki-shunsuke/aqua-installer/workflows/test/badge.svg)](https://github.com/suzuki-shunsuke/aqua-installer/actions)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/aqua-installer.svg)](https://github.com/suzuki-shunsuke/aqua-installer)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/aqua-installer/main/LICENSE)

Install [aqua](https://github.com/suzuki-shunsuke/aqua) quickly

* [Shell](#shell)
* [GitHub Actions](#github-actions)

## Shell

```
$ curl -sSfL https://raw.githubusercontent.com/suzuki-shunsuke/aqua-installer/v0.2.0/aqua-installer | bash
```

You can pass the following parameters.

* `-v [aqua version]`: aqua version
* `-i [aqua install path]`: aqua's install path (default: `/usr/local/bin/aqua`)

e.g.

```
$ curl -sSfL https://raw.githubusercontent.com/suzuki-shunsuke/aqua-installer/v0.2.0/aqua-installer | bash -s -- -v v0.7.12 -i bin/aqua
```

If the version isn't specified, the latest version would be installed.

## GitHub Actions

e.g.

```yaml
- uses: suzuki-shunsuke/aqua-installer@v0.2.0
  with:
    version: v0.7.12 # renovate: depName=suzuki-shunsuke/aqua
    install_path: /tmp/bin/aqua
```

Please see [int128/aqua-action](https://github.com/int128/aqua-action) too.

### Inputs

- version: (Required) installed aqua's version (e.g. `v0.1.0`)
- install_path: (Default: `/usr/local/bin/aqua`) aqua's install path

### Outputs

Nothing

## License

[MIT](LICENSE)
