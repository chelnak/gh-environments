#!/bin/bash
set -e

tag=$(git describe --tags --abbrev=0)
platforms=$(echo "darwin-amd64,linux-386,linux-arm,linux-amd64,linux-arm64,windows-386,windows-amd64" | tr "," "\n")
include="dist/*"

if [ -n "${GH_VERSION_SPEC}" ]; then
  # build the ldflags arg in two parts due to this limitation:
  # https://unix.stackexchange.com/questions/502715/escape-parameters-for-positional-arguments
  LDFLAG=-ldflags
  VERSION_SPEC="-X ${GH_VERSION_SPEC}=${tag}"
fi

for p in $platforms; do
  goos=$(echo $p | sed 's/-.*//')
  goarch=$(echo $p | sed 's/.*-//')
  ext=""
  if [[ "${goos}" == "windows" ]]; then
    ext=".exe"
  fi
  GOOS=${goos} GOARCH=${goarch} go build -o "dist/${goos}-${goarch}${ext}" ${LDFLAG} "${VERSION_SPEC}"
done

ls -A dist >/dev/null || (echo "no files found in dist/" && exit 1)

gh api repos/$GITHUB_REPOSITORY/releases/generate-notes \
  -f tag_name="${tag}" -q .body > CHANGELOG.md

gh release create $tag --notes-file CHANGELOG.md $include