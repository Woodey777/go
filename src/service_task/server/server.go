package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	// db
	engine *gin.Engine
}

func createGood(c *gin.Context) {

}

func updateGood(c *gin.Context) {

}

func removeGood(c *gin.Context) {

}

func listGoods(c *gin.Context) {

}

func reprioritizeGood(c *gin.Context) {

}

func NewServer() *Server {
	server := &Server{}
	engine := gin.Default()
	
	engine.GET("/goods/list", listGoods)
	engine.POST("/good/create", createGood)
	engine.PATCH("/good/update", updateGood)
	engine.PATCH("/good/reprioritize", listGoods)
	engine.DELETE("/good/remove", removeGood)
	
	server.engine = engine
	return server
}