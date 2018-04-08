package server

import (
	"net/http"
	"github.com/NozomiSugiyama/image-placeholder-server/server/handler"
)

type Route struct {
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"/",
		handler.Index,
	},
	Route{
		"GET",
		"/image-samples",
		handler.ImageSamples,
	},
}
