#!/usr/bin/env bash
set -e

cd $(dirname "$0")
CWD=$(pwd)

mkdir -p bin
mkdir -p mocks

# Use this to ensure that we have all the tools required to do a build.
export CGO_ENABLED=0
export GO111MODULE=on
export GOFLAGS="-mod=vendor"

MISSING=()

check() {
  local X=$1
  set +e
  command -v $X >/dev/null 2>&1
  local RESULT=$?
  set -e
  if [ $RESULT != 0 ]; then
    MISSING+=($X)
  fi
}

check git
check go
check helm
check jq
check kubectl
# check zip

if ! [ ${#MISSING[@]} -eq 0 ]; then
  echo "Missing prerequisites:"
  for X in $MISSING; do
    echo "  $X"
  done

  exit 1
fi

echo "Prerequisites present"

GIT_COMMIT=$(git rev-list -1 HEAD)
VERSION=$(git describe --tags)
if ! [ -z "$(git status --porcelain)" ]; then
  # There are untracked or unstaged changes
  GIT_COMMIT="DIRTY-${GIT_COMMIT}"
  VERSION="WIP-${VERSION}"
fi

echo "Build version '$VERSION', git SHA '$GIT_COMMIT'"
LDFLAGS_IMPORTS="-X github.com/object88/shipbot.GitCommit=${GIT_COMMIT} -X github.com/object88/shipbot.AppVersion=${VERSION}"

cd "$CWD"

# default to mostly true, set env val to override
DO_PACKAGE=${DO_PACKAGE:-"false"}
DO_TEST=${DO_TEST:-"true"}
DO_VERIFY=${DO_VERIFY:-"true"}
DO_VET=${DO_VET:-"true"}

while [[ $# -gt 0 ]]; do
  key="$1"
  case $key in
    --fast)
        DO_TEST="false"
        DO_VERIFY="false"
        DO_VET="false"
        shift
        ;;
    --no-test)
        DO_TEST="false"
        shift
        ;;
    --no-verify)
        DO_VERIFY="false"
        shift
        ;;
    --no-vet)
        DO_VET="false"
        shift
        ;;
    --package)
        DO_PACKAGE="true"
        shift
        ;;
    *)
      shift
      ;;
  esac
done

if [[ $DO_TEST == "true" ]]; then
  if ! [ -x $CWD/bin/mockgen ]; then
    echo "Building mockgen"
    time go build -o $CWD/bin/mockgen $CWD/vendor/github.com/golang/mock/mockgen
    echo ""
  fi
  echo "mockgen tool checked"

  echo "Generating mock for RoundTripper"
  time $CWD/bin/mockgen -destination=$CWD/mocks/mock_httproundtripper.go -package=mocks net/http RoundTripper
  echo ""
fi

if [[ $DO_VERIFY == "true" ]]; then
  echo "Verifying modules"
  # returns non-zero if this doesn't verify out
  time go mod verify
  echo ""
fi

if [[ $DO_VET == "true" ]]; then
  # Vet's exit code is non-zero for erroneous invocation of the tool
  # or if a problem was reported, and 0 otherwise. Note that the
  # tool does not check every possible problem and depends on
  # unreliable heuristics, so it should be used as guidance only,
  # not as a firm indicator of program correctness.
  # [snip]
  # By default, all checks are performed.
  #
  # https://golang.org/cmd/vet/
  echo "Running vet"
  time go vet $(go list ./...)
  echo ""
fi

# build executable(s)
# method found here https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04

echo "Building..."

DEFAULT_GOOS=$(uname | tr '[:upper:]' '[:lower:]')
PLATFORMS=( "$DEFAULT_GOOS/amd64" )
if [ "$BUILD_AND_RELEASE" == "true" ]; then
  PLATFORMS=( "linux/amd64" "darwin/amd64" )
fi

# build executable for each platform...
for PLATFORM in "${PLATFORMS[@]}"; do
  export GOOS=$(cut -d'/' -f1 <<< $PLATFORM)
  export GOARCH=$(cut -d'/' -f2 <<< $PLATFORM)
  BINARY_NAME="shipbot-${GOOS}-${GOARCH}"
  if [ $DEFAULT_GOOS == $GOOS ]; then
    export TEST_BINARY_NAME="$CWD/bin/$BINARY_NAME"
  fi
  echo "Building as $BINARY_NAME"

  if [ $(uname) == "Darwin" ]; then
    # Cannot do a static compilation on Darwin.
    time go build -o ./bin/$BINARY_NAME -ldflags "-s -w $LDFLAGS_IMPORTS" ./main/main.go
  else
    time go build -o ./bin/$BINARY_NAME -tags "netgo" -ldflags "-extldflags \"-static\" -s -w $LDFLAGS_IMPORTS" ./main/main.go
  fi

  if [ $DO_PACKAGE == "true" ]; then
    zip -j ./bin/$BINARY_NAME.zip ./bin/$BINARY_NAME
  fi
  echo ""
done

# test executables and binaries
if [[ $DO_TEST == "true" ]]; then
  export TEST_SHA=${GIT_COMMIT}
  export TEST_VERSION=${VERSION}

  echo "Testing with $TEST_BINARY_NAME"
  time go test ./... -count=1 -tags test_integration
  echo ""
fi

echo "Done"
