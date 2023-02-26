package cmd

import (
	"github.com/spf13/cobra"
)

var gemCmd = &cobra.Command{
	Use:   "gem",
	Short: "Commands for working with ruby gems",
}

func init() {
	rootCmd.AddCommand(gemCmd)
}