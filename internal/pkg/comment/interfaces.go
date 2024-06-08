package comment

import (
	"context"

	"github.com/MaratKamalovPD/o3_test_task/internal/models"
	"github.com/MaratKamalovPD/o3_test_task/internal/models/args"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go

type CommentRepository interface {
	CreateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error)
	GetCommentsByPost(ctx context.Context, postID, num, offset uint) ([]*models.Comment, error)
}

type CommentUsecases interface {
	CreateComment(ctx context.Context, args args.CreateCommentArgs) (any, error)
	GetCommentsByPost(ctx context.Context, args args.GetCommentsArgs) (any, error)
}
