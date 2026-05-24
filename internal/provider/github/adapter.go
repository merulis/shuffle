package github

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/internal/domain"
)

type Adapter struct {
	client *Client
	owner  string
	repo   string
}

func NewAdapter(client *Client, owner, repo string) *Adapter {
	return &Adapter{
		client,
		owner,
		repo,
	}
}

func (s *Adapter) List(ctx context.Context, loc domain.Locator) ([]domain.Artifact, error) {
	items, err := s.client.GetListContent(
		ctx,
		s.owner,
		s.repo,
		loc.Path,
		loc.Ref,
	)
	if err != nil {
		return nil, fmt.Errorf("get list content: %w", err)
	}

	artifacts := make([]domain.Artifact, 0, len(items))
	for _, item := range items {
		artifacts = append(artifacts, toArtifact(item))
	}

	return artifacts, nil
}

func (s *Adapter) Read(ctx context.Context, loc domain.Locator) ([]byte, error) {
	data, err := s.client.GetFileContent(
		ctx,
		s.owner,
		s.repo,
		loc.Path,
		loc.Ref,
	)
	if err != nil {
		return nil, fmt.Errorf("get file content: %w", err)
	}

	return data, nil
}
