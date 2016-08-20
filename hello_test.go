package main

import (
	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHello(t *testing.T) {
	r := gofight.New()

	r.GET("/").
		// enable debug mode
		SetDebug(true).
		Run(runLocal(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

			assert.Equal(t, "Hi there", string(r.Body.String()[:8]), "strings should be equal")
			assert.Equal(t, http.StatusOK, r.Code, "HTTP Status Code should be 200")
		})
}
