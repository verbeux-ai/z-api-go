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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var toReturn []TagsResponse
	if err = json.Unmarshal(body, &toReturn); err != nil {
		return nil, fmt.Errorf("%w: %s", err, string(body))
	}

	return toReturn, nil
}

type CreateTagRequest struct {
	Name string `json:"name"`
}

type CreateTagResponse struct {
	ID string `json:"id"`
}

func (s *Client) CreateTag(ctx context.Context, name string) (*CreateTagResponse, error) {
	req := &CreateTagRequest{Name: name}
	resp, err := s.request(ctx, req, http.MethodPost, fmt.Sprintf(createTagEndpoint, s.instance, s.token))
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var toReturn CreateTagResponse
	if err = json.Unmarshal(body, &toReturn); err != nil {
		return nil, fmt.Errorf("%w: %s", err, string(body))
	}

	return &toReturn, nil
}
