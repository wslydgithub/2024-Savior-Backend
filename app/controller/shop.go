package controller

import (
	"github.com/gin-gonic/gin"
	"miniproject/app/common/qiniuyun"
	"miniproject/app/model"
	"net/http"
)

// 动物 /shop/animinial GET
func Getshopaniminial(c *gin.Context) {
	var as [25]model.Animinal
	result := model.DB.Where("username=? AND planetname =?", model.User1.Name, model.Planetname).Find(as)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "商店读取动物信息错误"})
		return
	}
	var usersMap []gin.H
	for i, animinal := range as {
		usersMap = append(usersMap, gin.H{
			"name":  animinal.Name,
			"图片":    animinal.Image,
			"token": qiniuyun.Gettoken(model.Animaladdress[i]),
			"价格":    animinal.Price,
		})
	}
	c.JSON(200, usersMap)
}

// 植物 /shop/plants GET
func Getshopplants(c *gin.Context) {
	var pl [15]model.Plant
	result := model.DB.Where("username=? AND planetname =?", model.User1.Name, model.Planetname).Find(pl)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "商店读取植物信息错误"})
		return
	}
	var usersMap []gin.H
	for i, plant := range pl {
		usersMap = append(usersMap, gin.H{
			"name":  plant.Name,
			"图片":    plant.Image,
			"token": qiniuyun.Gettoken(model.Plantaddress[i]),
			"价格":    plant.Price,
		})
	}
	c.JSON(200, usersMap)
}

// 建筑物 /shop/buildings GET
func Getbuilding(c *gin.Context) {
	var bu [20]model.Goodbuilding
	result := model.DB.Where("username=? AND planetname =?", model.User1.Name, model.Planetname).Find(bu)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "商店读取建筑物信息错误"})
		return
	}
	c.JSON(200, gin.H{
		"名称1":    bu[0].Name,
		"价格1":    bu[0].Price,
		"图片1":    bu[0].Image,
		"token1": qiniuyun.Gettoken(model.Goodbuildingaddress[0]),
		"名称2":    bu[1].Name,
		"价格2":    bu[1].Price,
		"图片2":    bu[1].Image,
		"token2": qiniuyun.Gettoken(model.Goodbuildingaddress[1]),
		"名称3":    bu[2].Name,
		"价格3":    bu[2].Price,
		"图片3":    bu[2].Image,
		"token3": qiniuyun.Gettoken(model.Goodbuildingaddress[2]),
		"名称4":    bu[3].Name,
		"价格4":    bu[3].Price,
		"图片4":    bu[3].Image,
		"token4": qiniuyun.Gettoken(model.Goodbuildingaddress[3]),
	})
}

// 购买动物 /shop/animinial/buy POST
func Buyaniminal(c *gin.Context) {
	var a string
	a = c.PostForm("animinalname")
	model.Buyanimals(c, model.User1, a, model.Planetname)
}

// 购买植物 /shop/plant/buy POST
func Buyplant(c *gin.Context) {
	var a string
	a = c.PostForm("plantname")
	model.Buyplants(c, model.User1, a, model.Planetname)
}

// 购买建筑物 /shop/building/buy POST
func Buybuilding(c *gin.Context) {
	var a string
	a = c.PostForm("goodbuildingname")
	model.Buygoodbuildingss(c, model.User1, a, "西伦瑞亚", model.Planetname)
	model.Buygoodbuildingss(c, model.User1, a, "米尔勒拉", model.Planetname)
	model.Buygoodbuildingss(c, model.User1, a, "乌兰宇蒂", model.Planetname)
	model.Buygoodbuildingss(c, model.User1, a, "碦拉玛干", model.Planetname)
	model.Buygoodbuildingss(c, model.User1, a, "云格雷诺", model.Planetname)
}
