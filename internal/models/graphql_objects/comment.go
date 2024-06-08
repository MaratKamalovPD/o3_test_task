package graphqlobjects

import "github.com/graphql-go/graphql"

func NewCommentObject() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"postId": &graphql.Field{
				Type: graphql.Int,
			},
			"parentCommentId": &graphql.Field{
				Type: graphql.Int,
			},
			"userId": &graphql.Field{
				Type: graphql.Int,
			},
			"content": &graphql.Field{
				Type: graphql.String,
			},
			"createdAt": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	})
}
