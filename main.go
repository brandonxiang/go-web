package main

import (
    "net/http"
    "log"

	"github.com/brandonxiang/go-example/controller"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/", controller.Index)
    r.GET("/hello/:name", controller.Hello)
    r.GET("/user", controller.User)
    r.Run()

    log.Fatal(http.ListenAndServe(":8080", r))
}