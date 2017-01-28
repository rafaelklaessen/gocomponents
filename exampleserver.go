package main

import (
        "net/http"
        "./gocomponents"
)

func index(w http.ResponseWriter, r *http.Request) {
        gocomponents.Header(w, r)
}

func main() {
        http.HandleFunc("/", index)
        http.ListenAndServe(":666", nil)
}