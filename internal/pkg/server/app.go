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
	commentusecases "github.com/MaratKamalovPD/o3_test_task/internal/pkg/comment/usecases"
	"github.com/MaratKamalovPD/o3_test_task/internal/pkg/config"
	postinterface "github.com/MaratKamalovPD/o3_test_task/internal/pkg/post"
	postdelivery "github.com/MaratKamalovPD/o3_test_task/internal/pkg/post/delivery/schema"
	postinmemorystorage "github.com/MaratKamalovPD/o3_test_task/internal/pkg/post/repository/inmemory"
	postusecases "github.com/MaratKamalovPD/o3_test_task/internal/pkg/post/usecases"
	"github.com/graphql-go/handler"
)

const (
	Timeout    = time.Second * 15
	Address    = ":8080"
	tickerTime = 10 * time.Minute
)

// Server is a server that handles GraphQL schema
type Server struct {
	server *http.Server
}

// // NewServer creates a new GraphQL server
// func NewServer(schema *graphql.Schema) *Server {

// 	return &Server{
// 		server: server,
// 	}
// }

// Run starts HTTP server on the given port
func (srv *Server) Run() error {

	ctx := context.Background()

	cfg := config.ReadConfig()

	postRepo, err := createPostRepository(ctx, cfg)
	if err != nil {
		log.Fatal("Something went wrong while creating Post repository: ", err)
	}

	commentRepo, err := createCommentRepository(ctx, cfg)
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

func createPostRepository(ctx context.Context, cfg *config.Config) (postinterface.PostRepository, error) {
	switch cfg.Server.StorageType {
	case "inmemory":
		return postinmemorystorage.NewInMemoryPostRepository(), nil
	default:
		return nil, fmt.Errorf("unknown repository type: %s", cfg.Server.StorageType)
	}
}

func createCommentRepository(ctx context.Context, cfg *config.Config) (commentinterface.CommentRepository, error) {
	switch cfg.Server.StorageType {
	case "inmemory":
		return commentinmemorystorage.NewInMemoryCommentRepository(), nil
	default:
		return nil, fmt.Errorf("unknown repository type: %s", cfg.Server.StorageType)
	}
}
