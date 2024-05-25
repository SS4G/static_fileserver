package main

import (
    "fmt"
    "flag"
    //"net/http"
)

func main() {
    var port int
    var local_path string
    flag.IntVar(&port, "port", 8888, "server_port")
    flag.StringVar(&local_path, "path", "/data/", "local path to serve")
    flag.Parse()
    fmt.Printf("Listening on port=%d path=%s\n", port, local_path)
    //http.Handle("/", http.FileServer(http.Dir(local_path)))
    //fmt.Printf("Listening on port=%d path=%s", port, local_path)
    //http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
