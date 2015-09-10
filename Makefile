all:
	(cd public; go-bindata -ignore "\.go" -pkg public .)
	(cd templates; go-bindata -ignore "\.go" -pkg templates .)
	go build
