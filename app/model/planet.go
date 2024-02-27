package model

import (
	"time"
)

type Planet struct {
	Name string
	/*	Image      string*/
	Allenergy  float64
	Restenergy float64
	Username   string //`gorm:"size:255"`
}

// 初始化一个星球
func (planet *Planet) Initializeplanet( /*planetimage string,*/ user User, planetname string) {
	planet.Username = user.Name
	/*planet.Image = planetimage*/
	planet.Name = planetname
	planet.Allenergy = 1
	planet.Restenergy = planet.Allenergy
}

// 创建新星球
func Createplanet( /*planetimage string,*/ planetname string, user User) {
	var planet Planet
	planet.Initializeplanet( /*planetimage, */ user, planetname)
	DB.Create(&planet)
}

// 删除星球
func Deleteplanet(username string, planetname string) {
	var planet Planet
	var mainland Mainland
	var animinal Animinal
	var goodbuilding Goodbuilding
	var badbuilding Badbuilding
	var plant Plant
	DB.Where("name=? AND username=? ", planetname, username).Delete(&planet)
	DB.Where("planetname=? AND username=? ", planetname, username).Delete(&mainland)
	DB.Where("planetname=? AND username=? ", planetname, username).Delete(&animinal)
	DB.Where("planetname=? AND username=? ", planetname, username).Delete(&goodbuilding)
	DB.Where("planetname=? AND username=? ", planetname, username).Delete(&badbuilding)
	DB.Where("planetname=? AND username=? ", planetname, username).Delete(&plant)
}

// 计算能量值,每隔一天加一定的能量值
func Planetenergy(planetname string, user User) {
	var a Planet
	var b Mainland
	// 创建一个每隔一天执行一次的定时器
	ticker := time.NewTicker(24 * time.Hour)
	// 启动一个 goroutine 来执行定时任务
	go func() {
		for {
			select {
			case <-ticker.C:
				// 每隔一天执行一次增加数据的操作
				DB.Where("name=? AND username=? ", user.Name, planetname).Find(a)
				DB.Where("name=? AND username=? AND planetname=? AND Status = true", "西伦瑞亚", user.Name, planetname).Find(b)
				a.Allenergy = b.Allproduct + a.Allenergy
				a.Restenergy = a.Restenergy + b.Allproduct
				DB.Where("name=? AND username=? AND planetname=? AND Status = true", "米尔勒拉", user.Name, planetname).Find(b)
				a.Allenergy = b.Allproduct + a.Allenergy
				a.Restenergy = a.Restenergy + b.Allproduct
				DB.Where("name=? AND username=? AND planetname=? AND Status = true", "乌兰宇蒂", user.Name, planetname).Find(b)
				a.Allenergy = b.Allproduct + a.Allenergy
				a.Restenergy = a.Restenergy + b.Allproduct
				DB.Where("name=? AND username=? AND planetname=? AND Status = true", "碦拉玛干", user.Name, planetname).Find(b)
				a.Allenergy = b.Allproduct + a.Allenergy
				a.Restenergy = a.Restenergy + b.Allproduct
				DB.Where("name=? AND username=? AND planetname=? AND Status = true", "云格雷诺", user.Name, planetname).Find(b)
				a.Allenergy = b.Allproduct + a.Allenergy
				a.Restenergy = a.Restenergy + b.Allproduct
				DB.Save(a)
			}
		}
	}()
	// 阻塞主 goroutine，使程序持续运行
	select {}
}

// 初始化一个星球
/*func (planet *Planet) Initializeplanet(planetimage string, planetname string, username string) {
	planet.Username = username
	planet.Image = planetimage
	planet.Name = planetname
	//第一个
	planet.Mainlands[0].Name = "西伦瑞亚" //初始化5个大洲的称号，气候，地形
	planet.Mainlands[0].Climate = "热带雨林气候"
	planet.Mainlands[0].Terrian = "森林"
	planet.Mainlands[0].Successclean = 10
	//第二个
	planet.Mainlands[1].Name = "米尔勒拉"
	planet.Mainlands[1].Climate = "高原山地气候"
	planet.Mainlands[1].Terrian = "山地，丘陵"
	planet.Mainlands[1].Price = 10
	planet.Mainlands[1].Successclean = 20
	//第三个
	planet.Mainlands[2].Name = "乌兰宇蒂"
	planet.Mainlands[2].Climate = "热带草原气候"
	planet.Mainlands[2].Terrian = "草原"
	planet.Mainlands[2].Price = 20
	planet.Mainlands[2].Successclean = 30
	//第四个
	planet.Mainlands[3].Name = "碦拉玛干"
	planet.Mainlands[3].Climate = "热带沙漠气候"
	planet.Mainlands[3].Terrian = "沙漠"
	planet.Mainlands[3].Price = 30
	planet.Mainlands[3].Successclean = 40
	//第五个
	planet.Mainlands[4].Name = "云格雷诺"
	planet.Mainlands[4].Climate = "极地气候"
	planet.Mainlands[4].Terrian = "冰川"
	planet.Mainlands[4].Price = 40
	planet.Mainlands[4].Successclean = 50
	var i int
	for i = 1; i < 5; i++ {
		planet.Mainlands[i].Planetname = planetname
		planet.Mainlands[i].Status = false
	}
	planet.Allenergy = 1
	planet.Restenergy = planet.Allenergy
}
*/
// 星球总净化量和剩余净化量的计算
/*func (planet *Planet) Allenergynumber() {
	var i int
	for i = 0; i < 5; i++ {
		if planet.Mainlands[i].Status == true {
			planet.Restenergy = planet.Restenergy + planet.Mainlands[i].Allproduct
			planet.Allenergy = planet.Allenergy + planet.Mainlands[i].Allproduct
		}
	}
}
*/
