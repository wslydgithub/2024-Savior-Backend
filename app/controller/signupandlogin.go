package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"miniproject/app/common/email"
	"miniproject/app/common/jwt1"
	"miniproject/app/model"
	"net/http"
)

// /signup/sendcode POST
// 获取验证码ok
func Sendcode(c *gin.Context) {
	var user model.User
	model.Email = c.PostForm("username")
	model.DB.Where("name = ?", model.Email).Find(user)
	if user.Name != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "该用户已存在"})
		return
	}
	email.Send(model.Email)
}

// /signup POST
// 检验+转到设置密码ok
func Solvehttpsignup(c *gin.Context) {
	model.Codehttp = c.PostForm("testcode")
	if model.Codehttp == string(model.Code) {
		model.Usersign.Email = model.Email
		c.Redirect(http.StatusMovedPermanently, "/signup/password")
	} else {
		c.String(http.StatusUnauthorized, "你的验证码有问题")
		return
	}
}

// 读取密码并注册+得到token ok
func Solvehttpsignpassword(r *gin.Engine) {
	r.POST("/signup/getpassword", func(c *gin.Context) {
		model.Password = c.PostForm("password")
		model.Usersign.Password = model.Password
		model.Adduser()
		model.User1.Name = model.Email
		model.User1.Password = model.Password
		var claims jwt1.Usercliams
		claims.Username = model.Email
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte("muximiniproject"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "获取token有问题"})
			return
		}
		c.Header("jwt1-token", tokenString)
		r.Use(jwt1.CORSMiddleware())
		c.Redirect(http.StatusMovedPermanently, "saveplanets/introduction")
	})
}

// /login POST
// 登录+得到token
func Solvelogin(r *gin.Engine) {
	r.POST("/login", func(c *gin.Context) {
		model.Email = c.PostForm("username")
		model.Password = c.PostForm("password")
		var user model.User
		result := model.DB.Where("name=?", model.Email).Where("password=?", model.Password).First(&user)
		if result.Error != nil {
			c.String(http.StatusUnauthorized, "你的邮箱或密码不正确")
			return
		} else {
			model.User1.Name = model.Email
			model.User1.Password = model.Password
			var claims jwt1.Usercliams
			claims.Username = model.Email
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString([]byte("muximiniproject"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "获取token有问题"})
				return
			}
			c.Header("jwt1-token", tokenString)
			r.Use(jwt1.CORSMiddleware())
			c.Redirect(http.StatusMovedPermanently, "/planet")
		}
	})
}

// 找回密码 POST  /login/regettestcode
func Regettestcode(c *gin.Context) {
	var user model.User
	model.Email = c.PostForm("username")
	model.DB.Where("name = ?", model.Email).Find(user)
	if user.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "你还没有注册"})
		return
	}
	email.Send(model.Email)
}

// 找回验证码时，检查验证码是否正确 POST /login/next
func Checkandnext(c *gin.Context) {
	model.Codehttp = c.PostForm("testcode")
	if model.Codehttp == string(model.Code) {
		model.Usersign.Email = model.Email
		c.Redirect(http.StatusMovedPermanently, "/login/resetpassword")
	} else {
		c.String(http.StatusUnauthorized, "你的验证码有问题")
		return
	}
}

// 找回密码后重新设置密码
func Resetpassword(r *gin.Engine) {
	r.POST("/login/resetpassword", func(c *gin.Context) {
		var user model.User
		var repassword string
		repassword = c.PostForm("repassword")
		model.Password = c.PostForm("password")
		model.Usersign.Password = model.Password
		if repassword != model.Password {
			c.JSON(http.StatusBadRequest, gin.H{"error": "两次输入的密码不一致"})
			return
		}
		model.User1.Name = model.Email
		model.User1.Password = model.Password
		result := model.DB.Where("name =? ", model.User1.Name).Find(user)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "查找用户失败，重设密码失败"})
			return
		}
		user.Password = model.User1.Password
		model.DB.Save(user)
		var claims jwt1.Usercliams
		claims.Username = model.Email
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte("muximiniproject"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "获取token有问题"})
			return
		}
		c.Header("jwt1-token", tokenString)
		r.Use(jwt1.CORSMiddleware())
		c.Redirect(http.StatusMovedPermanently, "saveplanets/introduction")
	})
}
