package inmemory

import (
	"context"
	"sync"

	"github.com/MaratKamalovPD/o3_test_task/internal/models"
	postrepo "github.com/MaratKamalovPD/o3_test_task/internal/pkg/post"
)

type inMemoryPostRepository struct {
	posts       map[uint]*models.Post
	postCounter uint
	mu          sync.RWMutex
}

func NewInMemoryPostRepository() postrepo.PostRepository {
	return &inMemoryPostRepository{
		posts:       make(map[uint]*models.Post),
		postCounter: 0,
		mu:          sync.RWMutex{},
	}
}

func (r *inMemoryPostRepository) CreatePost(ctx context.Context, post *models.Post) (*models.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.postCounter++
	post.ID = r.postCounter

	r.posts[post.ID] = post

	return post, nil
}

func (r *inMemoryPostRepository) GetPosts(ctx context.Context) ([]*models.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	posts := make([]*models.Post, 0, len(r.posts))
	for _, post := range r.posts {
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *inMemoryPostRepository) GetPost(ctx context.Context, id uint) (*models.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	post := r.posts[id]

	return post, nil
}

func (r *inMemoryPostRepository) DisableComments(ctx context.Context, postID uint) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	post := r.posts[postID]
	if post.AreCommentsDisabled {
		post.AreCommentsDisabled = false
	} else {
		post.AreCommentsDisabled = true
	}

	return post.AreCommentsDisabled, nil
}

func (r *inMemoryPostRepository) PostExists(ctx context.Context, postID uint) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, exists := r.posts[postID]

	return exists, nil
}
