package trongrid

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func MyPost(strURL string, req any, rsp any) error {
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}

	payload := bytes.NewReader(b)

	r, err := http.NewRequest("POST", strURL, payload)
	if err != nil {
		return err
	}

	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(r)
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
