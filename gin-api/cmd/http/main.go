package http

import (
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	viperExt "gingo/extensions/viper"
)

func runServers() func(cmd *cobra.Command) {
	return func(cmd *cobra.Command) {
		var wg sync.WaitGroup
		wg.Add(2)
		go runAdmin(&wg)
		go runPublic(&wg)
		wg.Wait()
	}
}

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Run Gin Server",
	Run: func(cmd *cobra.Command, args []string) {
		mode, _ := cmd.Flags().GetString("mode")
		if mode == "dev" {
			viperExt.LoadViper("dev", "./config", false)
		} else if mode == "prod" {
			viperExt.LoadViper("prod", "./config", false)
		} else {
			log.Errorln("Unknown mode received!")
			os.Exit(1)
		}

		runServers()(cmd)
	},
}

// RegisterCommandRecursive is to register command
func RegisterCommandRecursive(parent *cobra.Command) {
	parent.AddCommand(httpCmd)
}
