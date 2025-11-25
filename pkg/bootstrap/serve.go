package bootstrap

import (
	"github.com/oshie15/go-blog.git/pkg/config"
	"github.com/oshie15/go-blog.git/pkg/routing"
)

func Serve() {

	config.Set()

	routing.Init()
	

	routing.RegisterRoutes()

	routing.Serve()
}