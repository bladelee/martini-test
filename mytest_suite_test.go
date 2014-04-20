package main_test

import (
	"github.com/go-martini/martini"

	"github.com/martini-contrib/render"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"

	"io"
	"testing"
)

var (
	response *httptest.ResponseRecorder
)

func TestMytest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mytest Suite")
}

func Request(method string, route string, handler martini.Handler) {
	m := martini.Classic()
	m.Get(route, handler)
	m.Use(render.Renderer())
	request, _ := http.NewRequest(method, route, nil)
	response = httptest.NewRecorder()
	m.ServeHTTP(response, request)
}

func PostRequest(method string, route string, jsonStruct interface{}, handler martini.Handler, body io.Reader) {
	m := martini.Classic()
	m.Post(route, jsonStruct, handler)
	m.Use(render.Renderer())
	request, _ := http.NewRequest(method, route, body)
	response = httptest.NewRecorder()
	m.ServeHTTP(response, request)
}

func PostRequestCommon(method string, route string, handler martini.Handler, body io.Reader) {
	m := martini.Classic()
	m.Post(route, handler)
	m.Use(render.Renderer())
	request, _ := http.NewRequest(method, route, body)
	response = httptest.NewRecorder()
	m.ServeHTTP(response, request)
}
