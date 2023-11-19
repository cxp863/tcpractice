package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "tcp-practice",
	Short: "a demo set of tcp",
	Long:  "a demo set of tcp",
	Run:   rootFunc,
}

func rootFunc(cmd *cobra.Command, args []string) {}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

var logs *zap.SugaredLogger

// init logs
func init() {
	if zapLog, err := zap.NewDevelopment(); err != nil {
		log.Fatal("zap log init failed")
	} else {
		logs = zapLog.Sugar()
	}
}

// FlushLog flush logs
func FlushLog() {
	err := logs.Sync()
	if err != nil {
		log.Fatal("zap log sync failed")
	}
}
