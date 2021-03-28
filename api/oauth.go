// Copyright 2021 the xiejiahe. All rights reserved. MIT license.
// https://github.com/settings/developers

package api

import (
  "embed"
  "encoding/json"
  "fmt"
  "io"
  "io/fs"
  "net/http"
  "net/url"
  "strings"
)

type Oauth struct {
  ClientId string `json:"client_id"`
  ClientSecret string `json:"client_secret"`
  Code string `json:"code"`
}

type Response struct {
  Message string `json:"message"`
  Data interface{} `json:"data"`
  Status int `json:"status"`
}

//go:embed config.json
var f embed.FS

func HandlerAuth(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Request-Method", "GET, POST, OPTIONS")
  w.Header().Set("Access-Control-Allow-Headers", "Content-Type, " +
    "X-Requested-With, Authorization, Origin, Accept")

  var oauth Oauth
  jsonContent, _ := fs.ReadFile(f, "config.json")
  errJson := json.Unmarshal(jsonContent, &oauth)

  if errJson != nil {
    panic(errJson)
  }

  code := r.FormValue("code")

  fmt.Println("code ==>", code)

  if code == "" {
    j, _ := json.Marshal(Response{
      Message: "code cannot be empty",
      Data: nil,
      Status: 401,
    })

    fmt.Fprintf(w, string(j))
    return
  }

  payload := Oauth{
    ClientId: oauth.ClientId,
    ClientSecret: oauth.ClientSecret,
    Code: code,
  }

  b, _ := json.Marshal(payload)

  resp, err := http.Post(
    "https://github.com.cnpmjs.org/login/oauth/access_token",
    "application/json",
    strings.NewReader(string(b)))

  defer resp.Body.Close()

  if err != nil {
    w.WriteHeader(http.StatusUnauthorized)
    j, _ := json.Marshal(Response{
      Message: "HTTP ERROR",
      Data: nil,
      Status: http.StatusUnauthorized,
    })

    fmt.Fprintf(w, string(j))
    return
  }

  contents, _ := io.ReadAll(resp.Body)
  response := string(contents)
  values, err := url.ParseQuery(response)
  accessToken := values.Get("access_token")

  if err != nil || accessToken == "" {
    w.WriteHeader(http.StatusUnauthorized)
    j, _ := json.Marshal(Response{
      Message: response,
      Data: nil,
      Status: http.StatusUnauthorized,
    })

    fmt.Fprintf(w, string(j))
    return
  }

  fmt.Println("response ==> ", response)

  j, _ := json.Marshal(Response{
    Message: "OK",
    Data: map[string]string{
      "accessToken": accessToken,
    },
    Status: http.StatusOK,
  })

  fmt.Fprintf(w, string(j))
}