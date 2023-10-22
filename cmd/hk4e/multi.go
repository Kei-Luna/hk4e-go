package main

import (
	"context"

	"hk4e/multi/app"

	"github.com/spf13/cobra"
)

func MultiCmd() *cobra.Command {
	var cfg string
	app.APPVERSION = VERSION
	c := &cobra.Command{
		Use:   "multi",
		Short: "multi server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.Run(context.Background(), cfg)
		},
	}
	c.Flags().StringVar(&cfg, "config", "application.toml", "config file")
	return c
}
