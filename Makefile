build: build-binary

build-binary:
	docker run --rm \
		-e CGO_ENABLED=0 \
		-e GOOS="linux" \
		-e GOARCH="amd64" \
		-v `pwd`:/go/src/github.com/SeerUK/reverb \
		-w /go/src/github.com/SeerUK/reverb \
		golang:1.7 \
		bash -c 'set -x \
			&& go get -v ./... \
			&& go build -a -installsuffix cgo -v -o dist/reverb-$${GOOS}-$${GOARCH} ./cmd/reverb/ \
		'

build-image: build-binary
	docker build -t seeruk/reverb .

.PHONY: build
.SILENT:
