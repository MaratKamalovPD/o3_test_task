package postgres

import (
	"context"

	"github.com/MaratKamalovPD/o3_test_task/internal/models"
	postrepo "github.com/MaratKamalovPD/o3_test_task/internal/pkg/post"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresPostRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresPostRepository(pool *pgxpool.Pool) postrepo.PostRepository {
	return &postgresPostRepository{
		pool: pool,
	}
}

func (r *postgresPostRepository) createPost(ctx context.Context, tx pgx.Tx, post *models.Post) (*models.Post, error) {

	SQLInsertPost := `INSERT INTO public.post(
		user_id, title, content, are_comments_disabled)
		VALUES ($1, $2, $3, $4)
		RETURNING id, user_id, title, content, created_time, are_comments_disabled;`

	postLine := tx.QueryRow(ctx, SQLInsertPost, post.UserID, post.Title, post.Content,
		post.AreCommentsDisabled)

	if err := postLine.Scan(&post.ID, &post.UserID, &post.Title,
		&post.Content, &post.CreatedAt, &post.AreCommentsDisabled); err != nil {

		return nil, err
	}

	return post, nil
}

func (r *postgresPostRepository) CreatePost(ctx context.Context, post *models.Post) (*models.Post, error) {

	var returningPost *models.Post

	err := pgx.BeginFunc(ctx, r.pool, func(tx pgx.Tx) error {
		returningCommentInner, err := r.createPost(ctx, tx, post)
		returningPost = returningCommentInner

		return err
	})

	if err != nil {

		return nil, err
	}

	return returningPost, nil
}

func (r *postgresPostRepository) getPosts(ctx context.Context, tx pgx.Tx) ([]*models.Post, error) {

	SQLSelectPosts := `SELECT id, user_id, title, content, created_time, are_comments_disabled
	FROM public.post;`

	rows, err := tx.Query(ctx, SQLSelectPosts)

	if err != nil {

		return nil, err
	}

	defer rows.Close()

	var postList []*models.Post

	for rows.Next() {
		returningPost := models.Post{}

		if err := rows.Scan(&returningPost.ID, &returningPost.UserID, &returningPost.Title,
			&returningPost.Content, &returningPost.CreatedAt, &returningPost.AreCommentsDisabled); err != nil {
			return nil, err
		}

		postList = append(postList, &returningPost)
	}

	if err := rows.Err(); err != nil {

		return nil, err
	}

	return postList, nil
}

func (r *postgresPostRepository) GetPosts(ctx context.Context) ([]*models.Post, error) {

	var postList []*models.Post

	err := pgx.BeginFunc(ctx, r.pool, func(tx pgx.Tx) error {
		postListInner, err := r.getPosts(ctx, tx)
		postList = postListInner

		return err
	})

	if err != nil {

		return nil, err
	}

	if postList == nil {
		postList = []*models.Post{}
	}

	return postList, nil
}

func (r *postgresPostRepository) getPost(ctx context.Context, tx pgx.Tx, id uint) (*models.Post, error) {

	SQLSelectPost := `SELECT id, user_id, title, content, created_time, are_comments_disabled
	FROM public.post
	WHERE id = $1;`

	postLine := tx.QueryRow(ctx, SQLSelectPost, id)

	returningPost := models.Post{}

	if err := postLine.Scan(&returningPost.ID, &returningPost.UserID, &returningPost.Title,
		&returningPost.Content, &returningPost.CreatedAt, &returningPost.AreCommentsDisabled); err != nil {

		return nil, err
	}

	return &returningPost, nil
}

func (r *postgresPostRepository) GetPost(ctx context.Context, id uint) (*models.Post, error) {
	var returningPost *models.Post

	err := pgx.BeginFunc(ctx, r.pool, func(tx pgx.Tx) error {
		returningCommentInner, err := r.getPost(ctx, tx, id)
		returningPost = returningCommentInner

		return err
	})

	if err != nil {

		return nil, err
	}

	return returningPost, nil
}

func (r *postgresPostRepository) disableComments(ctx context.Context, tx pgx.Tx, postID uint) (bool, error) {

	SQLInsertPost := `UPDATE public.post
	SET are_comments_disabled = CASE 
							   WHEN are_comments_disabled = TRUE THEN FALSE
							   ELSE TRUE
							 END
	WHERE id = $1
	RETURNING are_comments_disabled;`

	boolLine := tx.QueryRow(ctx, SQLInsertPost, postID)

	var areCommentsDisabled bool

	if err := boolLine.Scan(&areCommentsDisabled); err != nil {

		return false, err
	}

	return areCommentsDisabled, nil
}

func (r *postgresPostRepository) DisableComments(ctx context.Context, postID uint) (bool, error) {

	var areCommentsDisabled bool

	err := pgx.BeginFunc(ctx, r.pool, func(tx pgx.Tx) error {
		areCommentsDisabledInner, err := r.disableComments(ctx, tx, postID)
		areCommentsDisabled = areCommentsDisabledInner

		return err
	})

	if err != nil {

		return false, err
	}

	return areCommentsDisabled, nil
}

func (r *postgresPostRepository) postExists(ctx context.Context, tx pgx.Tx, postID uint) (bool, error) {

	SQLPostExists := `SELECT EXISTS(SELECT 1 FROM public.post WHERE id=$1 );`

	userLine := tx.QueryRow(ctx, SQLPostExists, postID)

	var exists bool

	if err := userLine.Scan(&exists); err != nil {

		return false, err
	}

	return exists, nil
}

func (r *postgresPostRepository) PostExists(ctx context.Context, postID uint) (bool, error) {
	var exists bool

	err := pgx.BeginFunc(ctx, r.pool, func(tx pgx.Tx) error {
		userExists, err := r.postExists(ctx, tx, postID)
		exists = userExists

		return err
	})

	if err != nil {

		return false, err
	}

	return exists, nil

}
