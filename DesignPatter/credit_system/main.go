package main

import routers "github.com/starichat/notes/DesignPatter/credit_system/http/routes"

func main() {
	router := routers.InitRouter()
	router.Run(":8080")
}
