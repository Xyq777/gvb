package middleware

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/models/ctype"
	"gvb/internal/models/serializition/res"
	"gvb/tools/jwt"
	"strings"
)

func JwtAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		// 检查头部是否以 "Bearer " 开头
		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			callback.FAIL(res.AuthFailed, "认证类型不符", c)
			c.Abort()
			return
		}
		token := strings.TrimPrefix(authHeader, bearerPrefix)
		claims, err := jwt.ParseToken(token)
		if err != nil {
			callback.FAIL(res.AuthFailed, res.CodeMsg(res.AuthFailed), c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
func RoleAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		role := claims.(jwt.CustomClaims).Role
		if role != ctype.PermissionAdmin {
			callback.FAIL(res.AuthFailed, "权限不足", c)
			c.Abort()
			return
		}

		c.Next()
	}
}
