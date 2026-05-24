package github

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) GetListContent(
	ctx context.Context,
	owner string,
	repo string,
	repoPath string,
	ref string,
) ([]ContentItem, error) {
	req, err := c.newRequest(
		ctx,
		http.MethodGet,
		makeContentPath(owner, repo, repoPath, ref),
	)
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

func (c *Client) GetFileContent(
	ctx context.Context,
	owner string,
	repo string,
	repoPath string,
	ref string,
) ([]byte, error) {
	req, err := c.newRequest(
		ctx,
		http.MethodGet,
		makeContentPath(owner, repo, repoPath, ref),
	)
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

	var file FileContent
	if err := json.Unmarshal(raw, &file); err != nil {
		return nil, fmt.Errorf("decode response body: %w", err)
	}

	if file.Type != "file" {
		return nil, fmt.Errorf("path %q is not a file, got type %q", repoPath, file.Type)
	}

	if file.Encoding != "base64" {
		return nil, fmt.Errorf("unsupported encoding %q", file.Encoding)
	}

	clean := strings.ReplaceAll(file.Content, "\n", "")

	decoded, err := base64.StdEncoding.DecodeString(clean)
	if err != nil {
		return nil, fmt.Errorf("decode base64 content: %w", err)
	}

	return decoded, nil
}

func makeContentPath(owner string, repo string, repoPath string, ref string) string {
	path := fmt.Sprintf(
		"/repos/%s/%s/contents/%s",
		url.PathEscape(owner),
		url.PathEscape(repo),
		repoPath,
	)

	if ref != "" {
		path += "?ref=" + url.QueryEscape(ref)
	}

	return path
}
