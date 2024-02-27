package controller

import (
	"github.com/gin-gonic/gin"
	"miniproject/app/model"
	"net/http"
)

// 获得大陆信息 /mainland/xilun GET
func Getxilun(c *gin.Context) {
	var a model.Mainland
	model.Mainlandenergy(model.User1, model.Planetname, "西伦瑞亚")
	result := model.DB.Where("username = ? AND planetname = ? AND name = 西伦瑞亚", model.User1.Name, model.Planetname).Find(a)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, "查询大陆信息出现问题")
		return
	}
	c.JSON(200, gin.H{
		"大陆净化率": a.Allcleanrate,
		"大陆总产能": a.Allproduct,
		"大陆名称":  "西伦瑞亚",
		"大陆气候":  a.Climate,
		"大陆地形":  a.Terrian,
		"其他信息":  a.Others,
	})
}

// 获得大陆信息 /mainland/mier GET
func Getmier(c *gin.Context) {
	var a model.Mainland
	model.Mainlandenergy(model.User1, model.Planetname, "米尔勒拉")
	result := model.DB.Where("username = ? AND planetname = ? AND name = 米尔勒拉", model.User1.Name, model.Planetname).Find(a)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, "查询大陆信息出现问题")
		return
	}
	c.JSON(200, gin.H{
		"大陆净化率": a.Allcleanrate,
		"大陆总产能": a.Allproduct,
		"大陆名称":  "米尔勒拉",
		"大陆气候":  a.Climate,
		"大陆地形":  a.Terrian,
		"其他信息":  a.Others,
	})
}

// 获得大陆信息 /mainland/wulan GET
func Getwulan(c *gin.Context) {
	var a model.Mainland
	model.Mainlandenergy(model.User1, model.Planetname, "乌兰宇蒂")
	result := model.DB.Where("username = ? AND planetname = ? AND name = 乌兰宇蒂", model.User1.Name, model.Planetname).Find(a)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, "查询大陆信息出现问题")
		return
	}
	c.JSON(200, gin.H{
		"大陆净化率": a.Allcleanrate,
		"大陆总产能": a.Allproduct,
		"大陆名称":  "乌兰宇蒂",
		"大陆气候":  a.Climate,
		"大陆地形":  a.Terrian,
		"其他信息":  a.Others,
	})
}

// 获得大陆信息 /mainland/kala GET
func Getkala(c *gin.Context) {
	var a model.Mainland
	model.Mainlandenergy(model.User1, model.Planetname, "喀拉玛干")
	result := model.DB.Where("username = ? AND planetname = ? AND name = 喀拉玛干", model.User1.Name, model.Planetname).Find(a)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, "查询大陆信息出现问题")
		return
	}
	c.JSON(200, gin.H{
		"大陆净化率": a.Allcleanrate,
		"大陆总产能": a.Allproduct,
		"大陆名称":  "喀拉玛干",
		"大陆气候":  a.Climate,
		"大陆地形":  a.Terrian,
		"其他信息":  a.Others,
	})
}

// 获得大陆信息 /mainland/yunluo GET
func Getyunluo(c *gin.Context) {
	var a model.Mainland
	model.Mainlandenergy(model.User1, model.Planetname, "云落雷诺")
	result := model.DB.Where("username = ? AND planetname = ? AND name = 云落雷诺", model.User1.Name, model.Planetname).Find(a)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, "查询大陆信息出现问题")
		return
	}
	c.JSON(200, gin.H{
		"大陆净化率": a.Allcleanrate,
		"大陆总产能": a.Allproduct,
		"大陆名称":  "云落雷诺",
		"大陆气候":  a.Climate,
		"大陆地形":  a.Terrian,
		"其他信息":  a.Others,
	})
}

// 产能报告 /mainland/report POST
func Getreport(c *gin.Context) {
	var a string
	a = c.PostForm("mainlandname")
	var report model.Report
	model.Createreport(model.User1, model.Planetname, a)
	model.Reportenergy(model.User1, model.Planetname, a)
	result := model.DB.Where("username = ? AND mainlandname =? AND planetname = ?", model.User1.Name, a, model.Planetname).Find(report)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "读取报告信息失败"})
		return
	}
	c.JSON(200, gin.H{
		"每日增产": report.Dayclean,
		"每日污染": report.Pollution,
		"每日净产": report.Cleanliness,
	})

}
