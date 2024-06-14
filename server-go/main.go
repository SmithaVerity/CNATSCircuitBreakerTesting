package main

import (
 "fmt"
 "net/http"
)

func main() {

 http.HandleFunc("/api/v1/ping", ping)
 fmt.Println("Server Application running.")
 http.ListenAndServe(":8081", nil)
}

func ping(w http.ResponseWriter, r *http.Request) {

 fmt.Fprintf(w, "Hiiii from Server!!!!")
}
