#!/bin/bash
set -xe

# go-fuzz doesn't support modules yet, so ensure we do everything
# in the old style GOPATH way
export GO111MODULE="off"

# install go-fuzz
go get github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build

# TODO: needed until https://github.com/actions/setup-go/issues/14 is fixed
# adds GOBIN to PATH so that go-fuzz-build is visible
GOB="$(go env GOPATH)/bin"
PATH=${PATH}:"${GOB}"

# target name can only contain lower-case letters (a-z), digits (0-9) and a dash (-)
# to add another target, make sure to create it with `fuzzit create target`
# before using `fuzzit create job`
TARGET=micro-starter-kit

go-fuzz-build -libfuzzer -o ${TARGET}.a .
clang -fsanitize=fuzzer ${TARGET}.a -o ${TARGET}

# install fuzzit for talking to fuzzit.dev service
# or latest version:
# https://github.com/fuzzitdev/fuzzit/releases/latest/download/fuzzit_Linux_x86_64
wget -q -O fuzzit https://github.com/fuzzitdev/fuzzit/releases/download/v2.4.35/fuzzit_Linux_x86_64
chmod a+x ./fuzzit

# upload fuzz target for long fuzz testing on fuzzit.dev server
# or run locally for regression
if [ "${GITHUB_EVENT_NAME}" == "push" ]; then
	TYPE=fuzzing
elif [ "${GITHUB_EVENT_NAME}" == "pull_request" ]; then
	TYPE=local-regression
else
    echo "Unexpected event '${GITHUB_EVENT_NAME}'"
    exit 1
fi

./fuzzit create job --type $TYPE kkowalczyk/${TARGET} ${TARGET}
