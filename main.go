package main

import (
    "github.com/xjh22222228/github-oauth2/api"
    "net/http"
)

func main()  {
    http.HandleFunc("/", api.HandlerIndex)
    http.HandleFunc("/api/oauth", api.HandlerAuth)
    err := http.ListenAndServe(":7001", nil)

    if err != nil {
        panic(err)
    }
}
