package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onnytra/first-go/external"
	"github.com/onnytra/first-go/internal/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			log.Println("authorization empty")
			utils.HTTPResponse(ctx, http.StatusUnauthorized, "Unauthorized", nil, nil)
			ctx.Abort()
			return
		}

		const bearerPrefix = "Bearer "
		if len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
			log.Println("invalid authorization format")
			utils.HTTPResponse(ctx, http.StatusUnauthorized, "Unauthorized", nil, nil)
			ctx.Abort()
			return
		}

		accessToken := authHeader[len(bearerPrefix):]
		userData, err := external.ValidateToken(ctx.Request.Context(), accessToken)
		if err != nil {
			log.Println(err)
			utils.HTTPResponse(ctx, http.StatusUnauthorized, "Unauthorized", nil, nil)
			ctx.Abort()
			return
		}

		ctx.Set("user", userData)
		ctx.Next()
	}
}
