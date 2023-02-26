package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"github.com/jedib0t/go-pretty/v6/table"
)

var createtokenCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new API token",
	Run: func(cmd *cobra.Command, args []string) {
		gems, _ := client.ListGems()

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
	},
}

func init() {
	tokenCmd.AddCommand(createtokenCmd)
}

func writeGemCredentials(dir string, body []byte) {
	fname := fmt.Sprintf("%s/credentials", dir)
	if _, err := os.Stat(fname); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
		f, err := os.Create(fname)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		j := map[string]string{}
		json.Unmarshal(body, &j)
		delete(j, "code")
		delete(j, "expire")
		j[":gemfast"] = fmt.Sprintf("Bearer %s", j["token"])
		delete(j, "token")
		data, err := yaml.Marshal(&j)
		_, err = f.WriteString(string(data))
		if err != nil {
			panic(err)
		}
	} else {
		data := make(map[interface{}]interface{})

		yfile, err := ioutil.ReadFile(fname)
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(yfile, &data)
		if err != nil {
			panic(err)
		}
		j := map[string]string{}
		json.Unmarshal(body, &j)
		data[":gemfast"] = fmt.Sprintf("Bearer %s", j["token"])
		out, err := yaml.Marshal(&data)
		err = ioutil.WriteFile(fname, out, 0)
		if err != nil {
			panic(err)
		}
	}
}

// writeGemCredentials(fmt.Sprintf("%s/.gem", usr.HomeDir), body)