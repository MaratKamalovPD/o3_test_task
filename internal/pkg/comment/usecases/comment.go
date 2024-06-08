package usecases

import (
	"context"
	"fmt"

	"github.com/MaratKamalovPD/o3_test_task/internal/models"
	"github.com/MaratKamalovPD/o3_test_task/internal/models/args"
	commentrepo "github.com/MaratKamalovPD/o3_test_task/internal/pkg/comment"
	postinterface "github.com/MaratKamalovPD/o3_test_task/internal/pkg/post"
	utils "github.com/MaratKamalovPD/o3_test_task/internal/pkg/utils"
)

type CommentUsecases struct {
	storage      commentrepo.CommentRepository
	usecasesPost postinterface.PostUsecases
}

func NewCommentUsecases(storage commentrepo.CommentRepository, usecasesPost postinterface.PostUsecases) *CommentUsecases {
	return &CommentUsecases{
		storage:      storage,
		usecasesPost: usecasesPost,
	}
}

func (uc *CommentUsecases) GetCommentsByPost(ctx context.Context, args args.GetCommentsArgs) (any, error) {
	if err := utils.ValidateID(args.PostID); err != nil {
		return nil, err
	}

	if err := utils.ValidatePaginationArgs(args.Limit, args.Offset); err != nil {
		return nil, err
	}

	err := uc.usecasesPost.PostExists(ctx, uint(args.PostID))

	if err != nil {

		return nil, err
	}

	comments, err := uc.storage.GetCommentsByPost(ctx, uint(args.PostID), uint(args.Limit), uint(args.Offset))
	if err != nil {
		return nil, fmt.Errorf("failed to get comments: %w", err)
	}

	return comments, nil
}

func (uc *CommentUsecases) CreateComment(ctx context.Context, args args.CreateCommentArgs) (any, error) {
	if err := utils.ValidateID(args.PostID, args.UserID); err != nil {
		return nil, err
	}

	if err := utils.ValidateCommentLength(args.Content); err != nil {

		return nil, err
	}

	comment := &models.Comment{
		PostID:          uint(args.PostID),
		ParentCommentID: args.ParentCommentID,
		UserID:          uint(args.UserID),
		Content:         args.Content,
	}

	err := uc.usecasesPost.PostExists(ctx, uint(args.PostID))

	if err != nil {

		return nil, err
	}

	err = uc.usecasesPost.CommentsDisabled(ctx, uint(args.PostID))

	if err != nil {

		return nil, err
	}

	savedComment, err := uc.storage.CreateComment(ctx, comment)
	if err != nil {

		return nil, fmt.Errorf("something went wrong while creating comment, err=%w", err)
	}

	return savedComment, nil
}
