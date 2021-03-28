package api

import (
  "fmt"
  "net/http"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Github OAuth2 => https://github." +
    "com/xjh22222228/github-oauth2")
}
