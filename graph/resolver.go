package graph

import (
	"github.com/SwArch-2025-1-2A/backend/graph/model"
	"github.com/SwArch-2025-1-2A/backend/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
	Repo  repository.Queries
}

// func (r *queryResolver) Groups(ctx context.Context) ([]*model.Group, error) {
// 	dbGroups, err := r.Repo.GetGroups(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("issue fetching groups: %w", err)
// 	}
// 	return dbGroups, nil
// }
