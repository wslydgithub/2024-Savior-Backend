package model

import "time"

type Report struct {
	Pollution    float64
	Cleanliness  float64
	Dayclean     float64
	Planetname   string
	Mainlandname string
	Username     string
}

// 创建并存入数据库中
func Createreport(user User, plantname string, mainlandname string) {
	var a Report
	a.Username = user.Name
	a.Planetname = plantname
	a.Mainlandname = mainlandname
	a.Dayclean = 0
	a.Pollution = 0
	a.Cleanliness = 0
	DB.Create(a)
}

// 计算并展现数据
func Reportenergy(user User, planetname string, mainlandname string) {
	var a [5]Animinal
	var b [3]Plant
	var c [4]Badbuilding
	var d [4]Goodbuilding
	var f Report
	ticker := time.NewTicker(24 * time.Hour)
	// 启动一个 goroutine 来执行定时任务
	go func() {
		for {
			select {
			case <-ticker.C:
				var i int
				DB.Where("username=? AND planetname = ? AND mainlandname = ?", user.Name, planetname, mainlandname).Find(&a)
				DB.Where("username=? AND planetname = ? AND mainlandname = ?", user.Name, planetname, mainlandname).Find(&b)
				DB.Where("username=? AND planetname = ? AND mainlandname = ?", user.Name, planetname, mainlandname).Find(&c)
				DB.Where("username=? AND planetname = ? AND mainlandname = ?", user.Name, planetname, mainlandname).Find(&d)
				DB.Where("username=? AND planetname = ? AND mainlandname = ?", user.Name, planetname, mainlandname).Find(&f)
				for i = 0; i < 5; i++ {
					f.Dayclean = f.Dayclean + a[i].Production*float64(a[i].Number)
				}
				for i = 0; i < 3; i++ {
					f.Dayclean = f.Dayclean + b[i].Production*float64(b[i].Number)
				}
				for i = 0; i < 4; i++ {
					f.Dayclean = f.Dayclean + d[i].Production*float64(d[i].Number)
					f.Pollution = f.Pollution + c[i].Pollution*float64(c[i].Number)
				}
				f.Cleanliness = f.Dayclean - f.Pollution
				DB.Save(f)
			}
		}
	}()
	// 阻塞主 goroutine，使程序持续运行
	select {}
}
