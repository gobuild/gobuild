#!/bin/bash -
#
# gorelease.sh: build and publish
#
# shorted url: <http://bitly.com/gorelease>
# need login to view stats: https://bitly.com/a/stats
#
# ref:
# - [travis-ci environment](http://docs.travis-ci.com/user/environment-variables/)

set -e
set -o pipefail

echo "Is Pull Request: $TRAVIS_PULL_REQUEST"

# Set environment variables
GORELEASE_GO_VERSION=${TRAVIS_GO_VERSION:-"1.5"}
BUILD_OS=${1:-"windows linux darwin"}
TMPDIR=$PWD/.gobuild-tmp
BRANCH=

if test -z "$TRAVIS"
then
	# Here for my test
	GORELEASE_TOKEN=12345678
	BRANCH=master
else
	GORELEASE_TOKEN=${GORELEASE_TOKEN:-"$GRTOKEN"}
	GORELEASE_TOKEN=${GORELEASE_TOKEN:?}
	BRANCH=${TRAVIS_BRANCH:-$TRAVIS_TAG}
fi
KEY_PREFIX=/gorelease${PWD#*/src/github.com}/${BRANCH:?}/

echo "BRANCH: $BRANCH"
echo "PREFIX: $KEY_PREFIX"

# Set build environment
if test -n "$TRAVIS"
then
	go get github.com/mitchellh/gox
	if test $GORELEASE_GO_VERSION != "1.5"
	then
		gox -os="${BUILD_OS}" -build-toolchain
	fi
	go get github.com/gorelease/qsync
	go get github.com/gorelease/gopack
	go get github.com/tools/godep
else
	BUILD_OS="darwin"
fi


/bin/mkdir -p $TMPDIR
DISTDIR=$TMPDIR/dist

# FIXME(ssx): need support build pack
# build standalone
gopack all \
	--output "$DISTDIR/{{.OS}}-{{.Arch}}/{{.Dir}}.zip" \
	--json "$DISTDIR/builds.json"

cat > $TMPDIR/conf.ini <<EOF
[qiniu]
uphost = http://up.qiniug.com
bucket = ""
accesskey = ""
secretkey = ""
keyprefix = $KEY_PREFIX

[local]
syncdir = $DISTDIR

[gorelease]
token = "$GORELEASE_TOKEN"
host = "qntoken.herokuapp.com"
EOF

set -eu

# upload
qsync -c $TMPDIR/conf.ini
