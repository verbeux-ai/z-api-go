package z_api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type WebhookDeliveryRequest struct {
	Value string `json:"phone,omitempty"`
}

type WebhookReceivedRequest struct {
	Value string `json:"phone,omitempty"`
}

func (s *Client) SetWebhookDelivery(d *WebhookDeliveryRequest) error {
	resp, err := s.request(d, http.MethodPut, fmt.Sprintf(webhookDeliveryEndpoint, s.instance, s.token))
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
		return fmt.Errorf("failed to set delivery webhook with code %d: %w", resp.StatusCode, bodyErr)
	}

	return nil
}

func (s *Client) SetWebhookReceived(d *WebhookReceivedRequest) error {
	resp, err := s.request(d, http.MethodPut, fmt.Sprintf(webhookReceivedEndpoint, s.instance, s.token))
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
		return fmt.Errorf("failed to set received webhook with code %d: %w", resp.StatusCode, bodyErr)
	}

	return nil
}
