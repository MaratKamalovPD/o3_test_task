package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/MaratKamalovPD/o3_test_task/internal/models"
	"github.com/MaratKamalovPD/o3_test_task/internal/models/args"
	postrepo "github.com/MaratKamalovPD/o3_test_task/internal/pkg/post"
	utils "github.com/MaratKamalovPD/o3_test_task/internal/pkg/utils"
)

// var (
// 	ErrPostNotFound     = fmt.Errorf("post not found")
// 	ErrCommentNotFound  = fmt.Errorf("comment not found")
// 	ErrCommentsDisabled = fmt.Errorf("comments are disabled")
// 	ErrNotAuthor        = fmt.Errorf("only author can disable comments")
// )

type PostUsecases struct {
	storage postrepo.PostRepository
}

func NewPostUsecases(storage postrepo.PostRepository) *PostUsecases {
	return &PostUsecases{storage: storage}
}

func (uc *PostUsecases) GetPosts(ctx context.Context) (any, error) {
	posts, err := uc.storage.GetPosts(ctx)
	if err != nil {

		return nil, fmt.Errorf("something went wrong while getting posts, err=%w", err)
	}

	return posts, nil
}

func (uc *PostUsecases) GetPost(ctx context.Context, args args.PostArgs) (any, error) {
	if err := utils.ValidateID(args.ID); err != nil {

		return nil, err
	}

	// if err := uc.postExists(ctx, args.ID); err != nil {

	// 	return nil, err
	// }

	post, err := uc.storage.GetPost(ctx, uint(args.ID))
	if err != nil {

		return nil, fmt.Errorf("something went wrong while getting single post, err=%w", err)
	}

	return post, nil
}

func (uc *PostUsecases) CreatePost(ctx context.Context, args args.CreatePostArgs) (any, error) {
	if err := utils.ValidateID(args.UserID); err != nil {

		return nil, err
	}

	post := &models.Post{
		Title:     args.Title,
		Content:   args.Content,
		UserID:    uint(args.UserID),
		CreatedAt: time.Now().UTC(),
	}

	savedPost, err := uc.storage.CreatePost(ctx, post)
	if err != nil {

		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	return savedPost, nil
}

func (uc *PostUsecases) DisableComments(ctx context.Context, args args.DisableCommentsArgs) (any, error) {
	if err := utils.ValidateID(args.PostID, args.UserID); err != nil {

		return nil, err
	}

	// if err := uc.postExists(ctx, args.PostID); err != nil {

	// 	return nil, err
	// }

	// if err := uc.authorizeAuthor(ctx, args.PostID, args.UserID); err != nil {

	// 	return nil, err
	// }

	err := uc.storage.DisableComments(ctx, uint(args.PostID))
	if err != nil {

		return nil, fmt.Errorf("failed to disable comments: %w", err)
	}

	return true, nil
}
