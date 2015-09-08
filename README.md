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
