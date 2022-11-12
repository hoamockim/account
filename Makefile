GOCMD=go
GOBUILD=$(GOCMD) build
GOMOD=$(GOCMD) mod
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

update:
	go mod tidy
all:
	make build && make run internal
build:
	$(GOBUILD) -v -ldflags="-extldflags=-static" -o $(SERVICE_NAME) app/cmd/main.go
test:
	$(GOTEST) -v ./...  -covermode=count -coverprofile=sample.cov
run $(app):
	./$(SERVICE_NAME) $(app)
