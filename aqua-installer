#!/usr/bin/env bash
# aqua-installer - shell script to install aqua
# https://github.com/aquaproj/aqua-installer

set -eu

usage_exit() {
  echo "Usage: $0 [-v version] [-i install_path]" 1>&2
  exit 1
}

install_path=/usr/local/bin/aqua

while getopts i:v: OPT
do
  case $OPT in
    v) version=$OPTARG
      ;;
    i) install_path=$OPTARG
      ;;
    \?) usage_exit
      ;;
  esac
done

shift $((OPTIND - 1))

uname_os() {
  local os
  os=$(uname -s | tr '[:upper:]' '[:lower:]')
  case "$os" in
    cygwin_nt*) os="windows" ;;
    mingw*) os="windows" ;;
    msys_nt*) os="windows" ;;
  esac
  echo "$os"
}

uname_arch() {
  local arch
  arch=$(uname -m)
  case $arch in
    x86_64) arch="amd64" ;;
    aarch64) arch="arm64" ;;
  esac
  echo ${arch}
}

OS="$(uname_os)"
ARCH="$(uname_arch)"

mkdir -p "$(dirname "$install_path")"

if [ -n "${version:-}" ]; then
  URL=https://github.com/aquaproj/aqua/releases/download/${version}/aqua_${OS}_${ARCH}.tar.gz
else
  URL=https://github.com/aquaproj/aqua/releases/latest/download/aqua_${OS}_${ARCH}.tar.gz
  version=latest
fi

tempdir=$(mktemp -d)
echo "===> Downloading $URL ..." >&2
curl --fail -L "$URL" -o "$tempdir/aqua_${OS}_${ARCH}.tar.gz"

(cd "$tempdir"; tar xvzf "aqua_${OS}_${ARCH}.tar.gz" > /dev/null)
echo "===> Install aqua $version ($OS/$ARCH) to $install_path" >&2

mv "$tempdir/aqua" "$install_path"

rm -R "$tempdir"

chmod a+x "$install_path"
