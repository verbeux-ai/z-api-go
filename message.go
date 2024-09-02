package z_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type TextMessageRequest struct {
	Phone         string  `json:"phone,omitempty"`
	Message       string  `json:"message,omitempty"`
	DelayMessage  float32 `json:"delayMessage,omitempty"`
	DelayTyping   float32 `json:"delayTyping,omitempty"`
	EditMessageId string  `json:"editMessageId,omitempty"`
}

type TextMessageResponse struct {
	ZaapId    string `json:"zaapId"`
	MessageId string `json:"messageId"`
	Id        string `json:"id"`
}

func (s *Client) SendTextMessage(d *TextMessageRequest) (*TextMessageResponse, error) {
	resp, err := s.request(d, http.MethodPost, fmt.Sprintf(textMessageEndpoint, s.instance, s.token))
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
		return nil, fmt.Errorf("failed to send text message with code %d: %w", resp.StatusCode, bodyErr)
	}

	var toReturn TextMessageResponse
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return &toReturn, nil
}
