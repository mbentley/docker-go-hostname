package main

import (
    "fmt"
    "log"
    "net/http"
    "net"
    "os"
)

func main() {
    log.SetFlags(log.LstdFlags | log.Lshortfile) // Enable short file names and standard flags

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        ip, _, err := net.SplitHostPort(r.RemoteAddr)
        if err != nil {
            log.Println("Received request:", r.URL.Path)
        } else {
            log.Println("Received request from", ip, ":", r.URL.Path)
        }
        fmt.Fprint(w, "The hostname is: "+os.Getenv("HOSTNAME"))
    })

    http.ListenAndServe(":8080", nil)
}
