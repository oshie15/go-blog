package routes

import (
	"github.com/gin-gonic/gin"
	homeRoutes "github.com/oshie15/go-blog.git/internal/modules/home/routes"
)

func RegisterRoutes(router *gin.Engine) {

	homeRoutes.Routes(router)
}