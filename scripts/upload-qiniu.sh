#!/bin/bash -
#
# ref travis-ci environment: http://docs.travis-ci.com/user/environment-variables/
#

test X${GR_VERSION:-$TRAVIS_GO_VERSION} != X${TRAVIS_GO_VERSION} && exit 0
echo "Go version: $TRAVIS_GO_VERSION"

GOOS=$(go env GOOS)
GOARCH=$(go env GOARCH)

TMPDIR=$PWD/gorelease-temp
/bin/mkdir -p $TMPDIR

BRANCH=
if test -z "$TRAVIS"
then
	# Here for my test
	BRANCH=$(git symbolic-ref --short HEAD)
	ACCESS_KEY=V6cm-H-uL5Lh0hrPbF28Y1KJ99dW8d2p9lUQRDMJ
	SECRET_KEY=gFatds2RE8MWZSqbVOwsztp8EAqtHUOnWC6NGKVU
	BUCKET=gorelease
else
	BRANCH=${TRAVIS_BRANCH:-$TRAVIS_TAG}
	ACCESS_KEY=${ACCESS_KEY:?}
	SECRET_KEY=${SECRET_KEY:?}
	BUCKET=${BUCKET:?}

	go get github.com/gorelease/qsync
fi
echo "Branch: $BRANCH"


KEY_PREFIX=gorelease/$(basename $PWD)/${BRANCH:?}/
#KEY_PREFIX=${PWD##$GOPATH/src/}/$branch/

#wget -q http://devtools.qiniu.com/qiniu-devtools-${GOOS}_${GOARCH}-current.tar.gz -O- | tar -xz -C $TMPDIR
#/bin/rm -fr $HOME/.qrsync

set -eu
DISTDIR=$TMPDIR/dist

cat > $TMPDIR/conf.ini <<EOF
[qiniu]
uphost = http://up.qiniug.com
bucket = $BUCKET
accesskey = "$ACCESS_KEY"
secretkey = "$SECRET_KEY"
keyprefix = $KEY_PREFIX

[local]
syncdir = $DISTDIR
EOF

cat > $DISTDIR/builds.json <<EOF
{
	"update_time": $(date +%s),
	"go_version": "$TRAVIS_GO_VERSION"
}
EOF
# upload
qsync -c $TMPDIR/conf.ini

