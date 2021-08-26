# aqua-installer

[![Build Status](https://github.com/suzuki-shunsuke/aqua-installer/workflows/test/badge.svg)](https://github.com/suzuki-shunsuke/aqua-installer/actions)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/aqua-installer.svg)](https://github.com/suzuki-shunsuke/aqua-installer)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/aqua-installer/main/LICENSE)

Install [aqua](https://github.com/suzuki-shunsuke/aqua) quickly

## How to install

```
$ curl -sSfL https://raw.githubusercontent.com/suzuki-shunsuke/aqua-installer/v0.1.0-0/aqua-installer | bash
```

You can pass the following parameters.

* `-v [aqua version]`: aqua version
* `-i [aqua install path]`: aqua's install path (default: `/usr/local/bin/aqua`)

e.g.

```
$ curl -sSfL https://raw.githubusercontent.com/suzuki-shunsuke/aqua-installer/v0.1.0-0/aqua-installer | bash -s -- -v v0.1.0-3 -i bin/aqua
```

## License

[MIT](LICENSE)
