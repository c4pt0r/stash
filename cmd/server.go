/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/c4pt0r/log"
	"github.com/spf13/cobra"
)

var (
	serverAddr *string
)

func init() {
	rootCmd.AddCommand(serverCmd)
	serverAddr = serverCmd.Flags().StringP("addr", "l", "0.0.0.0:6334", "listen addr")
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{Use: "server",
	Short: "create a new stash server listening in given addr",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := serve(*serverAddr)
		if err != nil {
			log.F(err)
		}
	},
}
