package main

import "net/http"


func RequestHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello World"))
}

func main() {
    r := http.NewServeMux()
    r.Handle("/", http.HandlerFunc(RequestHandler))
    http.ListenAndServe(":8000", r)
}
