
fmt:
	go fmt ./...
	go vet ./...

bin: fmt 
	go build

test: fmt
	go test ./...

bench: fmt
	go test ./... -bench=. -benchmem

diff-off: bin
	./raytrace -off=off/testdata/dragon.off -shift="0 0 -2" -out=testdata/differ.png
	diff testdata/differ.png testdata/og.png
	rm testdata/differ.png

diff-obj: bin
	./raytrace -obj=obj/testdata/teapot.obj -shift="0 -2 -10" -out=testdata/differ_obj.png
	diff testdata/differ_obj.png testdata/og_obj.png
	rm testdata/differ_obj.png

todos:
	grep -rn . -e "TODO"

sources:
	grep -rn . --binary-files=without-match -e "http" | sort | uniq
