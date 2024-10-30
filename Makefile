#@IgnoreInspection BashAddShebang
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export DEBUG=true
export APP=scimta-be
export LDFLAGS="-w -s"

all: init

init:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/air-verse/air@latest
	swag i
	go install

build:
	go build -race -v -o dist/scimta-be -ldflags $(LDFLAGS) .

build-win:
	go build -v -o dist/scimta-be -ldflags $(LDFLAGS) .

build-static:
	CGO_ENABLED=0 go build -race -v -o $(APP) -a -installsuffix cgo -ldflags $(LDFLAGS) .

run:
	go run -race .

air:
	air

air-win:
	air -build.cmd "go build -o .\tmp\main.exe ." -build.bin ".\tmp\main.exe" 
	
# dev:
# 	bash -c "trap 'docker-compose down' EXIT; docker-compose -f ./docker-compose.dev.yml up -d --build && air"

swagger:
	swag i

############################################################
# Test
############################################################

test:
	go test -v -race ./...

.PHONY: all build build-static run test air