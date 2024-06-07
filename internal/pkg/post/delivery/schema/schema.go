package schema

import (
	objects "github.com/MaratKamalovPD/o3_test_task/internal/models/graphql_objects"
	postinterface "github.com/MaratKamalovPD/o3_test_task/internal/pkg/post"
	"github.com/graphql-go/graphql"
)

// NewSchema creates a new GraphQL schema with the given resolver
func NewPostSchema(usecases postinterface.PostUsecases) (graphql.Schema, error) {
	post := objects.NewPostObject()

	rootQuery := query(post, usecases)
	rootMutation := mutation(post, usecases)

	schemaConfig := graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	}

	return graphql.NewSchema(schemaConfig)
}

func query(postType *graphql.Object, usecases postinterface.PostUsecases) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"posts": PostsField(postType, usecases),
			"post":  PostField(postType, usecases),
		},
	})
}

func mutation(postType *graphql.Object, usecases postinterface.PostUsecases) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createPost":      CreatePostField(postType, usecases),
			"disableComments": DisableCommentsField(usecases),
		},
	})
}
