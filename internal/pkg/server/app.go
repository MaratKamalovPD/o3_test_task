package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

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
	serverrepo "github.com/MaratKamalovPD/o3_test_task/internal/pkg/server/repository"
	"github.com/graphql-go/handler"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	Address = ":8080"
)

type Server struct {
	server *http.Server
}

func (srv *Server) Run() error {

	cfg := config.ReadConfig()

	var (
		postRepo    postinterface.PostRepository
		commentRepo commentinterface.CommentRepository
	)

	switch cfg.Server.StorageType {
	case "inmemory":
		postRepo = postinmemorystorage.NewInMemoryPostRepository()
		commentRepo = commentinmemorystorage.NewInMemoryCommentRepository()
	case "postgres":
		connPool, err := pgxpool.NewWithConfig(context.Background(), serverrepo.PGXPoolConfig(cfg.Server.DBConnectingURL))
		if err != nil {
			log.Fatal("something went wrong while creating connection pool, err=", err)
		}

		err = serverrepo.MakeMigrations(cfg.Server.DBMigrationsFolder, cfg.Server.DBConnectingURL)
		if err != nil {
			log.Fatal("something went wrong while making migrations, err=", err)
		}

		postRepo = postpostgresstorage.NewPostgresPostRepository(connPool)
		commentRepo = commentpostgresstorage.NewPostgresCommentRepository(connPool)
	default:
		return fmt.Errorf("unknown repository type: %s", cfg.Server.StorageType)
	}

	postUsecases := postusecases.NewPostUsecases(postRepo)
	commentUsecases := commentusecases.NewCommentUsecases(commentRepo, postUsecases)

	postSchema, err := postdelivery.NewPostSchema(postUsecases)
	if err != nil {
		log.Fatal("something went wrong while creating Post schema: ", err)
	}

	commentSchema, err := commentdelivery.NewCommentSchema(commentUsecases)
	if err != nil {
		log.Fatal("something went wrong while creating Comment schema: ", err)
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
		Handler: mux,
		Addr:    Address,
	}

	log.Println("server is running on port ", Address)

	return srv.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {

	return s.server.Shutdown(ctx)
}
