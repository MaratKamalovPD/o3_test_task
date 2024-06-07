package graphqlobjects

import "github.com/graphql-go/graphql"

func NewPostObject() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"content": &graphql.Field{
				Type: graphql.String,
			},
			"userId": &graphql.Field{
				Type: graphql.Int,
			},
			"commentsDisabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	})
}
