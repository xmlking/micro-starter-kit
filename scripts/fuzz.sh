#!/bin/bash
set -xe

# install go-fuzz
go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build

# This is current workaround to support go modules
find $GOPATH
cd $GOPATH/src/github.com/dvyukov/go-fuzz
git remote add fork https://github.com/fuzzitdev/go-fuzz
git fetch
git checkout fork
go install ./...

# TODO: needed until https://github.com/actions/setup-go/issues/14 is fixed
# adds GOBIN to PATH so that go-fuzz-build is visible
GOB="$(go env GOPATH)/bin"
PATH=${PATH}:"${GOB}"

# target name can only contain lower-case letters (a-z), digits (0-9) and a dash (-)
# to add another target, make sure to create it with `fuzzit create target`
# before using `fuzzit create job`
TARGET=crypto
cd ./shared/crypto
go-fuzz-build -libfuzzer -o ${TARGET}.a .
clang -fsanitize=fuzzer ${TARGET}.a -o ${TARGET}

# install fuzzit for talking to fuzzit.dev service
# or latest version:
# https://github.com/fuzzitdev/fuzzit/releases/latest/download/fuzzit_Linux_x86_64
wget -q -O fuzzit https://github.com/fuzzitdev/fuzzit/releases/download/v2.4.54/fuzzit_Linux_x86_64
chmod a+x ./fuzzit

./fuzzit create job --type $1 m-starter-kit/${TARGET} ${TARGET}
