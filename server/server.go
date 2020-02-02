package server

import (
	"github.com/joyching/golang-practice/config"
)

func Init() {
	config := config.GetConfig()

	router := NewRouter()
	router.Run(":" + config.GetString("server.port"))
}
