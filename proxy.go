package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"os"
)

func main() {
	port := getPort()
	fmt.Printf("# Proxy API\n")
	fmt.Printf("# Server listening on port %s...\n", port)
	fasthttp.ListenAndServe(":"+port, requestHandler)
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	targetURL := string(ctx.Request.Header.Peek("X-Proxy-Url"))
	ctx.Request.Header.Del("X-Proxy-Url")
	if targetURL == "" {
		sendErrorResponse(ctx, fasthttp.StatusBadRequest, "Missing X-Proxy-Url header")
		return
	}

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	ctx.Request.Header.CopyTo(&req.Header)
	req.SetBody(ctx.PostBody())
	req.SetRequestURI(targetURL)
	req.Header.SetUserAgent("Proxy API")

	if err := fasthttp.Do(req, resp); err != nil {
		sendErrorResponse(ctx, fasthttp.StatusInternalServerError, err.Error())
		return
	}

	resp.Header.CopyTo(&ctx.Response.Header)
	ctx.SetBody(resp.Body())
}

func sendErrorResponse(ctx *fasthttp.RequestCtx, statusCode int, message string) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	response := Response{
		Status:  false,
		Message: message,
	}
	_ = json.NewEncoder(ctx).Encode(response)
}

func getPort() string {
	if p := os.Getenv("PROXY_API_PORT"); p != "" {
		return p
	}
	return "9900"
}

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
