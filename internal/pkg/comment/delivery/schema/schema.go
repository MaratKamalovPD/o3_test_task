package schema

import (
	objects "github.com/MaratKamalovPD/o3_test_task/internal/models/graphql_objects"
	commentinterface "github.com/MaratKamalovPD/o3_test_task/internal/pkg/comment"
	"github.com/graphql-go/graphql"
)

// NewSchema creates a new GraphQL schema with the given resolver
func NewCommentSchema(usecases commentinterface.CommentUsecases) (graphql.Schema, error) {
	comment := objects.NewCommentObject()

	rootQuery := query(comment, usecases)
	rootMutation := mutation(comment, usecases)

	schemaConfig := graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	}

	return graphql.NewSchema(schemaConfig)
}

func query(commentType *graphql.Object, usecases commentinterface.CommentUsecases) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"commentsByPost": commentsByPostField(commentType, usecases),
		},
	})
}

func mutation(commentType *graphql.Object, usecases commentinterface.CommentUsecases) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createComment": createCommentField(commentType, usecases),
		},
	})
}
