package model

import (
	"github.com/gin-gonic/gin"
	"miniproject/app/common/qiniuyun"
)

type Plant struct {
	Name         string
	Price        float64
	Production   float64
	Introduction string
	Image        string
	Mainlandname string //`gorm:"size:255"`
	Number       int
	Planetname   string
	Username     string
}

// 初始化植物
func Initialplants(username string, planetname string) {
	qiniuyun.Plantimages()
	var i int
	for i = 0; i < 12; i++ {
		Plants[i].Username = username
		Plants[i].Planetname = planetname
		Plants[i].Image = Plantimages[i]
	}
}

// 将植物存入数据库中
func Creatplants(username string, planetname string) {
	Initialplants(username, planetname)
	DB.Create(&Plants)
}

// 购买植物
func Buyplants(c *gin.Context, user User, plantname string, planetname string) {
	var a Plant
	var b Planet
	planet := DB.Where("username =? AND name =?", user.Name, planetname).Find(&b)
	if planet.Error != nil {
		/*该处为http报错*/
		return
	}
	plant := DB.Where("name =? AND planetname =? AND username =?", plantname, planetname, user.Name).Find(&a)
	if plant.Error != nil {
		/*该处为http报错*/
		return
	}
	if b.Restenergy >= a.Price {
		b.Restenergy = b.Restenergy - a.Price
		a.Number++
		DB.Save(&b)
		DB.Save(&a)
	} else {
		c.JSON(200, gin.H{
			"无法购买植物": "余额不足",
		})
	}
}

/*func (plant *Plant) Initial(name string, price int, introduction string, image string, num int) {
	plant.Name = name
	plant.Price = price
	plant.Introduction = introduction
	plant.Image = image
	plant.Number = num
}
*/
