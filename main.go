package main

// https://docs.github.com/cn/rest/reference/activity#list-repositories-starred-by-a-user

import (
	"fmt"
	"github.com/xjh22222228/github-oauth2/api"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", api.HandlerIndex)
	http.HandleFunc("/api/oauth", api.HandlerAuth)
	http.HandleFunc("/api/user", api.HandlerUser)

	err := http.ListenAndServe(":7006", nil)
	if err != nil {
		panic(err)
	}
}
