package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

var (
	baseURL, _ = url.Parse("https://api.openai.com")
)

type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
	config     Config
}

func NewClient(httpClient *http.Client, config Config) *Client {
	return &Client{
		httpClient: httpClient,
		baseURL:    baseURL,
		config:     config,
	}
}

func (c *Client) doAPI(ctx context.Context, method string, pathname string, params any, respBody any) error {
	u, err := c.baseURL.Parse(pathname)
	if err != nil {
		return errors.Wrap(err, "failed to parse URL")
	}

	var reqBody io.Reader

	switch method {
	case http.MethodPost:
		if params == nil {
			break
		}

		b, err := json.Marshal(params)
		if err != nil {
			return errors.Wrap(err, "failed to encode params")
		}

		reqBody = bytes.NewBuffer(b)

	default:
		return errors.New("unsupported method")
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), reqBody)
	if err != nil {
		return errors.Wrap(err, "failed to prepare request")
	}

	req.Header.Set("Authorization", "Bearer "+c.config.APIKey)

	switch method {
	case http.MethodPost:
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		var errResp ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return errors.Wrap(err, "failed to decode error response body")
		}

		return &errResp.Error
	}

	if respBody != nil {
		if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
			return errors.Wrap(err, "failed to decode respnose body")
		}
	}

	return nil
}
