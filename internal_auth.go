package internalauth

import (
	"github.com/corytech/go-openapi"
	"github.com/gin-gonic/gin"
)

func InternalAuthHeaderCheck(internalAuthToken string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		internalAuthHeader := ctx.GetHeader("Internal-Authorization")

		if internalAuthHeader != internalAuthToken {
			openapi.NewError(AuthentificationFailed).Send(ctx)
			return
		}

		ctx.Next()
	}
}
