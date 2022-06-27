#!/usr/bin/env bash

set -eu
set -o pipefail

if [ -n "${CLIVM_INSTALL_PATH:-}" ]; then
	"$GITHUB_ACTION_PATH/clivm-installer" -v "${CLIVM_VERSION}" -i "${CLIVM_INSTALL_PATH}"
	# shellcheck disable=SC2086
	"${CLIVM_INSTALL_PATH}" i $CLIVM_OPTS
	exit 0
fi

"$GITHUB_ACTION_PATH/clivm-installer" -v "${CLIVM_VERSION}"
# shellcheck disable=SC2086
clivm i $CLIVM_OPTS
