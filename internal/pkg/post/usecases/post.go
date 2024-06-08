package usecases

import (
	"context"
	"fmt"

	"github.com/MaratKamalovPD/o3_test_task/internal/models"
	"github.com/MaratKamalovPD/o3_test_task/internal/models/args"
	postrepo "github.com/MaratKamalovPD/o3_test_task/internal/pkg/post"
	utils "github.com/MaratKamalovPD/o3_test_task/internal/pkg/utils"
)

var (
	errPostDoNotExist      = fmt.Errorf("post with such an ID don't exist")
	errUserIsNotAnAuthor   = fmt.Errorf("user isn't the aithor of the post")
	errCommentsAreDisabled = fmt.Errorf("comments are disabled")
)

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

	err := uc.PostExists(ctx, uint(args.ID))

	if err != nil {

		return nil, err
	}

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
		Title:   args.Title,
		Content: args.Content,
		UserID:  uint(args.UserID),
	}

	savedPost, err := uc.storage.CreatePost(ctx, post)
	if err != nil {

		return nil, fmt.Errorf("something whent wrong while creating the post, err=%w", err)
	}

	return savedPost, nil
}

func (uc *PostUsecases) DisableComments(ctx context.Context, args args.DisableCommentsArgs) (any, error) {
	if err := utils.ValidateID(args.PostID, args.UserID); err != nil {

		return nil, err
	}

	err := uc.PostExists(ctx, uint(args.PostID))

	if err != nil {

		return nil, err
	}

	post, err := uc.storage.GetPost(ctx, uint(args.PostID))

	if err != nil {

		return nil, fmt.Errorf("something went wrong while getting single post, err=%w", err)
	}

	if post.UserID != uint(args.UserID) {

		return nil, errUserIsNotAnAuthor
	}

	areCommentsDisabled, err := uc.storage.DisableComments(ctx, uint(args.PostID))
	if err != nil {

		return nil, fmt.Errorf("something whent wrong while disabling comments, err=%w", err)
	}

	return areCommentsDisabled, nil
}

func (uc *PostUsecases) PostExists(ctx context.Context, postID uint) error {

	ok, err := uc.storage.PostExists(ctx, postID)
	if err != nil {

		return fmt.Errorf("something whent wrong while checking post existence, err=%w", err)
	}

	if !ok {

		return errPostDoNotExist
	}

	return nil
}

func (uc *PostUsecases) CommentsDisabled(ctx context.Context, postID uint) error {

	post, err := uc.storage.GetPost(ctx, postID)
	if err != nil {

		return fmt.Errorf("something went wrong while getting single post, err=%w", err)
	}

	if post.AreCommentsDisabled {

		return errCommentsAreDisabled
	}

	return nil
}
