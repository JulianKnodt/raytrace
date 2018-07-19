
fmt:
	go fmt ./... && go vet

binary: fmt 
	go build
