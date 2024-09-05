package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"go-couchbase/internal/database"
	"go-couchbase/internal/repository"
)

var runServer = &cobra.Command{
	Use:   "server",
	Short: "run server",
	Long:  `This subcommand start the server`,
	Run:   server,
}

func init() {
	RootCmd.AddCommand(runServer)
}

func server(cmd *cobra.Command, args []string) {
	// connect database couchbase
	dbCouchbase, closeCluster := database.InitializeDatabaseCouchbase()
	defer closeCluster()

	// register repo
	emailLogRepository := repository.NewEmailLogRepository(dbCouchbase)

	/*
			// insert
			newDataLog, _ := emailLogRepository.CreateEmailLog(context.Background(), &entity.EmailLog{
				Sources:   "eraspace",
				Recipient: "reoshby@gmail.com",
				Status:    "pending",
			})
		logrus.Info(newDataLog)
	*/

	/*
		emailLog, _ := emailLogRepository.GetEmailLogByID(context.Background(), 1725577474930742000)
		marshal, _ := json.Marshal(emailLog)
		logrus.Info(string(marshal))
	*/

	recipient, _ := emailLogRepository.GetEmailLogByRecipient(context.Background(), "reoshby1@gmail.com")
	marshal, _ := json.Marshal(&recipient)
	fmt.Println(string(marshal))
}
