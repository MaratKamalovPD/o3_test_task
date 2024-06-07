package post

import (
	"context"

	"github.com/MaratKamalovPD/o3_test_task/internal/models"
	"github.com/MaratKamalovPD/o3_test_task/internal/models/args"
)

type PostRepository interface {
	CreatePost(ctx context.Context, post *models.Post) (*models.Post, error)
	GetPosts(ctx context.Context) ([]*models.Post, error)
	GetPost(ctx context.Context, id uint) (*models.Post, error)
	DisableComments(ctx context.Context, postID uint) error
}

type PostUsecases interface {
	GetPosts(ctx context.Context) (any, error)
	GetPost(ctx context.Context, args args.PostArgs) (any, error)
	CreatePost(ctx context.Context, args args.CreatePostArgs) (any, error)
	DisableComments(ctx context.Context, args args.DisableCommentsArgs) (any, error)
}
