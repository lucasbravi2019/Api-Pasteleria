package core

import "net/http"

type Route struct {
	Path        string
	HandlerFunc http.HandlerFunc
	Method      string
}

type Routes []Route
