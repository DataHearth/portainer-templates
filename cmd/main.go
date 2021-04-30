package main

import (
	"os"

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
		level := "info"
		if os.Getenv("LOG_LEVEL") != "" {
			level = os.Getenv("LOG_LEVEL")
		}

		logLevel, err := logrus.ParseLevel(level)
		if err != nil {
			logger.WithError(err).WithField("log-level", level).Errorln("invalid log level. Using default INFO")
			logrus.SetLevel(logrus.InfoLevel)
		} else {
			logrus.SetLevel(logLevel)
		}

		db, err := db.NewDB(logger)
		if err != nil {
			logger.WithError(err).Fatalln("failed to create a database instance")
		}

		port := "4345"
		if os.Getenv("PORT") != "" {
			port = os.Getenv("PORT")
		}

		srv, err := server.NewServer(logger, db, os.Getenv("HOST"), port)
		if err != nil {
			logger.WithError(err).Fatalln("failed to create a server instance")
		}

		srv.RegisterHandlers()
		if err := srv.Start(); err != nil {
			logger.WithError(err).Fatalln("Server error")
		}
	},
}

func main() {
	rootCmd.Execute()
}
