t = ${shell date "+20%y%m%d"}
p = ginpgx

main:
	## For Mac
	## Gin uses encoding/json as default json package
	## but you can change to jsoniter by build from other tags.
	go build -tags=jsoniter .

linux:
	# Cross compile to Linux.
	# export CGO_ENABLED=0 GOARCH=amd64 GOOS=linux
	# export GOARCH=amd64 GOOS=linux
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -tags=jsoniter .

clear:
	# rm ${p}

