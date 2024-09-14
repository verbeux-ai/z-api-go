package z_api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type TagsResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Color int    `json:"color"`
}

func (s *Client) GetTags(ctx context.Context) ([]TagsResponse, error) {
	resp, err := s.request(ctx, nil, http.MethodGet, fmt.Sprintf(tagsEndpoint, s.instance, s.token))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		bodyErr := errors.New(string(body))
		return nil, fmt.Errorf("failed to get qr code with code %d: %w", resp.StatusCode, bodyErr)
	}

	var toReturn []TagsResponse
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return toReturn, nil
}
