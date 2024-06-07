package inmemory_test

import (
	"context"
	"testing"

	"github.com/MaratKamalovPD/o3_test_task/internal/models"
	commentinmemrepo "github.com/MaratKamalovPD/o3_test_task/internal/pkg/comment/repository/inmemory"
	"github.com/stretchr/testify/assert"
)

func TestCreateComment(t *testing.T) {
	repo := commentinmemrepo.NewInMemoryCommentRepository()

	sampleComment := &models.Comment{Content: "Пример комментария"}

	newComment, err := repo.CreateComment(context.Background(), sampleComment)
	assert.NoError(t, err)
	assert.NotNil(t, newComment)
	assert.Equal(t, "Пример комментария", newComment.Content)
}

func TestGetCommentsByPost(t *testing.T) {
	repo := commentinmemrepo.NewInMemoryCommentRepository()

	// Создаем несколько комментариев
	sampleComment1 := &models.Comment{Content: "Комментарий 1", PostID: 1}
	sampleComment2 := &models.Comment{Content: "Комментарий 2", PostID: 1}
	sampleComment3 := &models.Comment{Content: "Комментарий 3", PostID: 2}

	_, _ = repo.CreateComment(context.Background(), sampleComment1)
	_, _ = repo.CreateComment(context.Background(), sampleComment2)
	_, _ = repo.CreateComment(context.Background(), sampleComment3)

	comments, err := repo.GetCommentsByPost(context.Background(), 1, 10, 0)
	assert.NoError(t, err)
	assert.Len(t, comments, 2)
	assert.Contains(t, comments, sampleComment1)
	assert.Contains(t, comments, sampleComment2)
}
