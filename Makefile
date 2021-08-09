all: build

build:
	go build -o out/ ./cmd/algodexidxsvr ./cmd/algodexidxsvr-cli

generate:
	go get goa.design/goa/v3/...
	goa gen algodexidx/design
	goa example algodexidx/design