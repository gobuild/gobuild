#!/bin/bash -
#

set -x
echo $TRAVIS_GO_VERSION
test X${GR_VERSION:-$TRAVIS_GO_VERSION} != X${TRAVIS_GO_VERSION} && exit 0

if test -n "$TRAVIS"
then
	go get github.com/mitchellh/gox
	gox -os="linux darwin windows" -build-toolchain
fi

set -eu

TARGETDIR=gorelease-temp/dist
mkdir -p $TARGETDIR

go get -v

NAME=$(basename $PWD)
if test $(go env GOOS) = "windows"
then
	NAME=${NAME}.exe
fi

copy(){
	if ! test -e "$1"
	then
		echo "$1 not exists, ignore"
		return
	fi
	cp -r "$1" "$2"
}
build(){
	local BUILDDIR=$TARGETDIR/tmp
	local OUTPUT=$TARGETDIR/$GOOS-$GOARCH
	mkdir -p $OUTPUT

	go build -o $BUILDDIR/$NAME
	for resource in templates public scripts conf etc README.md LICENSE
	do
		copy $resource $BUILDDIR
	done
	(cd $BUILDDIR; zip -r ../$GOOS-$GOARCH/$NAME.zip *)
	rm -fr ${BUILDDIR:?}
}

GOOS=linux GOARCH=amd64 build
GOOS=windows GOARCH=amd64 build
GOOS=darwin GOARCH=amd64 build
