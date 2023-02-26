package cmd

import (
	"fmt"
	"os"

	"github.com/gemfast/cli/gemfast"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gemfast",
	Short: "The gemfast server command line interface",
}

var client *gemfast.Client

func init() {
	client = gemfast.NewClient("http://localhost:8881", &gemfast.LocalAuth{JWTToken: os.Getenv("GEMFAST_AUTH_TOKEN")})
	if config := gemfast.ReadInConfig(); config != nil {
		client = gemfast.NewClient("http://localhost:8881", &gemfast.LocalAuth{JWTToken: config.Token})
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "'%s'", err)
		os.Exit(1)
	}
}