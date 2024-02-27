package main

// @title saveplanet
// @version 1.0
// @description  救星API
// @contact.name lyd
// @contact.email 2771730573@qq.com
// @host 这里写服务的host
// @BasePath api
import (
	"miniproject/app/core/gorm"
	"miniproject/app/routers"
)

func main() {
	var db = gorm.Linktodatabase()
	gorm.Migrate(db)
	a := routers.Routerinit()
	a.Run(":8080")
}
