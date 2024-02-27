package controller

import (
	"github.com/gin-gonic/gin"
	"miniproject/app/model"
	"net/http"
)

// 命名星球 /planet/name POST
func Nameplanet(c *gin.Context) {
	var planetname string
	var i int
	planetname = c.PostForm("planetname")
	model.Planetname = planetname
	model.Createplanet(planetname, model.User1)
	model.Createmainlands(model.User1, planetname)
	model.Creategoodbuildings(model.User1.Name, planetname)
	model.Createbadbuildings(model.User1.Name, planetname)
	model.Creatplants(model.User1.Name, planetname)
	model.Createanimals(model.User1.Name, planetname)
	for i = 0; i < 5; i++ {
		model.Createreport(model.User1, planetname, model.Mainlands[i].Name)
	}
}

// 主界面 /planet GET
func Zhujiemian(c *gin.Context) {
	var planet model.Planet
	result := model.DB.Where("username = ?", model.User1.Name).Find(planet)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "查询星球数据错误"})
		return
	}
	model.Planetenergy(planet.Name, model.User1)
	c.JSON(200, gin.H{
		"星球名":  planet.Name,
		"总能量":  planet.Allenergy,
		"剩余能量": planet.Restenergy,
	})
}
