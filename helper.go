package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Get(uri string, authName string, authValue string, queries map[string]string) (*http.Response, error) {
	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("%s %s", authName, authValue))

	if len(queries) > 0 {
		q := request.URL.Query()

		for key, value := range queries {
			q.Add(key, value)
		}

		request.URL.RawQuery = q.Encode()
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func Post(uri string, authName string, authValue string, data any) (*http.Response, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(jsonBytes)
	request, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("%s %s", authName, authValue))

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
