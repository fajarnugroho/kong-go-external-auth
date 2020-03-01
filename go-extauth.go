package main

import (
	"net/http"
	"strings"

	"github.com/Kong/go-pdk"
)

type Config struct {
	Url     string
	Method  string
	Headers string
}

func New() interface{} {
	return &Config{}
}
func CreateResponseMessage(msg string) string {
	return "{\"message\": \"" + msg + "\"}"
}
func (conf *Config) Access(kong *pdk.PDK) {
	client := &http.Client{}
	responseHeader := map[string][]string{
		"Content-Type": []string{"application/json"},
	}
	req, err := http.NewRequest(conf.Method, conf.Url, nil)
	if err != nil {
		kong.Log.Err(err)
		kong.Response.Exit(500, CreateResponseMessage("Internal Server Error"), responseHeader)
	}

	headers, err := kong.Request.GetHeaders(-1)
	if err != nil {
		kong.Log.Err(err)
		kong.Response.Exit(400, CreateResponseMessage("No Header Specified"), responseHeader)
	}
	configHeader := strings.Split(conf.Headers, ",")
	for _, header := range configHeader {
		if val, ok := headers[header]; ok {
			req.Header.Set(header, strings.Join(val, " "))
		} else {
			kong.Response.Exit(400, CreateResponseMessage("Missing '"+header+"'header"), responseHeader)
		}
	}
	res, err := client.Do(req)
	if err != nil {
		kong.Log.Err(err)
		kong.Response.Exit(500, CreateResponseMessage("Internal Server Error"), responseHeader)
	}
	if res.StatusCode != 200 {
		kong.Response.Exit(200, CreateResponseMessage("Your are not authorized"), responseHeader)
	}
}
