default: build

package = {aws-tools}
binary = byron

.PHONY: default run build

run:fmt
				 GOPATH=~/${package} go run ${binary}.go ${token} 

build:
				 GOPATH=~/${package} go build -o build/${binary}

fmt:
				go fmt *.go

vet:
				go vet *.go

get:
				GOPATH=~/${package} go get ${p}
