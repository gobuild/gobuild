#!/bin/bash -
#
# cross compile

test X${GR_VERSION:-$TRAVIS_GO_VERSION} != X${TRAVIS_GO_VERSION} && exit 0

if test -n "$TRAVIS"
then
	go get github.com/mitchellh/gox
	gox -os="linux darwin windows" -build-toolchain
fi

OS=${1:-"windows linux darwin"}
gox -os "$OS" -output "gorelease-temp/dist/{{.OS}}-{{.Arch}}/{{.Dir}}"
