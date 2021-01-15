package main

import (
	"fmt"
	"net/http"
)
func handler() {
List := func list(rw http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc("/",list)
	}
done := func done(rw http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc("done",done)
add := func add(rw http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc("/add", add)
	}

http.ListenAndServe(":8080", nil)
}
