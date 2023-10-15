package main

import (
	"fmt"
	"net/http"
	"os"
)


var name string

func RequestHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello World from Server " + fmt.Sprint(name)))
}

func main() {
    name = os.Getenv("NAME")
    r := http.NewServeMux()
    r.Handle("/", http.HandlerFunc(RequestHandler))
    http.ListenAndServe(":8000", r)
}
