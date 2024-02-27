package jwt1

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

type Usercliams struct {
	jwt.RegisteredClaims
	Username string
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 允许所有来源
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// 允许特定的HTTP方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")

		// 允许特定的HTTP头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// 允许携带凭证（例如，cookies）
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 处理OPTIONS预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		// 继续处理请求
		c.Next()
	}
}

func Jwtcheck(r *gin.Engine) {
	r.Use(func(c *gin.Context) {
		//不用验证的部分
		if c.Request.URL.Path == "/signup/sendcode" {
			return
		}
		if c.Request.URL.Path == "/signup" {
			return
		}
		if c.Request.URL.Path == "/signup/getpassword" {
			return
		}
		if c.Request.URL.Path == "/login" {
			return
		}
		if c.Request.URL.Path == "/login/regettestcode" {
			return
		}
		if c.Request.URL.Path == "/login/next" {
			return
		}
		if c.Request.URL.Path == "/login/resetpassword" {
			return
		}
		//验证的部分
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.String(http.StatusUnauthorized, "你未登录")
			return
		}
		segs := strings.Split(auth, " ")
		if len(segs) != 2 {
			c.String(http.StatusUnauthorized, "你未登录")
			return
		}
		tokenStr := segs[1]
		var claims = &Usercliams{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("muximiniproject"), nil
		})
		if err != nil {
			c.String(http.StatusUnauthorized, "你没登录")
		}
		if token == nil || !token.Valid || claims.Username == "" {
			c.String(http.StatusUnauthorized, "你未登录")
			return
		}
	})

}
