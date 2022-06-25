package trongrid

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func MyGet(strURL string, rsp any) error {
	req, err := http.NewRequest("GET", strURL, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, rsp)
	return err
}
