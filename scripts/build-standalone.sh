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

#language: go
#go:
#  - 1.4
#env:
#  - "PATH=/home/travis/gopath/bin:$PATH"
#before_install:
#  - go get github.com/mitchellh/gox
#  - gox -build-toolchain
#  - go get github.com/tcnksm/ghr
#script:
#  - go test -v ./...
#after_success:
#  - gox -output "dist/{{.OS}}-{{.Arch}}/{{.Dir}}"
#  - ghr --username codeskyblue --token $GITHUB_TOKEN --replace --prerelease --debug pre-release dist/
