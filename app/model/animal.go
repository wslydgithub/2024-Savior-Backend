package model

import (
	"github.com/gin-gonic/gin"
	"miniproject/app/common/qiniuyun"
)

type Animinal struct {
	Name         string
	Price        float64
	Production   float64
	Introduction string
	Image        string
	Number       int
	Planetname   string
	Mainlandname string
	Username     string
}

// 初始化各种动物
func Initialanimals(username string, planetname string) {
	qiniuyun.Animinalimages()
	var i int
	for i = 0; i < 20; i++ {
		Animals[i].Username = username
		Animals[i].Planetname = planetname
		Animals[i].Image = Animalsimages[i]
	}
}

// 将动物存入数据库中
func Createanimals(username string, planetname string) {
	Initialanimals(username, planetname)
	DB.Create(&Animals)
}

// 购买动物
func Buyanimals(c1 *gin.Context, user User, animinalname string, planetname string) {
	var c Animinal
	var d Planet
	planet := DB.Where("username =? AND name =?", user.Name, planetname).Find(&d)
	if planet.Error != nil {
		/*该处为http报错*/
		return
	}
	animinal := DB.Where("name =? AND planetname =? AND username =?", animinalname, planetname, user.Name).Find(&c)
	if animinal.Error != nil {
		/*该处为http报错*/
		return
	}
	if d.Restenergy >= c.Price {
		d.Restenergy = d.Restenergy - c.Price
		c.Number++
		DB.Save(&d)
		DB.Save(&c)
	} else {
		c1.JSON(200, gin.H{
			"不能购买动物": "能量不足",
		})
	}
}
