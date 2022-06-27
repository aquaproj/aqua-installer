#!/usr/bin/env bash

set -eu
set -o pipefail

if [ -n "${AQUA_INSTALL_PATH:-}" ]; then
	"$GITHUB_ACTION_PATH/clivm-installer" -v "${AQUA_VERSION}" -i "${AQUA_INSTALL_PATH}"
	# shellcheck disable=SC2086
	"${AQUA_INSTALL_PATH}" i $AQUA_OPTS
	exit 0
fi

"$GITHUB_ACTION_PATH/clivm-installer" -v "${AQUA_VERSION}"
# shellcheck disable=SC2086
clivm i $AQUA_OPTS
