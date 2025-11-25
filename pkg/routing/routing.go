package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/oshie15/go-blog.git/internal/providers/routes"
)

func Init() {
	router = gin.Default()
}

func GetRouter() *gin.Engine {
	return router
}


func RegisterRoutes(){
	routes.RegisterRoutes(GetRouter())
}