// Copyright 2021 the xiejiahe. All rights reserved. MIT license.
// https://github.com/settings/developers

package api

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/xjh22222228/github-oauth2/utils"
	"net/http"
	"net/url"
	"os"
)

type Oauth struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
}

type Map map[string]interface{}

func HandlerAuth(w http.ResponseWriter, r *http.Request) {
	utils.Cors(w)

	var oauth Oauth

	code := r.FormValue("code")
	clientId := os.Getenv("client_id")
	clientSecret := os.Getenv("client_secret")

	fmt.Println("code ==>", code)

	if code == "" {
		utils.Body(w, utils.Stringify(Response{
			Message: "code cannot be empty",
			Data:    nil,
			Status:  401,
		}))
		return
	}

	payload := &Oauth{
		ClientId:     oauth.ClientId,
		ClientSecret: oauth.ClientSecret,
		Code:         code,
	}

	if clientId != "" {
		payload.ClientId = clientId
	}

	if clientSecret != "" {
		payload.ClientSecret = clientSecret
	}
	response := utils.Fetch(&utils.FetchConf{
		Method: http.MethodPost,
		Url:    "https://github.com/login/oauth/access_token",
		Data:   payload,
		R:      r,
	})

	values, err := url.ParseQuery(response)
	accessToken := values.Get("access_token")

	if err != nil || accessToken == "" {
		w.WriteHeader(http.StatusUnauthorized)
		utils.Body(w, utils.Stringify(Response{
			Message: "Failed",
			Data:    nil,
			Status:  http.StatusUnauthorized,
		}))
		return
	}

	userRes := utils.Fetch(&utils.FetchConf{
		Method: "GET",
		Url:    "https://api.github.com/user",
		Headers: map[string]string{
			"Authorization": "token " + accessToken,
		},
		R: r,
	})
	var user User
	json.Unmarshal([]byte(userRes), &user)

	fmt.Println("user", user)

	utils.Body(w, utils.Stringify(Response{
		Message: "OK",
		Data: Map{
			"accessToken": accessToken,
			"user":        user,
		},
		Status: http.StatusOK,
	}))
}
