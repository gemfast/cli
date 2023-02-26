package gemfast

import (
	"encoding/json"
	"io/ioutil"
)

type ListGems [][]struct {
	Name     string `json:"Name"`
	Version  string `json:"Version"`
	Platform string `json:"Platform"`
}

func (c *Client) ListGems() (ListGems, error) {
	res, err := c.get("/gems")
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, err
	}
	body, _ := ioutil.ReadAll(res.Body)
	var gemList ListGems
	json.Unmarshal(body, &gemList)
	return gemList, nil
}