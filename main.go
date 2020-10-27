package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func mainPage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World .... !")
}

func handleRequests() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", mainPage)

    log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
    handleRequests()
}