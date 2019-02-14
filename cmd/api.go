package cmd

import (
	"github.com/CryptoLarry/ssm/conf"
	"github.com/CryptoLarry/ssm/database/postgres"
	"github.com/CryptoLarry/ssm/server"
	"github.com/CryptoLarry/ssm/ssmapi"
	cli "github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
	rootCmd.AddCommand(apiCmd)
}

var (
	apiCmd = &cli.Command{
		Use:   "api",
		Short: "Start API",
		Long:  `Start API`,
		Run: func(cmd *cli.Command, args []string) { // Initialize the databse

			var thingStore ssmapi.SSMStore
			var err error

			thingStore, err = postgres.New()
			if err != nil {
				logger.Fatalw("Database Error", "error", err)
			}

			// Create the server
			s, err := server.New(thingStore)
			if err != nil {
				logger.Fatalw("Could not create server",
					"error", err,
				)
			}
			err = s.ListenAndServe()
			if err != nil {
				logger.Fatalw("Could not start server",
					"error", err,
				)
			}

			if err != nil {
				logger.Fatalw("Database Error", "error", err)
			}
			<-conf.Stop.Chan() // Wait until StopChan
			conf.Stop.Wait()   // Wait until everyone cleans up
			zap.L().Sync()     // Flush the logger

		},
	}
)
