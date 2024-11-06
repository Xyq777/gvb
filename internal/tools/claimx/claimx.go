package claimx

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gvb/internal/global"
	"gvb/tools/jwt"
)

func GetClaim(c *gin.Context) (*jwt.CustomClaims, error) {
	_claims, ok := c.Get("claims")
	if !ok {
		global.Log.Error("未设置claims")
		return nil, errors.New("未设置claims")
	}
	claims := _claims.(*jwt.CustomClaims)
	return claims, nil
}
