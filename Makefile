linux:
	go get github.com/docopt/docopt-go
	go build -v miniserv.go

all:
	echo ${PWD}
	go get github.com/docopt/docopt-go
	@for GOOS in darwin linux; do\
		for GOARCH in 386 amd64; do \
			go build -v -o miniserv;\
			tar -cvzf dist/miniserv-$${GOOS}-$${GOARCH}.tar.gz miniserv;\
			rm miniserv;\
		done \
	done
	@for GOARCH in 386 amd64; do \
		go build -v -o miniserv.exe; \
		zip dist/miniserv-win-$${GOARCH}.zip miniserv.exe;\
		rm miniserv.exe;\
	done

clean:
	rm miniserv-*