package claimx

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/global"
	"gvb/tools/jwt"
)

func GetClaim(c *gin.Context) *jwt.CustomClaims {
	_claims, ok := c.Get("claims")
	if !ok {
		global.Log.Error("未设置claims")
		return nil
	}
	claims := _claims.(*jwt.CustomClaims)
	return claims
}
