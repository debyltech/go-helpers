package helper

import (
	"fmt"
	"net/http"
)

func Get(uri string, authName string, authValue string, queries map[string]string) (*http.Response, error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
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
