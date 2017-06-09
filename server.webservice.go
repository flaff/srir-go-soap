package main

import (
	"net/http"
	"github.com/codegangsta/martini"
)

type WebService interface {
	GetPath() string
	WebDelete(params martini.Params) (int, string)
	WebGet(params martini.Params) (int, string)
	WebPost(params martini.Params, req *http.Request) (int, string)
	WebPut(params martini.Params, req *http.Request) (int, string)
	WebPatch(params martini.Params, req *http.Request) (int, string)
}

func RegisterWebService(webService WebService,
	classicMartini *martini.ClassicMartini) {
	path := webService.GetPath()

	classicMartini.Get(path, webService.WebGet)
	classicMartini.Get(path+"/:id", webService.WebGet)

	classicMartini.Post(path, webService.WebPost)
	classicMartini.Post(path+"/:id", webService.WebPost)

	classicMartini.Delete(path, webService.WebDelete)
	classicMartini.Delete(path+"/:id", webService.WebDelete)

	classicMartini.Put(path, webService.WebPut)
	classicMartini.Put(path+"/:id", webService.WebPut)

	classicMartini.Patch(path, webService.WebPatch)
	classicMartini.Patch(path+"/:id", webService.WebPatch)
}