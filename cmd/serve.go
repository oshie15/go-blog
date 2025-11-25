package cmd

import (
	"github.com/oshie15/go-blog.git/pkg/bootstrap"
	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
  Use:   "serve",
  Short: "Serve app on the dev server",
  Long:  `Application will be served on host and port defined in the config.yml`,
  Run: func(cmd *cobra.Command, args []string) {
    serve()
  },
}


func serve() {
	bootstrap.Serve()
}

