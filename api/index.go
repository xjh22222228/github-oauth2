package api

import (
	"fmt"
	"net/http"
	"os"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	clientId := os.Getenv("client_id")
	fmt.Fprintf(w, `
Github OAuth2 => <a href="https://github.com/xjh22222228/github-oauth2" target="_blank">https://github.com/xjh22222228/github-oauth2</a>
`+"\nclient_id="+clientId)
}
