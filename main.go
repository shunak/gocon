package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // クライアントからリクエストを受けた場合
        log.Println("received request")
        // HTTPリクエストを受けた場合
        fmt.Fprintf(w, "Hello Docker!!")
    })
    http.HandleFunc("/page2", func(w http.ResponseWriter, r *http.Request) {
        // クライアントからリクエストを受けた場合
        log.Println("received request2")
        // HTTPリクエストを受けた場合
        fmt.Fprintf(w, "You Win!!")
    })

    log.Println("start server")
    server := &http.Server{Addr: ":8080"}
    if err := server.ListenAndServe(); err != nil {
        log.Println(err)
    }
}
