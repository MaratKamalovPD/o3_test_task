package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	commentinterface "github.com/MaratKamalovPD/o3_test_task/internal/pkg/comment"
	commentdelivery "github.com/MaratKamalovPD/o3_test_task/internal/pkg/comment/delivery/schema"
	commentinmemorystorage "github.com/MaratKamalovPD/o3_test_task/internal/pkg/comment/repository/inmemory"
	commentpostgresstorage "github.com/MaratKamalovPD/o3_test_task/internal/pkg/comment/repository/postgres"
	commentusecases "github.com/MaratKamalovPD/o3_test_task/internal/pkg/comment/usecases"
	"github.com/MaratKamalovPD/o3_test_task/internal/pkg/config"
	postinterface "github.com/MaratKamalovPD/o3_test_task/internal/pkg/post"
	postdelivery "github.com/MaratKamalovPD/o3_test_task/internal/pkg/post/delivery/schema"
	postinmemorystorage "github.com/MaratKamalovPD/o3_test_task/internal/pkg/post/repository/inmemory"
	postpostgresstorage "github.com/MaratKamalovPD/o3_test_task/internal/pkg/post/repository/postgres"
	postusecases "github.com/MaratKamalovPD/o3_test_task/internal/pkg/post/usecases"
	pgxpoolconfig "github.com/MaratKamalovPD/o3_test_task/internal/pkg/server/repository"
	"github.com/graphql-go/handler"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	Timeout    = time.Second * 15
	Address    = ":8080"
	tickerTime = 10 * time.Minute
)

type Server struct {
	server *http.Server
}

func (srv *Server) Run() error {

	cfg := config.ReadConfig()

	connPool, err := pgxpool.NewWithConfig(context.Background(), pgxpoolconfig.PGXPoolConfig())
	if err != nil {
		log.Fatal("Error while creating connection to the database")
	}

	postRepo, err := createPostRepository(cfg, connPool)
	if err != nil {
		log.Fatal("Something went wrong while creating Post repository: ", err)
	}

	commentRepo, err := createCommentRepository(cfg, connPool)
	if err != nil {
		log.Fatal("Something went wrong while creating Comment repository: ", err)
	}

	postUsecases := postusecases.NewPostUsecases(postRepo)
	commentUsecases := commentusecases.NewCommentUsecases(commentRepo)

	postSchema, err := postdelivery.NewPostSchema(postUsecases)
	if err != nil {
		log.Fatal("Something went wrong while creating Post schema: ", err)
	}

	commentSchema, err := commentdelivery.NewCommentSchema(commentUsecases)
	if err != nil {
		log.Fatal("Something went wrong while creating Comment schema: ", err)
	}

	postHandler := handler.New(&handler.Config{
		Schema:   &postSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	commentHandler := handler.New(&handler.Config{
		Schema:   &commentSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	mux := http.NewServeMux()

	mux.Handle("/post", postHandler)
	mux.Handle("/comment", commentHandler)

	srv.server = &http.Server{
		Handler:      mux,
		Addr:         Address,
		ReadTimeout:  Timeout,
		WriteTimeout: Timeout,
	}

	return srv.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func createPostRepository(cfg *config.Config, pool *pgxpool.Pool) (postinterface.PostRepository, error) {
	switch cfg.Server.StorageType {
	case "inmemory":
		return postinmemorystorage.NewInMemoryPostRepository(), nil
	case "postgres":
		return postpostgresstorage.NewPostgresPostRepository(pool), nil
	default:
		return nil, fmt.Errorf("unknown repository type: %s", cfg.Server.StorageType)
	}
}

func createCommentRepository(cfg *config.Config, pool *pgxpool.Pool) (commentinterface.CommentRepository, error) {
	switch cfg.Server.StorageType {
	case "inmemory":
		return commentinmemorystorage.NewInMemoryCommentRepository(), nil
	case "postgres":
		return commentpostgresstorage.NewPostgresCommentRepository(pool), nil
	default:
		return nil, fmt.Errorf("unknown repository type: %s", cfg.Server.StorageType)
	}
}
