package main

import (
	"com.higginslee/IMSystem/router"
	"com.higginslee/IMSystem/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	r := router.Router()
	r.Run(":8001")
}
