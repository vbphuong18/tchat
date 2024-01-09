package middleware

import (
	"TChat/pkg/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func HTTPAuthentication(ctx *gin.Context) {
	cookie, err := ctx.Cookie("token")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	claims := &utils.Claims{}
	tkn, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (any, error) {
		return utils.JwtKey, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.Next()
}
