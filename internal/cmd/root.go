package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go-couchbase",
	Short: "go-couchbase",
	Long:  `This is go couchbase service`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
