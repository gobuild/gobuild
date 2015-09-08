all:
	(cd public; go-bindata -pkg public .)
	(cd templates; go-bindata -pkg templates .)
	go build
