package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) ListContent(
	ctx context.Context,
	owner string,
	repo string,
	repoPath string,
	ref string,
) ([]ContentItem, error) {
	path := fmt.Sprintf(
		"/repos/%s/%s/contents/%s",
		url.PathEscape(owner),
		url.PathEscape(repo),
		repoPath,
	)

	if ref != "" {
		path += "?ref=" + url.QueryEscape(ref)
	}

	req, err := c.newRequest(ctx, http.MethodGet, path)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github api returned status: %d: %s", resp.StatusCode, string(raw))
	}

	var items []ContentItem
	if err := json.Unmarshal(raw, &items); err != nil {
		return nil, fmt.Errorf("decode contents  response: %w", err)
	}

	return items, nil
}
