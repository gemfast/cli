package gemfast

import (
	"encoding/json"
	"io/ioutil"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *Client) Login(username string, password string) ([]byte, error){
	login := Login{Username: username, Password: password}
	jsonData, _ := json.Marshal(login)
	res, err := c.post("/login", jsonData)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, err
	}
	body, _ := ioutil.ReadAll(res.Body)
	return body, nil
}