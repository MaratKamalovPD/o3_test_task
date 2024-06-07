package inmemory_test

import (
	"context"
	"testing"

	"github.com/MaratKamalovPD/o3_test_task/internal/models"
	postinmemrepo "github.com/MaratKamalovPD/o3_test_task/internal/pkg/post/repository/inmemory"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	repo := postinmemrepo.NewInMemoryPostRepository()

	samplePost := &models.Post{Title: "New Test Post"}

	newPost, err := repo.CreatePost(context.Background(), samplePost)
	assert.NoError(t, err)
	assert.NotNil(t, newPost)
	assert.Equal(t, "New Test Post", newPost.Title)
}

func TestGetPosts(t *testing.T) {
	repo := postinmemrepo.NewInMemoryPostRepository()

	samplePost1 := &models.Post{ID: 1, Title: "Test Post 1"}
	samplePost2 := &models.Post{ID: 2, Title: "Test Post 2"}

	_, _ = repo.CreatePost(context.Background(), samplePost1)
	_, _ = repo.CreatePost(context.Background(), samplePost2)

	posts, err := repo.GetPosts(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 2, len(posts))
	assert.Contains(t, posts, samplePost1)
	assert.Contains(t, posts, samplePost2)
}

func TestGetPost(t *testing.T) {
	repo := postinmemrepo.NewInMemoryPostRepository()

	samplePost := &models.Post{ID: 1, Title: "Test Post"}

	_, _ = repo.CreatePost(context.Background(), samplePost)

	post, err := repo.GetPost(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, post)
	assert.Equal(t, "Test Post", post.Title)
}

func TestDisableComments(t *testing.T) {
	repo := postinmemrepo.NewInMemoryPostRepository()

	samplePost := &models.Post{ID: 1, Title: "Test Post"}

	_, _ = repo.CreatePost(context.Background(), samplePost)

	err := repo.DisableComments(context.Background(), 1)
	assert.NoError(t, err)

	assert.True(t, samplePost.AreCommentsDisabled)
}
