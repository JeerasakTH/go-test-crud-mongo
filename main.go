package main

import (
	"github.com/JeerasakTH/go-test-crud/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.Router(r)

	r.Run()
}
