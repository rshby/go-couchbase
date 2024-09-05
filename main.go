package main

import (
	"go-couchbase/config"
	"go-couchbase/internal/cmd"
	"go-couchbase/internal/logger"
)

func init() {
	logger.SetupLogger()
	config.LoadConfig()
}

func main() {
	cmd.Execute()
}
