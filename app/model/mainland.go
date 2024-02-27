package model

import "time"

type Mainland struct {
	Name  string
	Image string
	/*	Stageimage   string*/
	Status       bool    //注意默认的是全不未解锁
	Allproduct   float64 //暂无
	Allcleanrate float64 //暂无
	Climate      string
	Terrian      string
	Others       string
	Planetname   string
	Successclean float64
	Username     string
}

// 初始化五个大陆的基本信息
func Initialmainlands(username string, planetname string) {
	var i int
	for i = 0; i < 5; i++ {
		Mainlands[i].Image = Mainlandmages[i]
		Mainlands[i].Status = false
		Mainlands[i].Username = username
		Mainlands[i].Planetname = planetname
		/*	Mainlands[i].Stageimage = Stateimages[i*3]*/
	}
	//第一个
	Mainlands[0].Name = "西伦瑞亚" //初始化5个大洲的称号，气候，地形
	Mainlands[0].Climate = "热带雨林气候"
	Mainlands[0].Terrian = "森林"
	Mainlands[0].Others = "暂无" //等资料
	Mainlands[0].Status = true
	Mainlands[0].Successclean = 5
	Mainlands[0].Allproduct = 0
	Mainlands[0].Allcleanrate = 0
	//第二个
	Mainlands[1].Name = "米尔勒拉"
	Mainlands[1].Climate = "高原山地气候"
	Mainlands[1].Terrian = "山地，丘陵"
	Mainlands[1].Others = "暂无" //等资料
	Mainlands[1].Successclean = 5
	Mainlands[1].Allproduct = 0
	Mainlands[1].Allcleanrate = 0
	//第三个
	Mainlands[2].Name = "乌兰宇蒂"
	Mainlands[2].Climate = "热带草原气候"
	Mainlands[2].Terrian = "草原"
	Mainlands[2].Others = "暂无" //等资料
	Mainlands[2].Successclean = 10
	Mainlands[2].Allproduct = 0
	Mainlands[2].Allcleanrate = 0
	//第四个
	Mainlands[3].Name = "碦拉玛干"
	Mainlands[3].Climate = "热带沙漠气候"
	Mainlands[3].Terrian = "沙漠"
	Mainlands[3].Others = "暂无" //等资料
	Mainlands[3].Successclean = 10
	Mainlands[3].Allproduct = 0
	Mainlands[3].Allcleanrate = 0
	//第五个
	Mainlands[4].Name = "云格雷诺"
	Mainlands[4].Climate = "极地气候"
	Mainlands[4].Terrian = "冰川"
	Mainlands[4].Others = "暂无" //等资料
	Mainlands[4].Successclean = 20
	Mainlands[4].Allproduct = 0
	Mainlands[4].Allcleanrate = 0
}

// 将大陆信息存入数据库中
func Createmainlands(user User, planetname string) {
	Initialmainlands(user.Name, planetname)
	DB.Create(Mainlands)
}

// 计算产率和产量+解锁
func Mainlandenergy(user User, planetname string, mainlandname string) {
	var a [5]Animinal
	var b [3]Plant
	var e [4]Badbuilding
	var f [4]Goodbuilding
	var c Mainland
	// 创建一个每隔一天执行一次的定时器
	ticker := time.NewTicker(24 * time.Hour)
	// 启动一个 goroutine 来执行定时任务
	go func() {
		for {
			select {
			case <-ticker.C:
				var i int
				DB.Where("name = ? AND username = ? AND planetname = ?", mainlandname, user.Name, planetname).Find(&c)
				DB.Where("username=? AND planetname = ? AND mainlandname = ?", user.Name, planetname, mainlandname).Find(&a)
				DB.Where("username=? AND planetname = ? AND mainlandname = ?", user.Name, planetname, mainlandname).Find(&b)
				DB.Where("username=? AND planetname = ? AND mainlandname = ?", user.Name, planetname, mainlandname).Find(&e)
				DB.Where("username=? AND planetname = ? AND mainlandname = ?", user.Name, planetname, mainlandname).Find(&f)
				for i = 0; i < 5; i++ {
					c.Allproduct = c.Allproduct + a[i].Production*float64(a[i].Number)
				}
				for i = 0; i < 3; i++ {
					c.Allproduct = c.Allproduct + b[i].Production*float64(b[i].Number)
				}
				for i = 0; i < 4; i++ {
					c.Allproduct = c.Allproduct - e[i].Pollution*float64(e[i].Number) + f[i].Production*float64(f[i].Number)
				}
				c.Allcleanrate = c.Allproduct / c.Successclean
				if c.Allcleanrate >= 1 {
					var g Mainland
					switch c.Name {
					case "西伦瑞亚":
						{
							DB.Where("name = ? AND username = ? AND planetname = ?", "米尔勒拉", user.Name, planetname).Find(&g)
							g.Status = true
							DB.Save(g)
						}
					case "米尔勒拉":
						{
							DB.Where("name = ? AND username = ? AND planetname = ?", "乌兰宇蒂", user.Name, planetname).Find(&g)
							g.Status = true
							DB.Save(g)
						}
					case "乌兰宇蒂":
						{
							DB.Where("name = ? AND username = ? AND planetname = ?", "碦拉玛干", user.Name, planetname).Find(&g)
							g.Status = true
							DB.Save(g)
						}
					case "碦拉玛干":
						{
							DB.Where("name = ? AND username = ? AND planetname = ?", "云格雷诺", user.Name, planetname).Find(&g)
							g.Status = true
							DB.Save(g)
						}
					}
				}
			}
		}
	}()
	// 阻塞主 goroutine，使程序持续运行
	select {}

}

/*// 初始大陆(badbuilding)
func (mainland *Mainland) Unlockmainland(stageimages [3]string) {
	mainland.Status = true
	var i int
	//初始化大陆上的动物，植物，图片
	for i = 0; i < 4; i++ {
		mainland.Goodbuildings = Goodbuildings[i*4 : i*4+4]
		mainland.Animinals = Animals[i*4 : i*4+4]
	}
	for i = 0; i < 3; i++ {
		mainland.Plants = Plants[i*3 : i*3+3]
		//mainland.Stageimage = common.Statusimage[i*3 : i*3+3]
	}
	mainland.Dayclean_Allproduct(mainland.Badbuildings, mainland.Goodbuildings)

}

// 计算大陆每日净产量+大陆总产量+大陆总净化率
func (mainland *Mainland) Dayclean_Allproduct(badbuildings []Badbuilding, goodbuildings []Goodbuilding) {
	var i = 0
	for i = 0; i < 4; i++ {
		mainland.Report.Pollution = badbuildings[i].Pollution*float64(badbuildings[i].Number) + mainland.Report.Pollution
		mainland.Report.Cleanliness = goodbuildings[i].Production*float64(goodbuildings[i].Number) + mainland.Report.Cleanliness
	}
	mainland.Report.Dayclean = mainland.Report.Pollution + mainland.Report.Cleanliness
	mainland.Allproduct = mainland.Allproduct + mainland.Report.Dayclean
	mainland.Allcleanrate = mainland.Allproduct / mainland.Successclean
	var radis = math.Pow(10, 1)
	mainland.Allcleanrate = math.Round(mainland.Allcleanrate*radis) / radis
	//改变shu
	if mainland.Allcleanrate < 0.5 {
		mainland.Image = mainland.Stageimage[0]
	} else if mainland.Allcleanrate >= 0.5 || mainland.Allcleanrate < 1 {
		mainland.Image = mainland.Stageimage[1]
	} else {
		mainland.Image = mainland.Stageimage[2]
	}
}

// 放置回收动物
func (mainland *Mainland) Placeandback_animal(n int, m int) { //n为动物的类型标号,m为放置的数量
	mainland.Animinals[n].Number = mainland.Animinals[n].Number + m
}

// 放置回收植物
func (mainland *Mainland) Placeandback_plant(n int, m int) { //n为植物的类型标号,m为放置的数量
	mainland.Plants[n].Number = mainland.Plants[n].Number + m
}

// 放置回收环保的建筑物
func (mainland *Mainland) Placeandback_building(n int, m int) { //n为建筑物类型，m为放置的数量
	mainland.Goodbuildings[n].Number = mainland.Goodbuildings[n].Number + m
}

// 拆除污染的建筑物
func (mainland *Mainland) Back_badbuilding(n int, m int) {
	mainland.Badbuildings[n].Number = mainland.Badbuildings[n].Number + m
}
*/
