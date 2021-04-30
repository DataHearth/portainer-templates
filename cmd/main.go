package main

import (
	"github.com/datahearth/portainer-templates/pkg/db"
	"github.com/datahearth/portainer-templates/pkg/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "portainer-templates",
	Short: "portainer-templates is a program that allows user to manage Portainer templates",
	Run: func(cmd *cobra.Command, args []string) {
		logger := logrus.StandardLogger()
		db, err := db.NewDB(logger)
		if err != nil {
			logger.WithError(err).Errorln("failed to create a database instance")
		}

		srv, err := server.NewServer(logger, db)
		if err != nil {
			logger.WithError(err).Errorln("failed to create a server instance")
		}

		srv.RegisterHandlers()
		if err := srv.Start(":4345"); err != nil {
			logger.WithError(err).Errorln("failed to create a server instance")
		}
	},
}

func main() {
	rootCmd.Execute()
}
