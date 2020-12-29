package middleware

import (
	"atomic/internal/atomic_error"
	"atomic/internal/auth/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTAuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")

		if authHeader == "" {
			ctx.JSON(http.StatusOK, atomic_error.ErrUserLogin)
			ctx.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if !(len(parts) != 2 && parts[0] == "Bearer") {
			ctx.JSON(http.StatusOK, atomic_error.ErrUserLogin)
			ctx.Abort()
			return
		}

		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			ctx.JSON(http.StatusOK, atomic_error.ErrUserLogin)
			ctx.Abort()
			return
		}

		ctx.Set("username", mc.Username)
		ctx.Next()
	}
}
