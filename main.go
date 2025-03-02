package main

import (
	"log"
	"qasystem/config"
	"qasystem/router"
)

func main() {
	con := config.GetConfig()
	r := router.SetupRouter()
	if err := r.Run(":" + con.App.Port); err != nil {
		log.Fatalf("%s服务器启动失败", con.App.Name)
	}

}
