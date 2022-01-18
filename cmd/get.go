/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	getCmdFlagVerbose *bool
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get specified item by key",
	Long: `Get specified item by item key. For example:
$ stash get foo
`,
	Run: func(cmd *cobra.Command, args []string) {
		k := []byte(args[0])
		item, err := Provider().Get(k)
		if err != nil {
			log.Println(err)
			return
		}
		if *getCmdFlagVerbose {
			fmt.Printf("Key => str: [%s] hex: [%v]\n", string(k), k)
			fmt.Printf("Value => str: [%s] hex: [%v]\n", string(item.Value), item.Value)
			fmt.Printf("Meta => %v", item.Meta)
		} else {
			os.Stdout.Write(item.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmdFlagVerbose = getCmd.Flags().BoolP("verbose", "v", false, "output item in verbose mode")
}
