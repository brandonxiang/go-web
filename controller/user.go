package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/brandonxiang/go-example/model"
)

func User(c *gin.Context) {
  db := model.Init()
	users := model.GetUser(db)


	//  model.GetUser()
	 c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}