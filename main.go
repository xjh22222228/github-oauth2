// Copyright 2021 the xiejiahe. All rights reserved. MIT license.

package main

import (
    "embed"
    "encoding/json"
    "fmt"
    "github.com/gin-gonic/gin"
    cors "github.com/rs/cors/wrapper/gin"
    "io"
    "io/fs"
    "net/http"
    "net/url"
    "strings"
)

// https://github.com/settings/developers

type Oauth struct {
    ClientId string `json:"client_id"`
    ClientSecret string `json:"client_secret"`
    Code string `json:"code"`
}

//go:embed config.json
var f embed.FS

func main()  {
    var oauth Oauth
    jsonContent, err := fs.ReadFile(f, "config.json")
    errJson := json.Unmarshal(jsonContent, &oauth)

    if errJson != nil {
        panic(errJson)
    }

    r := gin.Default()
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders: []string{
            "Content-Type", "X-Requested-With",
            "Authorization", "Origin",
            "Accept",
        },
    })

    r.Use(c)

    r.GET("/", func(context *gin.Context) {
        context.String(200, "Github OAuth2")
    })

    r.GET("/api/oauth/token", func(c *gin.Context) {
        code := c.Query("code")

        fmt.Println("code==>", code)

        if code == "" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "message": "code cannot be empty",
                "data": nil,
            })
            return
        }

        payload := Oauth{
            ClientId: oauth.ClientId,
            ClientSecret: oauth.ClientSecret,
            Code: code,
        }

        b, _ := json.Marshal(payload)

        r, err := http.Post(
            "https://github.com/login/oauth/access_token",
            "application/json",
            strings.NewReader(string(b)))

        defer r.Body.Close()

        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{
                "message": "HTTP ERROR",
                "data": nil,
            })
            return
        }

        contents, _ := io.ReadAll(r.Body)
        response := string(contents)
        values, err := url.ParseQuery(response)
        accessToken := values.Get("access_token")

        if err != nil || accessToken == "" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "message": response,
                "data": nil,
            })
            return
        }

        fmt.Println("response ==> ", response)

        c.JSON(http.StatusOK, gin.H{
            "message": "OK",
            "data": gin.H{
                "accessToken": accessToken,
            },
        })
    })


    errors := r.Run(":7001")
    if errors != nil {
        panic(err)
    }
}
