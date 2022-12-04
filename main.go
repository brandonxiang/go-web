package main

import (
    "net/http"
    "log"

    "github.com/julienschmidt/httprouter"
		"github.com/brandonxiang/go-example/controller"
)





func main() {
    router := httprouter.New()
    router.GET("/", controller.Welcome)
    router.GET("/hello/:name", controller.hello)

    log.Fatal(http.ListenAndServe(":8080", router))
}