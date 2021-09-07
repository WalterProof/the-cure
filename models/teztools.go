package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TezTools interacts with teztools api.
type TezTools interface {
	XTZPrice() (float64, error)
}

type tezTools struct {
}

func newTezTools() *tezTools {
	return &tezTools{}
}

func (tt *tezTools) XTZPrice() (float64, error) {
	responseBytes, err := tt.makeRequest("xtz-price")
	if err != nil {
		return 0, fmt.Errorf("error reading HTTP response body: %w", err)
	}

	var r map[string]interface{}
	if err := json.Unmarshal(responseBytes, &r); err != nil {
		return 0, fmt.Errorf("error deserializing price data: %w", err)
	}

	return r["price"].(float64), nil
}

func (tt *tezTools) makeRequest(endpoint string) ([]byte, error) {
	// TODO mock teztools on local env to avoid spamming
	// TODO basic cache on prod
	req, err := http.NewRequest(http.MethodGet, "https://api.teztools.io/v1/"+endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %w", err)
	}
	return ioutil.ReadAll(res.Body)
}
