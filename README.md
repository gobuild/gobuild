# gorelease
experiment, gobuild5 - for easily public released go binary.

## Step1
Save the following content into `.travis.yml`, and put it into your repository.

	language: go
	go:
	  - 1.4
	env:
	  - "PATH=/home/travis/gopath/bin:$PATH"
	before_install:
	  - go get github.com/mitchellh/gox
	  - gox -os="linux darwin windows" -build-toolchain
	script:
	  - go test -v ./...
	after_success:
	  - gox -os="linux darwin windows" -output "gorelease-temp/dist/{{.OS}}-{{.Arch}}/{{.Dir}}"
      - bash -c "$(curl -fsSL https://raw.githubusercontent.com/codeskyblue/gorelease/master/scripts/upload-qiniu.sh)"

## Step2
You need a account in [QiniuCDN](http://www.qiniu.com)

In `travis-ci.org` setting page. set three env vars (Copied from qiniu). for example

`BUCKET`也就是空间地址, 没有空间的话，选择创建一个新的空间就可以了.

	ACCESS_KEY=LKJFLSkdjfkj23lkjrl23kjflkzsjdfljwerf2w3
	SECRET_KEY=kljdlFLSDKFJo9iwejflkjLkjsdfoijw4elfkjsd
	BUCKET=gorelease

## LICENSE
[MIT](LICENSE)
