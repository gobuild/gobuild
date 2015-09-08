#!/bin/bash -
#

GOOS=$(go env GOOS)
GOARCH=$(go env GOARCH)

TMPDIR=$PWD/gorelease-temp
mkdir $TMPDIR

ACCESS_KEY=V6cm-H-uL5Lh0hrPbF28Y1KJ99dW8d2p9lUQRDMJ
SECRET_KEY=gFatds2RE8MWZSqbVOwsztp8EAqtHUOnWC6NGKVU
BUCKET=gorelease

ACCESS_KEY=${ACCESS_KEY:?}
SECRET_KEY=${SECRET_KEY:?}
BUCKET=${BUCKET:?}

KEY_PREFIX=${PWD##$GOPATH/src/}/

wget -q http://devtools.qiniu.com/qiniu-devtools-${GOOS}_${GOARCH}-current.tar.gz -O- | tar -xz -C $TMPDIR

DISTDIR=$TMPDIR/dist

cat > $TMPDIR/conf.json <<EOF
{
	"src": "$DISTDIR",
	"dest": "qiniu:access_key=$ACCESS_KEY&secret_key=$SECRET_KEY&bucket=$BUCKET&key_prefix=$KEY_PREFIX",
	"debug_level": 1
}
EOF

# upload
/bin/rm -fr $HOME/.qrsync
$TMPDIR/qrsync $TMPDIR/conf.json
