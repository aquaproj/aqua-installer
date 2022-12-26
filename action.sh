#!/usr/bin/env bash

set -eu
set -o pipefail

"$GITHUB_ACTION_PATH/aqua-installer" -v "${AQUA_VERSION}"
# shellcheck disable=SC2086
aqua i $AQUA_OPTS
