package postgres

import (
	"context"

	"github.com/MaratKamalovPD/o3_test_task/internal/models"
	commentrepo "github.com/MaratKamalovPD/o3_test_task/internal/pkg/comment"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresCommentRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresCommentRepository(pool *pgxpool.Pool) commentrepo.CommentRepository {
	return &postgresCommentRepository{
		pool: pool,
	}
}

func (r *postgresCommentRepository) createComment(ctx context.Context, tx pgx.Tx, comment *models.Comment) (*models.Comment, error) {

	SQLInsertComment := `INSERT INTO public.comment(
		post_id, parent_comment_id, user_id, "content")
		VALUES ($1, $2, $3, $4)
		RETURNING id, post_id, parent_comment_id, user_id, "content", created_time;`

	userLine := tx.QueryRow(ctx, SQLInsertComment, comment.PostID, comment.ParentCommentID, comment.UserID,
		comment.Content)

	if err := userLine.Scan(&comment.ID, &comment.PostID, &comment.ParentCommentID,
		&comment.UserID, &comment.Content, &comment.CreatedAt); err != nil {

		return nil, err
	}

	return comment, nil
}

func (r *postgresCommentRepository) CreateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error) {

	var returningComment *models.Comment

	err := pgx.BeginFunc(ctx, r.pool, func(tx pgx.Tx) error {
		returningCommentInner, err := r.createComment(ctx, tx, comment)
		returningComment = returningCommentInner

		return err
	})

	if err != nil {

		return nil, err
	}

	return returningComment, nil
}

func (r *postgresCommentRepository) getCommentsByPost(ctx context.Context, tx pgx.Tx, postID, limit, offset uint) ([]*models.Comment, error) {

	SQLGetFavouritesByUserID := `SELECT id, post_id, parent_comment_id, user_id, content, created_time
	FROM public.comment
	WHERE post_id = $1
	LIMIT $2 
	OFFSET $3;`

	rows, err := tx.Query(ctx, SQLGetFavouritesByUserID, postID, limit, offset)

	if err != nil {

		return nil, err
	}

	defer rows.Close()

	var commentList []*models.Comment

	for rows.Next() {
		returningComment := models.Comment{}

		if err := rows.Scan(&returningComment.ID, &returningComment.PostID, &returningComment.ParentCommentID,
			&returningComment.UserID, &returningComment.Content, &returningComment.CreatedAt); err != nil {
			return nil, err
		}

		commentList = append(commentList, &returningComment)
	}

	if err := rows.Err(); err != nil {

		return nil, err
	}

	return commentList, nil
}

func (r *postgresCommentRepository) GetCommentsByPost(ctx context.Context, postID, limit, offset uint) ([]*models.Comment, error) {

	var commentList []*models.Comment

	err := pgx.BeginFunc(ctx, r.pool, func(tx pgx.Tx) error {
		commentListInner, err := r.getCommentsByPost(ctx, tx, postID, limit, offset)
		commentList = commentListInner

		return err
	})

	if err != nil {

		return nil, err
	}

	if commentList == nil {
		commentList = []*models.Comment{}
	}

	return commentList, nil
}
