package main

import (
	"atomic/atomic_api/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	controller.WebUser(gin.Default(), 12002)
}
