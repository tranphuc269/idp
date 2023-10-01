package webreq

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type HTTP struct {
	client http.Client
}

func (h HTTP) JSON(
	method string,
	url string,
	headers map[string]string,
	body string,
	v interface{},
) error {
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	for key, val := range headers {
		req.Header.Add(key, val)
	}

	res, err := h.client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode >= http.StatusBadRequest {
		return errors.New(res.Status)
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, v)
	if err != nil {
		return err
	}
	return nil
}

func NewHTTP(client http.Client) HTTP {
	return HTTP{
		client: client,
	}
}

func NewHTTPClient() http.Client {
	return http.Client{}
}
