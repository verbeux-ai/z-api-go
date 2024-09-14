package z_api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type ChatMetadataResponse struct {
	CommunityId         interface{} `json:"communityId"`
	IsGroupAnnouncement bool        `json:"isGroupAnnouncement"`
	IsGroup             bool        `json:"isGroup"`
	Name                string      `json:"name"`
	Phone               string      `json:"phone"`
	Unread              string      `json:"unread"`
	LastMessageTime     string      `json:"lastMessageTime"`
	IsMuted             string      `json:"isMuted"`
	IsMarkedSpam        string      `json:"isMarkedSpam"`
	AgentId             interface{} `json:"agentId"`
	Tags                []string    `json:"tags"`
	Archived            string      `json:"archived"`
	Pinned              string      `json:"pinned"`
	MuteEndTime         interface{} `json:"muteEndTime"`
	ProfileThumbnail    string      `json:"profileThumbnail"`
	EphemeralExpiration int         `json:"ephemeralExpiration"`
	About               string      `json:"about"`
}

func (s *Client) GetChat(ctx context.Context, phone string) (*ChatMetadataResponse, error) {
	resp, err := s.request(ctx, nil, http.MethodGet, fmt.Sprintf(chatEndpoint, s.instance, s.token, phone))
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

	var toReturn ChatMetadataResponse
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return &toReturn, nil
}

func (s *Client) SetChatTag(ctx context.Context, tagID string, phone string) error {
	resp, err := s.request(ctx, nil, http.MethodPut, fmt.Sprintf(chatAddTagEndpoint, s.instance, s.token, phone, tagID))
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
		return fmt.Errorf("failed to set chat tag with code %d: %w", resp.StatusCode, bodyErr)
	}

	return nil
}
