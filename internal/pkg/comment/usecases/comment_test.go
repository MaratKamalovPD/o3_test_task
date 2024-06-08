package usecases_test

import (
	"context"
	"testing"

	"github.com/MaratKamalovPD/o3_test_task/internal/models"
	"github.com/MaratKamalovPD/o3_test_task/internal/models/args"
	commentmocks "github.com/MaratKamalovPD/o3_test_task/internal/pkg/comment/mocks"
	"github.com/MaratKamalovPD/o3_test_task/internal/pkg/comment/usecases"
	postmocks "github.com/MaratKamalovPD/o3_test_task/internal/pkg/post/mocks"
	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestGetCommentsByPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := commentmocks.NewMockCommentRepository(ctrl)
	mockPostUsecases := postmocks.NewMockPostUsecases(ctrl)

	defer ctrl.Finish()

	mockRepo.EXPECT().GetCommentsByPost(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Comment{}, nil).Times(1)
	mockPostUsecases.EXPECT().PostExists(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	uc := usecases.NewCommentUsecases(mockRepo, mockPostUsecases)

	args := args.GetCommentsArgs{
		PostID: 1,
		Limit:  10,
		Offset: 0,
	}

	_, err := uc.GetCommentsByPost(context.Background(), args)
	assert.NoError(t, err)

}

func TestCreateComment(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := commentmocks.NewMockCommentRepository(ctrl)
	mockPostUsecases := postmocks.NewMockPostUsecases(ctrl)

	defer ctrl.Finish()

	mockRepo.EXPECT().CreateComment(gomock.Any(), gomock.Any()).Return(&models.Comment{}, nil).Times(1)
	mockPostUsecases.EXPECT().PostExists(gomock.Any(), gomock.Any()).Return(nil).Times(1)
	mockPostUsecases.EXPECT().CommentsDisabled(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	uc := usecases.NewCommentUsecases(mockRepo, mockPostUsecases)

	args := args.CreateCommentArgs{
		PostID:          1,
		UserID:          1,
		Content:         "Valid content",
		ParentCommentID: nil,
	}

	_, err := uc.CreateComment(context.Background(), args)
	assert.NoError(t, err)

}
