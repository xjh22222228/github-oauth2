package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Cors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Request-Method", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, "+
		"X-Requested-With, Authorization, Origin, Accept")
}

func Stringify(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}
	return string(b)
}

type FetchConf struct {
	Method  string
	Url     string
	Data    interface{}
	Headers map[string]string
	R       *http.Request
}

func Fetch(c *FetchConf) string {
	body := ""
	if c.Method == http.MethodPost {
		body = Stringify(c.Data)
	}
	url := c.Url
	request, err := http.NewRequest(c.Method, url, strings.NewReader(body))
	if err != nil {
		return ""
	}

	request.Header.Add("Content-Type", "application/json")
	accessToken := c.R.Header.Get("Authorization")
	if accessToken != "" {
		request.Header.Add("Authorization", "token "+accessToken)
	}
	for k, v := range c.Headers {
		request.Header.Add(k, v)
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("%v %v, 请求错误: %v\n\n", c.Method, c.Url, err.Error())
		return ""
	}

	defer resp.Body.Close()
	contents, _ := io.ReadAll(resp.Body)
	fmt.Printf("%v %v, Message: %v\n\n", c.Method, c.Url, string(contents))
	return string(contents)
}

func Body(w io.Writer, s string) {
	fmt.Fprintf(w, s)
}
