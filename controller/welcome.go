package controller

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/brandonxiang/go-example/model"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
	model.init();
}