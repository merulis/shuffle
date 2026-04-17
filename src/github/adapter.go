package github

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/src/domain"
)

type Storage struct {
	client *Client
	owner  string
	repo   string
}

func NewStorage(client *Client, owner, repo string) *Storage {
	return &Storage{
		client,
		owner,
		repo,
	}
}

func (s *Storage) List(ctx context.Context, loc domain.Locator) ([]domain.Artifact, error) {
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

func (s *Storage) Read(ctx context.Context, loc domain.Locator) ([]byte, error) {
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
