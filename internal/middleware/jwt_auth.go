package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/global"
	"gvb/internal/models/ctype"
	"gvb/internal/models/dto/res"
	"gvb/tools/jwt"
	"strconv"
	"strings"
	"time"
)

func TokenRefresh(c *gin.Context) {
	rt, err := c.Cookie("refreshToken")
	if err != nil {
		callback.FAIL(res.AuthFailed, res.CodeMsg(res.AuthFailed), c, errors.New("cookie没有refreshToken"))
		c.Abort()
		return
	}

	claims, err := jwt.ParseToken(rt)
	if err != nil {
		callback.FAIL(res.AuthFailed, res.CodeMsg(res.AuthFailed), c)
		c.Abort()
		return
	}
	jti := claims.ID
	exist, err := global.Redis.HExists(c, strconv.Itoa(int(claims.UserID)), jti).Result()
	if err != nil {
		global.Log.Errorln(err)
		callback.FAIL(res.RedisGetFailed, res.CodeMsg(res.RedisGetFailed), c)
		return
	}
	if !exist {
		callback.FAIL(res.AuthFailed, res.CodeMsg(res.AuthFailed), c)
		return
	}
	// 生成新的token
	exp := time.Duration(global.Config.System.Jwt.ATExpires) * time.Second
	at, err := jwt.GenAccessToken(claims.Payload, exp)
	if err != nil {
		global.Log.Errorln(err)
		callback.FAIL(res.TokenGenerateFailed, res.CodeMsg(res.TokenGenerateFailed), c)
		return
	}
	callback.OK(at, c)
}
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
