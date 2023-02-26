package cmd

import (
	"fmt"
	"os"
	"os/user"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var loginCmd = &cobra.Command{
	Use:   "login [username]",
	Short: "Authenticate with the gemfast server",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		login(args[0])
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(username string) {
	fmt.Print("password: ")
	pass, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println("'" + string(pass) + "'")
	if err != nil {
		panic(err)
	}
	res, err := client.Login(username, string(pass))
	if err != nil {
		fmt.Println("\nlogin failed")
		os.Exit(1)
	} else {
		fmt.Println("\nlogin succeeded")
	}
	usr, _ := user.Current()
	gemfdir := fmt.Sprintf("%s/.gemfast", usr.HomeDir)
	err = os.MkdirAll(gemfdir, os.ModePerm)
	if err != nil {
		panic(err)
	}
	path := fmt.Sprintf("%s/config.json", gemfdir)
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(string(res))
	if err != nil {
		panic(err)
	}
}
