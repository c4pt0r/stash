/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"io"
	"log"
	"time"

	"github.com/c4pt0r/stash/datasource"
	"github.com/spf13/cobra"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "put item, put [key]",
	Long: `A longer description that spans multiple lines and likely contains examples
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		k := []byte(args[0])
		v, err := io.ReadAll(cmd.InOrStdin())
		if err != nil {
			log.Fatal(err)
		}
		meta := datasource.ItemMeta{
			Tp:        datasource.TypeBytes,
			Len:       len(v),
			CreateAt:  time.Now(),
			Tags:      []string{},
			Namespace: datasource.DefaultNS,
		}
		Provider().Put(k, v, meta)
		Provider().Sync()
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
}
