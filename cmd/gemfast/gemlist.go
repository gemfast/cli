package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/jedib0t/go-pretty/v6/table"
)

var gemListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all ruby gems stored on the gemfast server",
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		listGems(output)
	},
}

func init() {
	gemListCmd.PersistentFlags().String("output", "table", "Output format. Options are table or json")
	gemCmd.AddCommand(gemListCmd)
}

func listGems(output string) {
	gems, _ := client.ListGems()
	if output == "table" {
		t := table.NewWriter()
  t.SetOutputMirror(os.Stdout)
  t.AppendHeader(table.Row{"name", "version", "platform"})
  for _, gem := range gems {
	   for _, g := range gem {
	    t.AppendRows([]table.Row{{g.Name, g.Version, g.Platform}})
	  }
	  t.AppendSeparator()
	}
  t.Render()
	}
}