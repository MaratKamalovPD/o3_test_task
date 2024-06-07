package schema

import (
	"log"

	"github.com/MaratKamalovPD/o3_test_task/internal/models/args"
	commentinterface "github.com/MaratKamalovPD/o3_test_task/internal/pkg/comment"
	"github.com/graphql-go/graphql"
)

func commentsByPostField(commentType *graphql.Object, usecases commentinterface.CommentUsecases) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(commentType),
		Description: "Get comments by post id",
		Args: graphql.FieldConfigArgument{
			"postId": &graphql.ArgumentConfig{Type: graphql.Int},
			"limit":  &graphql.ArgumentConfig{Type: graphql.Int},
			"offset": &graphql.ArgumentConfig{Type: graphql.Int},
		},
		Resolve: func(p graphql.ResolveParams) (any, error) {
			postId, _ := p.Args["postId"].(int)
			limit, _ := p.Args["limit"].(int)
			offset, _ := p.Args["offset"].(int)
			res, err := usecases.GetCommentsByPost(p.Context, args.GetCommentsArgs{
				PostID: postId,
				Limit:  limit,
				Offset: offset,
			})

			if err != nil {
				log.Println("Error response:", err)
			}

			return res, err
		},
	}
}

// func CommentsByParentField(commentType *graphql.Object, usecases commentinterface.CommentUsecases) *graphql.Field {
// 	return &graphql.Field{
// 		Type:        graphql.NewList(commentType),
// 		Description: "Get comments by parent comment id",
// 		Args: graphql.FieldConfigArgument{
// 			"parentId": &graphql.ArgumentConfig{Type: graphql.Int},
// 			"limit":    &graphql.ArgumentConfig{Type: graphql.Int},
// 			"offset":   &graphql.ArgumentConfig{Type: graphql.Int},
// 		},
// 		Resolve: func(p graphql.ResolveParams) (any, error) {
// 			parentId, _ := p.Args["parentId"].(int)
// 			limit, _ := p.Args["limit"].(int)
// 			offset, _ := p.Args["offset"].(int)
// 			// making parentId nil if it's 0
// 			var pointerToParent *int
// 			if parentId != 0 {
// 				pointerToParent = &parentId
// 			}
// 			res, err := usecases.GetCommentsByParent(p.Context, args.GetCommentsArgs{
// 				ParentID: pointerToParent,
// 				Limit:    limit,
// 				Offset:   offset,
// 			})
// 			logIfNotNil(err)
// 			return res, err
// 		},
// 	}
// }

func createCommentField(commentType *graphql.Object, usecases commentinterface.CommentUsecases) *graphql.Field {
	return &graphql.Field{
		Type:        commentType,
		Description: "Create new comment",
		Args: graphql.FieldConfigArgument{
			"postId":          &graphql.ArgumentConfig{Type: graphql.Int},
			"parentCommentId": &graphql.ArgumentConfig{Type: graphql.Int},
			"userId":          &graphql.ArgumentConfig{Type: graphql.Int},
			"content":         &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(p graphql.ResolveParams) (any, error) {
			postId, _ := p.Args["postId"].(int)
			parentCommentId, _ := p.Args["parentCommentId"].(int)
			userId, _ := p.Args["userId"].(int)
			content, _ := p.Args["content"].(string)

			parentCommentIdUint := uint(parentCommentId)

			var pointerToParentComment *uint

			if parentCommentIdUint != 0 {
				pointerToParentComment = &parentCommentIdUint
			}

			res, err := usecases.CreateComment(p.Context, args.CreateCommentArgs{
				PostID:          postId,
				ParentCommentID: pointerToParentComment,
				UserID:          userId,
				Content:         content,
			})

			if err != nil {
				log.Println("Error response:", err)

				return nil, err
			}

			return res, err
		},
	}
}
