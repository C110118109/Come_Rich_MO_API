package middleware

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/jwe"
	"eirc.app/internal/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Verify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		privateKey := "-----BEGIN PRIVATE KEY-----\n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\nix/Ccf7Enn38xwIXaZjumpva1dg7bdc=\n-----END PRIVATE KEY-----"
		j := &jwe.JWT{
			PrivateKey: privateKey,
			Token:      ctx.GetHeader("Authorization"),
		}

		if len(j.Token) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, code.GetCodeMessage(code.JWTRejected, "jwe is null"))
			return
		}

		j, err := j.Verify()
		if err != nil {
			log.Error(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, code.GetCodeMessage(code.JWTRejected, err.Error()))
			return
		}

		ctx.Set("account", j.Other["account"])
		ctx.Set("account_id", j.Other["account_id"])
		ctx.Set("company_id", j.Other["company_id"])
		ctx.Next()
	}
}
