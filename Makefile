
fmt:
	go fmt ./... && go vet ./...

bin: fmt 
	go build

test: fmt
	go test ./...

diff: bin
	./raytrace -off=off/testdata/dragon.off -shift="0 0 -2" -out=testdata/differ.png
	diff testdata/differ.png testdata/og.png
	rm testdata/differ.png
