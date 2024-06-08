package usecases_test

import (
	"context"
	"testing"

	"github.com/MaratKamalovPD/o3_test_task/internal/models"
	"github.com/MaratKamalovPD/o3_test_task/internal/models/args"

	postmocks "github.com/MaratKamalovPD/o3_test_task/internal/pkg/post/mocks"
	"github.com/MaratKamalovPD/o3_test_task/internal/pkg/post/usecases"
	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestGetPosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := postmocks.NewMockPostRepository(ctrl)
	defer ctrl.Finish()

	mockRepo.EXPECT().GetPosts(gomock.Any()).Return([]*models.Post{}, nil).Times(1)

	uc := usecases.NewPostUsecases(mockRepo)

	result, err := uc.GetPosts(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGetPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := postmocks.NewMockPostRepository(ctrl)
	defer ctrl.Finish()

	post := &models.Post{}
	mockRepo.EXPECT().PostExists(gomock.Any(), gomock.Any()).Return(true, nil).Times(1)
	mockRepo.EXPECT().GetPost(gomock.Any(), gomock.Any()).Return(post, nil).Times(1)

	uc := usecases.NewPostUsecases(mockRepo)

	args := args.PostArgs{ID: 1}
	result, err := uc.GetPost(context.Background(), args)
	assert.NoError(t, err)
	assert.Equal(t, post, result)
}

func TestCreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := postmocks.NewMockPostRepository(ctrl)
	defer ctrl.Finish()

	mockRepo.EXPECT().CreatePost(gomock.Any(), gomock.Any()).Return(&models.Post{}, nil).Times(1)

	uc := usecases.NewPostUsecases(mockRepo)

	args := args.CreatePostArgs{UserID: 1, Title: "Test Title", Content: "Test Content"}
	result, err := uc.CreatePost(context.Background(), args)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestDisableComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := postmocks.NewMockPostRepository(ctrl)
	defer ctrl.Finish()

	mockRepo.EXPECT().PostExists(gomock.Any(), gomock.Any()).Return(true, nil).Times(1)
	mockRepo.EXPECT().GetPost(gomock.Any(), gomock.Any()).Return(&models.Post{UserID: 1}, nil).Times(1)
	mockRepo.EXPECT().DisableComments(gomock.Any(), gomock.Any()).Return(true, nil).Times(1)

	uc := usecases.NewPostUsecases(mockRepo)

	args := args.DisableCommentsArgs{PostID: 1, UserID: 1}
	result, err := uc.DisableComments(context.Background(), args)
	assert.NoError(t, err)
	assert.True(t, (result).(bool))
}
