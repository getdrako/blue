#!/usr/bin/env bash

PACKAGE="drako/server/pkg"

usage() {
    echo "Use: $0 --config=Debug|Release"
    exit 1
}

if [[ $# -ne 1 ]]; then
    usage
fi

CONFIG=""
for i in "$@"; do
    case $i in
        --config=*)
            CONFIG="${i#*=}"
            shift
            ;;
        *)
            usage
            ;;
    esac
done

if [[ "$CONFIG" != "Debug" && "$CONFIG" != "Release" ]]; then
    usage
fi

if [[ "$CONFIG" == "Debug" ]]; then
    VERSION_TYPE="0x01"
    DRAKO_DEBUG_BOOL="true"
else
    VERSION_TYPE="0x09"
    DRAKO_DEBUG_BOOL="false"
fi

VERSION_TAG=$(git describe --tags --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null)

if [[ -n "$VERSION_TAG" ]]; then
    VERSION="${VERSION_TAG#v}"
else
    VERSION="0.0.0"
fi

IFS='.' read -r MAJOR MINOR PATCH <<< "$VERSION"

COMMIT_HASH=$(git rev-parse --short HEAD)
BUILD_TIMESTAMP=$(date '+%Y-%m-%dT%H:%M:%S')

LDFLAGS=(
    "-X '${PACKAGE}/build.DRAKO_VERSION_MAJOR=${MAJOR}'"
    "-X '${PACKAGE}/build.DRAKO_VERSION_MINOR=${MINOR}'"
    "-X '${PACKAGE}/build.DRAKO_VERSION_PATCH=${PATCH}'"
    "-X '${PACKAGE}/build.VersionType=${VERSION_TYPE}'"
    "-X '${PACKAGE}/build.DRAKO_DEBUG_BOOL=${DRAKO_DEBUG}'"
    "-X '${PACKAGE}/build.Commit=${COMMIT_HASH}'"
    "-X '${PACKAGE}/build.Datetime=${BUILD_TIMESTAMP}'"
)

go build -x -ldflags="${LDFLAGS[*]}" -o ./bin/drako ./cmd/
