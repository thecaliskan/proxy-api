package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"os"
)

func main() {
	port := getPort()
	fmt.Printf("Server listening on port %s...\n", port)
	fasthttp.ListenAndServe(":"+port, receiveHandler)
}

func receiveHandler(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) != fasthttp.MethodPost {
		sendErrorResponse(ctx, fasthttp.StatusMethodNotAllowed, "Only POST requests are allowed")
		return
	}

	targetURL := string(ctx.Request.Header.Peek("proxy-url"))
	if targetURL == "" {
		sendErrorResponse(ctx, fasthttp.StatusBadRequest, "Missing proxy-url header")
		return
	}

	statusCode, body, err := fasthttp.Post(nil, targetURL, ctx.PostArgs())
	if err != nil {
		sendErrorResponse(ctx, fasthttp.StatusInternalServerError, "Failed to send: "+err.Error())
		return
	}

	ctx.SetStatusCode(statusCode)
	ctx.SetBody(body)
}

func sendErrorResponse(ctx *fasthttp.RequestCtx, statusCode int, message string) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	response := map[string]interface{}{"status": false, "message": message}
	_ = json.NewEncoder(ctx).Encode(response)
}

func getPort() string {
	if p := os.Getenv("PROXY_API_PORT"); p != "" {
		return p
	}
	return "9900"
}
