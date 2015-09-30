#!/bin/bash -x
#

if ! test -f .bowerrf
then
cat > .bowerrc <<EOF
{
	"directory": "public/components"
}
EOF
fi

bower install
gorelease -debug
