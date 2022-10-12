package httpx

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func PostWithJson(target string, devices interface{}) ([]byte, error) {

	d, err := json.Marshal(devices)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(d)
	reqs, err := http.NewRequest("POST", target, reader)
	if err != nil {
		return nil, err
	}

	reqs.ContentLength = int64(reader.Len())
	reqs.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(reqs)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return body, err
}

func Get(target string) ([]byte, error) {
	r, err := http.Get(target)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	return body, err
}
