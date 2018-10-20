linux:
	go get github.com/docopt/docopt-go
	go build -v miniserv.go

all:
	go get github.com/docopt/docopt-go
	@for GOOS in darwin linux; do\
		for GOARCH in 386 amd64; do \
			go build -v -o miniserv-$${GOOS}-$${GOARCH};\
		done \
	done
	@for GOARCH in 386 amd64; do \
		go build -v -o miniserv-windows-$${GOARCH}.exe; \
	done

clean:
	rm miniserv-*