package routers

import (
	"github.com/gin-gonic/gin"
	"miniproject/app/controller"
)

func Routerinit() *gin.Engine {
	a := gin.Default()
	//登录和注册
	a.POST("/signup/sendcode", controller.Sendcode)
	a.POST("/signup", controller.Solvehttpsignup)
	controller.Solvehttpsignpassword(a)
	controller.Solvelogin(a)
	a.POST("/login/regettestcode", controller.Regettestcode)
	a.POST("/login/next", controller.Checkandnext)
	controller.Resetpassword(a)
	//星球的展示
	a.POST("/planet/name", controller.Nameplanet)
	a.GET("/planet", controller.Zhujiemian)
	//大陆展示
	a.GET("/mainland/xilun", controller.Getxilun)
	a.GET("/mainland/mier", controller.Getmier)
	a.GET("/mainland/wulan", controller.Getwulan)
	a.GET("/mainland/kala", controller.Getkala)
	a.GET("/mainland/yunluo", controller.Getyunluo)
	a.POST("/mainland/report", controller.Getreport)
	//仓库
	a.GET("/home/animinial", controller.Showaniminial)
	a.GET("/home/plants", controller.Showplants)
	a.GET("/home/buildings", controller.Showbuidlings)
	a.POST("/home/upgrade", controller.Upgradebuidling)
	//商店
	a.GET("/shop/animinial", controller.Getshopaniminial)
	a.GET("/shop/plants", controller.Getshopplants)
	a.GET("/shop/buildings", controller.Getbuilding)
	a.POST("/shop/animinial/buy", controller.Buyaniminal)
	a.POST("/shop/plant/buy", controller.Buyplant)
	a.POST("/shop/building/buy", controller.Buybuilding)
	return a
}
