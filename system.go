package z_api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type StatusResponse struct {
	Connected           bool   `json:"connected"`
	Session             bool   `json:"session"`
	Created             int64  `json:"created"`
	Error               string `json:"error"`
	SmartphoneConnected bool   `json:"smartphoneConnected"`
}

type QrCodeImageResponse struct {
	Value string `json:"value"`
}

func (s *Client) Status(ctx context.Context) (*StatusResponse, error) {
	resp, err := s.request(ctx, nil, http.MethodGet, fmt.Sprintf(statusEndpoint, s.instance, s.token))
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
		return nil, fmt.Errorf("failed to get status with code %d: %w", resp.StatusCode, bodyErr)
	}

	var toReturn StatusResponse
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return &toReturn, nil
}

func (s *Client) QrCodeImage(ctx context.Context) (*QrCodeImageResponse, error) {
	resp, err := s.request(ctx, nil, http.MethodGet, fmt.Sprintf(qrCodeImageEndpoint, s.instance, s.token))
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

	var toReturn QrCodeImageResponse
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return &toReturn, nil
}

func (s *Client) Disconnect(ctx context.Context) error {
	resp, err := s.request(ctx, nil, http.MethodGet, fmt.Sprintf(disconnectEndpoint, s.instance, s.token))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		bodyErr := errors.New(string(body))
		return fmt.Errorf("failed to get qr code with code %d: %w", resp.StatusCode, bodyErr)
	}

	return nil
}
