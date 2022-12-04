package main

import (
    "net/http"
    "log"

    "github.com/julienschmidt/httprouter"
		"github.com/brandonxiang/go-example/controller"
)





func main() {
    router := httprouter.New()
    router.GET("/", controller.Index)
    router.GET("/hello/:name", controller.Hello)

    log.Fatal(http.ListenAndServe(":8080", router))
}