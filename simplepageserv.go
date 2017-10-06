package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	content := `<html>
    <header>
        <title>Hej</title>
    </header>
    <body>
Hello world
    </body>
</html>
`
	io.WriteString(w, content)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
