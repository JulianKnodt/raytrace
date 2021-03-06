bin: fmt
	go build

fmt:
	@go fmt ./...
	@go vet ./...

test: fmt
	go test ./...

test-race: fmt
	go test ./... -race

bench: fmt
	go test ./... -bench=. -benchmem

diff-off: bin
	./raytrace -off=off/testdata/dragon.off -shift="0 0 -0.5" -out=testdata/differ.png
	diff testdata/differ.png testdata/og_off.png
	rm testdata/differ.png

diff-obj: bin
	./raytrace -obj=obj/testdata/teapot.obj -shift="0 -2 -3" -out=testdata/differ_obj.png
	diff testdata/differ_obj.png testdata/og_obj.png
	rm testdata/differ_obj.png

todos:
	grep -rn . -e "TODO"

sources:
	grep -rn . --binary-files=without-match -e "http" | sort | uniq

only-sources:
	grep -hrn . --binary-files=without-match -e "http" | sort | uniq
