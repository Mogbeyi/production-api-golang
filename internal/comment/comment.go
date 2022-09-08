package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented = errors.New("not implemented")
)

type Comment struct {
	ID string
	Slug string
	Body string
	Author string
}

// Defines all methods that our service
// needs to operate
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
		return Comment{}, ErrFetchingComment
	}

	return cmt, nil
}