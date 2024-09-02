package z_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (s *Client) request(reqBody any, method, endpoint string) (*http.Response, error) {
	var bodyReader io.Reader
	if reqBody != nil {
		marshalledBody, err := json.Marshal(reqBody)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(marshalledBody)
	}

	url := fmt.Sprintf("%s/%s", s.baseUrl, endpoint)

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Client-Token", s.secret)

	return s.httpClient.Do(req)
}
