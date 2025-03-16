package cmd

import (
	"github.com/Useurmind/solar-controller/pkg/config"
	"github.com/Useurmind/solar-controller/pkg/log"
	"github.com/Useurmind/solar-controller/pkg/server"
	"github.com/spf13/cobra"
)

type ServerParams struct {
	configFilePath string
}

var serverParams = ServerParams{}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the solar controller server",
	RunE: executeServer,
}

func initServer(parent *cobra.Command) {
	serverCmd.Flags().StringVarP(&serverParams.configFilePath, "config", "c", "/etc/solar-controller.yaml", "The path to the config file for the solar controller server")

	parent.AddCommand(serverCmd)
}

func executeServer(cmd *cobra.Command, args []string) error {
	cfg, err := config.GetConfig(serverParams.configFilePath)
	log.CheckError(err, "reading config %s", serverParams.configFilePath)

	server := server.NewServer(cfg)
	err = server.Run()
	log.CheckError(err, "running server")

	return nil
}
