package cmd

import (
	"github.com/spf13/cobra"
	"log/slog"
)

var rootCmd = &cobra.Command{
	Use:   "tcp-practice",
	Short: "a demo set of tcp",
	Long:  "a demo set of tcp",
	Run:   rootFunc,
}

func rootFunc(cmd *cobra.Command, args []string) {
	slog.Info("tcp-practice init OK")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
