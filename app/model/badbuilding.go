package model

import (
	"github.com/gin-gonic/gin"
	"miniproject/app/common/qiniuyun"
)

type Badbuilding struct {
	Name         string
	Image        string
	Introduction string
	Number       int
	Pollution    float64
	Price        float64
	Planetname   string
	Mainlandname string
	Username     string
}

// 初始化污染建筑物
func Initialbadbuildings(username string, planetname string) {
	qiniuyun.Badbuildingimages()
	var i int
	j := 0
	for i = 0; i < 20; i++ {
		Badbuildings[i].Username = username
		Badbuildings[i].Planetname = planetname
		if j == 4 {
			j = 0
		}
		Badbuildings[i].Image = Badbuildingimages[j]
	}
}

// 把污染建筑物存入数据库
func Createbadbuildings(username string, planetname string) {
	Initialbadbuildings(username, planetname)
	DB.Create(&Badbuildings)
}

// 拆除建筑物
func Chaibadbuilding(c *gin.Context, user Usersl, planetname string, badbuildingname string, mainlandname string) {
	var a Planet
	var b Badbuilding
	planet := DB.Where("username =? AND planetname =?", user.Email, planetname).Find(&a)
	if planet.Error != nil {
		/*此处为http报错*/
		return
	}
	badbuilding := DB.Where("username=? AND planetname =? AND mainlandname =? AND name=?", user.Email, planetname, mainlandname, badbuildingname).Find(&b)
	if badbuilding.Error != nil {
		/*此处为http报错*/
		return
	}
	if a.Restenergy >= b.Price {
		a.Restenergy = a.Restenergy - b.Price
		b.Number--
		DB.Save(&a)
		DB.Save(&b)
	} else {
		c.JSON(200, gin.H{
			"不能拆除": "能量不足",
		})
	}
}

/*func (badbuilding *Badbuilding) Initial(name string, pollution float64, introduction string, num int) {
	badbuilding.Name = name
	badbuilding.Pollution = pollution
	badbuilding.Introduction = introduction
	badbuilding.Number = num
}
*/
