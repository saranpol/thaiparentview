package main

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", index)
}


func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "<h1>Thai Parent View</h1>")
}
