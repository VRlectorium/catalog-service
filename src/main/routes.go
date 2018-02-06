package main

import (
	"encoding/json"
	"net/http"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

// header of our json
func setHeader(ctx *fasthttp.RequestCtx, code int) {
	ctx.SetContentType("application/json; charset=utf-8")
	ctx.Response.SetStatusCode(code)
}

// GET /version
func versionEndpoint(ctx *fasthttp.RequestCtx) {
	setHeader(ctx, http.StatusOK)
	js, err := json.Marshal(Version{Version: "0.0"})
	if err == nil {
		ctx.Response.AppendBody(js)
	}
}

func notFoundEndpoint(ctx *fasthttp.RequestCtx) {
	setHeader(ctx, http.StatusNotFound)
	js, err := json.Marshal(Error{Error: "Not Found"})
	if err == nil {
		ctx.Response.AppendBody(js)
	}
}

// all endpoints
func Routes() *fasthttprouter.Router {
	router := fasthttprouter.New()
	router.GET("/version", versionEndpoint)
	router.NotFound = notFoundEndpoint
	return router
}
