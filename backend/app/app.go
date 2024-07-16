package app

import (
	"github.com/codescalersinternships/Secret-noteAPI-SPA-NabilaSherif/models"
	"github.com/gin-gonic/gin"
)

type App struct {
	port    string
	DB     models.GormDB
	router *gin.Engine
}

func NewApp(port string) {
	
}

func NewGinEngine() *gin.Engine {
	return gin.Default()
}

func Run(listenAddress string) error {
	r := NewGinEngine()
	return r.Run(listenAddress)
}
