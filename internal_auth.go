package internalauth

import (
	"os"

	"github.com/corytech/go-openapi"
	"github.com/gin-gonic/gin"
)

func InternalAuthHeaderCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		internalAuthToken := os.Getenv("INTERNAL_AUTH_TOKEN")
		internalAuthHeader := ctx.GetHeader("Internal-Authorization")

		if internalAuthHeader != internalAuthToken {
			openapi.NewError(AuthentificationFailed).Send(ctx)
			return
		}

		ctx.Next()
	}
}
