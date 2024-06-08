package post

import (
	"context"

	"github.com/MaratKamalovPD/o3_test_task/internal/models"
	"github.com/MaratKamalovPD/o3_test_task/internal/models/args"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go

type PostRepository interface {
	CreatePost(ctx context.Context, post *models.Post) (*models.Post, error)
	GetPosts(ctx context.Context) ([]*models.Post, error)
	GetPost(ctx context.Context, id uint) (*models.Post, error)
	DisableComments(ctx context.Context, postID uint) (bool, error)
	PostExists(ctx context.Context, postID uint) (bool, error)
}

type PostUsecases interface {
	GetPosts(ctx context.Context) (any, error)
	GetPost(ctx context.Context, args args.PostArgs) (any, error)
	CreatePost(ctx context.Context, args args.CreatePostArgs) (any, error)
	DisableComments(ctx context.Context, args args.DisableCommentsArgs) (any, error)

	PostExists(ctx context.Context, postID uint) error
	CommentsDisabled(ctx context.Context, postID uint) error
}
