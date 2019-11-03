package workcloud

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// get to the url of args with cookie.
func (workCloud *WorkCloud) get(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("User-Agent", workCloud.agent)

	res, err := workCloud.client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)

	return string(b), nil
}

// post to the url of args with cookie.
func (workCloud *WorkCloud) post(uri, token string, values url.Values) (string, error) {
	req, err := http.NewRequest("POST", uri, strings.NewReader(values.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Add("User-Agent", workCloud.agent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if len(token) != 0 {
		req.Header.Add("X-CSRF-TOKEN", token)
	}

	res, err := workCloud.client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)

	return string(b), nil
}
