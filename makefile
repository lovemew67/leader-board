## Define variables

# static variables
APP=leader-board
SOURCE=./...
CONF=local.toml
PKGPATH=github.com/lovemew67/public-misc/cornerstone

# derived variables

# variables from shell
GOPATH=$(shell env | grep GOPATH | cut -d'=' -f 2)
REVISION=$(shell git rev-list -1 HEAD)
BR=$(shell git rev-parse --abbrev-ref HEAD)
TAG=$(shell git tag --sort=-version:refname --points-at HEAD | head -n1)
DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
THISDIR=$(dir $(abspath $(firstword $(MAKEFILE_LIST))))

# if-else blocks
ifeq ($(TAG),)
TAG=$(REVISION)
endif

# exported variables


## Define targets

genproto:
	protoc -I ./proto \
	--go_out ./gen/go/proto --go_opt paths=source_relative \
	--openapiv2_out ./gen/openapiv2 --openapiv2_opt logtostderr=true \
	./proto/*.proto

injecttag:
	protoc-go-inject-tag -input=./gen/go/proto/staff.pb.go

genall: genproto injecttag

mock:
	./mock-script

swaggerui:
	docker run --rm --name swaggerui -p 8080:8080 \
	-v $(THISDIR)gen/openapiv2/staff_service.swagger.json:/app/swagger.json \
	swaggerapi/swagger-ui:v3.37.2

updatevendor:
	go mod tidy
	go mod vendor
	go vet ./...

test: updatevendor	
	go test -mod=vendor -v -race -cover -timeout 180s ./...

install: updatevendor
	go install -mod=vendor -v -ldflags "-s -X $(PKGPATH).appName=$(APP) -X $(PKGPATH).gitCommit=$(REVISION) -X $(PKGPATH).gitBranch=$(BR) -X $(PKGPATH).appVersion=$(TAG) -X $(PKGPATH).buildDate=$(DATE)" $(SOURCE) 

leaderboard: install
	$(GOPATH)/bin/$(APP) leaderboard --config $(CONF)
