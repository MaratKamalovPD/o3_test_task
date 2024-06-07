package schema

import (
	"log"

	"github.com/MaratKamalovPD/o3_test_task/internal/models/args"
	postinterface "github.com/MaratKamalovPD/o3_test_task/internal/pkg/post"
	"github.com/graphql-go/graphql"
)

func PostsField(postType *graphql.Object, usecases postinterface.PostUsecases) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(postType),
		Description: "Get all posts",
		Resolve: func(p graphql.ResolveParams) (any, error) {
			res, err := usecases.GetPosts(p.Context)
			if err != nil {
				log.Println("Error response:", err)
			}

			return res, err
		},
	}
}

func PostField(postType *graphql.Object, usecases postinterface.PostUsecases) *graphql.Field {
	return &graphql.Field{
		Type:        postType,
		Description: "Get post by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (any, error) {
			id, _ := p.Args["id"].(int)
			res, err := usecases.GetPost(p.Context, args.PostArgs{ID: id})

			if err != nil {
				log.Println("Error response:", err)
			}

			return res, err
		},
	}
}

func CreatePostField(postType *graphql.Object, usecases postinterface.PostUsecases) *graphql.Field {
	return &graphql.Field{
		Type:        postType,
		Description: "Create new post",
		Args: graphql.FieldConfigArgument{
			"title":   &graphql.ArgumentConfig{Type: graphql.String},
			"content": &graphql.ArgumentConfig{Type: graphql.String},
			"userId":  &graphql.ArgumentConfig{Type: graphql.Int},
		},
		Resolve: func(p graphql.ResolveParams) (any, error) {
			title, _ := p.Args["title"].(string)
			content, _ := p.Args["content"].(string)
			userId, _ := p.Args["userId"].(int)
			res, err := usecases.CreatePost(p.Context, args.CreatePostArgs{
				Title:   title,
				Content: content,
				UserID:  userId,
			})

			if err != nil {
				log.Println("Error response:", err)
			}

			return res, err
		},
	}
}

func DisableCommentsField(usecases postinterface.PostUsecases) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.Boolean,
		Description: "Disable new comments for post",
		Args: graphql.FieldConfigArgument{
			"postId": &graphql.ArgumentConfig{Type: graphql.Int},
			"userId": &graphql.ArgumentConfig{Type: graphql.Int},
		},
		Resolve: func(p graphql.ResolveParams) (any, error) {
			postId, _ := p.Args["postId"].(int)
			userId, _ := p.Args["userId"].(int)
			res, err := usecases.DisableComments(p.Context, args.DisableCommentsArgs{PostID: postId, UserID: userId})

			if err != nil {
				log.Println("Error response:", err)
			}

			return res, err
		},
	}
}
