package comment

import (
	"context"
	"fmt"
)

type Comment struct {
	ID string
	Slug string
	Body string
	Author string
}

type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Struct all logic will be built upon
type Service struct {
	Store Store
}

// Returns a pointer to a new service
func NewService(store Store) *Service {
	return &Service {
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a comment")
	cmt, err := s.Store.GetComment(ctx, id)

	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}

	return cmt, nil
}