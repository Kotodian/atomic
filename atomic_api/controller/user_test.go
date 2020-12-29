package controller

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestWebUser(t *testing.T) {
	g := gin.Default()
	WebUser(g, 8081)
}
