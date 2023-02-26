package gemfast

import (
  "encoding/json"
  "fmt"
  "os"
  "io/ioutil"
  "os/user"
)

type Config struct {
	Token string `json:"token"`
}

func ReadInConfig() (*Config){
	usr, _ := user.Current()
	configPath := fmt.Sprintf("%s/.gemfast/config.json", usr.HomeDir)
	if _, err := os.Stat(configPath); err == nil {
		configFile, err := os.Open(configPath)
		if err != nil {
			panic(err)
		}
		byteValue, _ := ioutil.ReadAll(configFile)
		var cfg Config
		json.Unmarshal(byteValue, &cfg)
		return &cfg
	}
	return &Config{}
}