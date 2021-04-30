package portainertemplates

import (
	"github.com/datahearth/portainer-templates/pkg/db"
	"github.com/datahearth/portainer-templates/pkg/server"
	"github.com/sirupsen/logrus"
)

func Start() {
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
}
