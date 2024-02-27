package controller

import (
	"github.com/gin-gonic/gin"
	"miniproject/app/common/qiniuyun"
	"miniproject/app/model"
	"net/http"
)

// 动物展示 /home/animinial GET
func Showaniminial(c *gin.Context) {
	var animinials [25]model.Animinal
	result := model.DB.Where("username = ? AND planetname = ?", model.User1.Name, model.Planetname).Find(animinials)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "仓库查询动物数据出现问题",
		})
		return
	}
	var usersMap []gin.H
	for i, animinal := range animinials {
		usersMap = append(usersMap, gin.H{
			"name":  animinal.Name,
			"图片":    animinal.Image,
			"token": qiniuyun.Gettoken(model.Animaladdress[i]),
			"数量":    animinal.Number,
		})
	}
	c.JSON(200, usersMap)
}

// 植物展示 /home/plants GET
func Showplants(c *gin.Context) {
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
			"数量":    plant.Number,
		})
	}
	c.JSON(200, usersMap)
}

// 建筑物展示 /home/buildings GET
func Showbuidlings(c *gin.Context) {
	var bu [20]model.Goodbuilding
	result := model.DB.Where("username=? AND planetname =?", model.User1.Name, model.Planetname).Find(bu)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "商店读取建筑物信息错误"})
		return
	}
	c.JSON(200, gin.H{
		"名称1":    bu[0].Name,
		"图片1":    bu[0].Image,
		"token1": qiniuyun.Gettoken(model.Goodbuildingaddress[0]),
		"名称2":    bu[1].Name,
		"图片2":    bu[1].Image,
		"token2": qiniuyun.Gettoken(model.Goodbuildingaddress[1]),
		"名称3":    bu[2].Name,
		"图片3":    bu[2].Image,
		"token3": qiniuyun.Gettoken(model.Goodbuildingaddress[2]),
		"名称4":    bu[3].Name,
		"图片4":    bu[3].Image,
		"token4": qiniuyun.Gettoken(model.Goodbuildingaddress[3]),
	})
}

// 升级建筑物 /home/upgrade POST
func Upgradebuidling(c *gin.Context) {
	var a string
	a = c.PostForm("buidingname")
	model.Upgradegoodbuildings(model.User1, a, "西伦瑞亚", model.Planetname)
	model.Upgradegoodbuildings(model.User1, a, "米尔勒拉", model.Planetname)
	model.Upgradegoodbuildings(model.User1, a, "乌兰宇蒂", model.Planetname)
	model.Upgradegoodbuildings(model.User1, a, "碦拉玛干", model.Planetname)
	model.Upgradegoodbuildings(model.User1, a, "云格雷诺", model.Planetname)

}
