package inmemory

import (
	"context"
	"sync"

	"github.com/MaratKamalovPD/o3_test_task/internal/models"
	commentrepo "github.com/MaratKamalovPD/o3_test_task/internal/pkg/comment"
)

type inMemoryCommentRepository struct {
	comments       map[uint]*models.Comment
	commentCounter uint
	mu             sync.RWMutex
}

func NewInMemoryCommentRepository() commentrepo.CommentRepository {
	return &inMemoryCommentRepository{
		comments:       make(map[uint]*models.Comment),
		commentCounter: 0,
		mu:             sync.RWMutex{},
	}
}

func (r *inMemoryCommentRepository) CreateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.commentCounter++

	comment.ID = r.commentCounter
	r.comments[comment.ID] = comment

	return comment, nil
}

func (r *inMemoryCommentRepository) filterByPostID(postID uint, comment *models.Comment) bool {

	return comment.PostID == postID
}

func (r *inMemoryCommentRepository) GetCommentsByPost(ctx context.Context, postID, limit, offset uint) ([]*models.Comment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	comments := make([]*models.Comment, 0, limit)
	var count uint = 0
	for _, comment := range r.comments {
		if r.filterByPostID(postID, comment) {
			if count >= offset {
				comments = append(comments, comment)
				if uint(len(comments)) == limit {
					break
				}
			}
			count++
		}
	}

	return comments, nil
}
