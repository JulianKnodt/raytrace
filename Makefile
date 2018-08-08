
fmt:
	go fmt ./... && go vet ./...

bin: fmt 
	go build

test: fmt
	go test ./...
