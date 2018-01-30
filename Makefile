GOFILES=$(shell find . -type f -iname '*.go' | grep -v vendor)

bin/uuidgen_linux_armv6: ${GOFILES}
	GOOS=linux GOARCH=arm GOARM=6 go build -o "$@" .

bin/uuidgen_darwin_amd64: ${GOFILES}
	GOOS=darwin GOARCH=amd64 go build -o "$@" .

bins: bin/uuidgen_darwin_amd64 bin/uuidgen_linux_armv6
