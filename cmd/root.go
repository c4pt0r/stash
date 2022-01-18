/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"
	"sync"

	"github.com/c4pt0r/stash/datasource"
	"github.com/spf13/cobra"
)

var (
	_global_ds   datasource.DataSource
	_global_once sync.Once
)

func dsn() string {
	return "local://localds.json"
}

func Provider() datasource.DataSource {
	_global_once.Do(func() {
		var err error
		_global_ds, err = datasource.Init(dsn())
		if err != nil {
			log.Fatal(err)
		}
	})
	return _global_ds
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stash",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
