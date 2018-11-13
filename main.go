package main

import (
	"github.com/codegangsta/martini"
	"github.com/henosteven/heigo/httpservice"
)

func main() {
	m := martini.Classic()
	m.Get("/", httpservice.Hello)
	m.Get("/get", httpservice.Get)
	m.Get("/set", httpservice.Set)
	m.Run()
}
