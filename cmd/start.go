package main

import (
	"github.com/gin-gonic/gin"
	"github.com/socketbasehq/socketbase/pkg/db"
	"github.com/socketbasehq/socketbase/pkg/modules"
	"github.com/socketbasehq/socketbase/pkg/modules/auth"
	"github.com/socketbasehq/socketbase/pkg/server"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		db.ConnectDB()

		auth.CreateAdminUser()

		app := fx.New(
			modules.Modules,
			fx.Provide(server.NewServer),
			fx.Invoke(func(app *gin.Engine) {
				go app.Run(":3000")
			}),
		)

		app.Run()
	},
}
