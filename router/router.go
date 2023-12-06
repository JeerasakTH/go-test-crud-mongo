package router

import (
	"github.com/JeerasakTH/go-test-crud/controller"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.GET("/user", controller.GetUser())
}
