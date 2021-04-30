package main

import (
	portainertemplates "github.com/datahearth/portainer-templates"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "portainer-templates",
	Short: "portainer-templates is a program that allows user to manage Portainer templates",
	Run: func(cmd *cobra.Command, args []string) {
		portainertemplates.Start()
	},
}

func main() {
	rootCmd.Execute()
}
