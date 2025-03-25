package cmd

import (
	"anonsurvey/client"
	"anonsurvey/config"
	"anonsurvey/internal/adapter"
	"anonsurvey/internal/adapter/controller"
	"anonsurvey/internal/core"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {
	conf := config.Config()

	dbClient := client.MySQL{}
	dbClient.ConnectDB(conf.DB)

	userRepo := adapter.NewUserRepository(&dbClient)
	userSvc := core.NewUserService(userRepo)
	userCtrl := controller.NewUserController(userSvc)

	server := adapter.NewHTTPServer(
		adapter.HTTPServerPort(conf.App.Port),
	)

	routers := []adapter.HTTPRouter{userCtrl}
	server.Start(routers)
}
