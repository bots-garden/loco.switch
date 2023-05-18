// Package main
package main

import (
	capsule "github.com/bots-garden/capsule-module-sdk"
)

func main() {
	capsule.SetHandleHTTP(Handle)
}

// Handle function 
func Handle(param capsule.HTTPRequest) (capsule.HTTPResponse, error) {
	
	htmlPage := "<h1>👋 Hello World! 🌍</h1><h2>"+capsule.GetEnv("MESSAGE")+"</h2>"

	return capsule.HTTPResponse{
		TextBody: htmlPage,
		Headers: `{
			"Content-Type": "text/html; charset=utf-8",
			"Cache-Control": "no-cache",
			"X-Powered-By": "capsule-module-sdk"
		}`,
		StatusCode: 200,
	}, nil
}
