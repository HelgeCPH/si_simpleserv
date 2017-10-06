package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"runtime"

	"github.com/docopt/docopt-go"
)

// App creates a mux and binds the root route for processing
// static files.
func startServing(pathToServe string) http.Handler {

	// Create a new mux for this service.
	m := http.NewServeMux()

	// Bind the route for serving static files using the
	// default FileServer. This will load the home page.
	if pathToServe == "" {
		m.Handle("/", http.FileServer(http.Dir(staticDir())))
	} else {
		m.Handle("/", http.FileServer(http.Dir(pathToServe)))
	}

	return m
}

// staticDir builds a full path to the 'static' directory
// that is relative to this file.
func staticDir() string {

	// Locate from the runtime the location of
	// the apps static files.
	_, filename, _, _ := runtime.Caller(1)

	// Return a path to the static folder.
	return path.Join(path.Dir(filename), "static")
}

func main() {
	usage := `Mini HTTP File Server.

Usage:
  miniserv [<path>] [--port=<port>]
  miniserv -h | --help
  miniserv --version

Options:
  -h --help     Show this screen.
  --version     Show version.
  --port=<port>  Port for serving [default: 8080].
`
	arguments, _ := docopt.Parse(usage, nil, true, "Mini HTTP File Server 1.0", false)
	path := "./"
	if arguments["<path>"] != nil {
		path = arguments["<path>"].(string)
	}
	port := arguments["--port"].(string)

	log.Print(fmt.Sprintf("Serving %s on localhost:%s", path, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", port), startServing(path)))
}
