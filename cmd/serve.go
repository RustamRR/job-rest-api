package cmd

import (
	"github.com/RustamRR/job-rest-api/internal/app/apiserver"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Запуск echo сервера",
	Long:  `Запуск echo сервера`,
	Run: func(cmd *cobra.Command, args []string) {
		apiserver.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
