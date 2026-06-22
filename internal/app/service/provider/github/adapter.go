package pgithub

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/internal/app/entity"
	"github.com/merulis/shuffle/internal/app/service"
)

type Adapter struct {
	client *Client
	owner  string
	repo   string
}

func NewAdapter(client *Client, owner, repo string) service.Provider {
	return &Adapter{
		client,
		owner,
		repo,
	}
}

func (a *Adapter) List(ctx context.Context, loc entity.Locator) ([]entity.Artifact, error) {
	items, err := a.client.GetListContent(
		ctx,
		a.owner,
		a.repo,
		loc.Path,
		loc.Ref,
	)
	if err != nil {
		return nil, fmt.Errorf("get list content: %w", err)
	}

	artifacts := make([]entity.Artifact, 0, len(items))
	for _, item := range items {
		artifacts = append(artifacts, toArtifact(item))
	}

	return artifacts, nil
}

func (a *Adapter) Read(ctx context.Context, loc entity.Locator) ([]byte, error) {
	data, err := a.client.GetFileContent(
		ctx,
		a.owner,
		a.repo,
		loc.Path,
		loc.Ref,
	)
	if err != nil {
		return nil, fmt.Errorf("get file content: %w", err)
	}

	return data, nil
}
