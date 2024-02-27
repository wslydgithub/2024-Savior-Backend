package model

import (
	"github.com/gin-gonic/gin"
	"miniproject/app/common/qiniuyun"
	"net/http"
)

type Goodbuilding struct {
	Name         string
	Price        float64
	Image        string
	Grade        int
	Introduction string
	Number       int
	Production   float64
	Planetname   string
	Username     string
	Mainlandname string
}

// 初始化环保建筑物
func Initialgoodbuildings(username string, planetname string) {
	qiniuyun.Goodbuildingimages()
	var i int
	j := 0
	for i = 0; i < 20; i++ {
		Goodbuildings[i].Username = username
		Goodbuildings[i].Planetname = planetname
		if j == 4 {
			j = 0
		}
		Goodbuildings[i].Image = Goodbuildingimages[j]
		j++
	}
}

// 存入环保建筑物的数据
func Creategoodbuildings(username string, planetname string) {
	Initialgoodbuildings(username, planetname)
	DB.Create(&Goodbuildings)
}

// 购买环保建筑物
func Buygoodbuildingss(c1 *gin.Context, user User, goodbuildingname string, mainlandname string, planetname string) {
	var c Goodbuilding
	var d Planet
	planet := DB.Where("username =? AND name =?", user.Name, planetname).Find(&d)
	if planet.Error != nil {
		c1.JSON(http.StatusBadRequest, gin.H{
			"error": "查找不到对应的星球",
		})
		return
	}
	goodbuilding := DB.Where("name =? AND planetname =? AND username =? AND mainlandname =?", goodbuildingname, planetname, user.Name, mainlandname).Find(&c)
	if goodbuilding.Error != nil {
		c1.JSON(http.StatusBadRequest, gin.H{
			"error": "查找不到对应的建筑物",
		})
		return
	}
	if d.Restenergy >= c.Price {
		d.Restenergy = d.Restenergy - c.Price
		c.Number++
		DB.Save(&d)
		DB.Save(&c)
	} else {
		c1.JSON(http.StatusBadRequest, gin.H{
			"不能购买建筑物": "能量不足",
		})
		return
	}
}

// 升级环保建筑物
func Upgradegoodbuildings(user User, goodbuildingname string, mainlandname string, planetname string) {
	var c Goodbuilding
	var d Planet
	planet := DB.Where("username =? AND name =?", user.Name, planetname).Find(&d)
	if planet.Error != nil {
		/*该处为http报错*/
		return
	}
	goodbuilding := DB.Where("name =? AND planetname =? AND username =? AND mainlandname =?", goodbuildingname, planetname, user.Name, mainlandname).Find(&c)
	if goodbuilding.Error != nil {
		/*该处为http报错*/
		return
	}
	c.Grade++
	c.Production = c.Production * 1.1
	DB.Save(c)
}

/*unc (goodbuilding *Goodbuilding) Initial(name string, price int, introduction string, image string, num int, production float64) {
	goodbuilding.Name = name
	goodbuilding.Price = price
	goodbuilding.Introduction = introduction
	goodbuilding.Image = image
	goodbuilding.Number = num
	goodbuilding.Production = production
	goodbuilding.Grade = 1
}
*/
