package main

import (
	"fmt"
	"github.com/GatilovSergey/test-NBA/pkg/server"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "run the test-NBA web server",
	}
	fPort := cmd.Flags().String("host", ":8080", "host the server will listen on")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		app, err := server.NewApp(*fPort)
		if err != nil {
			return fmt.Errorf("cannot create server: %w", err)
		}
		return app.Run()
	}
	Root.AddCommand(cmd)
}
