HOST=docker.io
PROJECT=go-rest-api
ORG=k8shop
VERSION=0.0.1
OUTPUT_DIR=./tmp/output

code/build/binary:
	go build -o ${OUTPUT_DIR}/bin/go-rest-api cmd/go-rest-api/main.go

image/build:
	docker build -t ${HOST}/${ORG}/${PROJECT}:v${VERSION} .

image/push:
	docker push ${HOST}/${ORG}/${PROJECT}:v${VERSION}

image/build/push: image/build image/push
