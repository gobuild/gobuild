#!/bin/bash -
#

GOOS=$(go env GOOS)
GOARCH=$(go env GOARCH)

TMPDIR=$PWD/gorelease-temp
/bin/mkdir -p $TMPDIR

#ACCESS_KEY=V6cm-H-uL5Lh0hrPbF28Y1KJ99dW8d2p9lUQRDMJ
#SECRET_KEY=gFatds2RE8MWZSqbVOwsztp8EAqtHUOnWC6NGKVU
#BUCKET=gorelease

ACCESS_KEY=${ACCESS_KEY:?}
SECRET_KEY=${SECRET_KEY:?}
BUCKET=${BUCKET:?}

branch=$(git symbolic-ref --short HEAD)
KEY_PREFIX=$(basename $PWD)/$branch/
#KEY_PREFIX=${PWD##$GOPATH/src/}/$branch/

#wget -q http://devtools.qiniu.com/qiniu-devtools-${GOOS}_${GOARCH}-current.tar.gz -O- | tar -xz -C $TMPDIR
#/bin/rm -fr $HOME/.qrsync

go get -v github.com/codeskyblue/qsync

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

# upload
qsync -c $TMPDIR/conf.ini
