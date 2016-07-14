package data

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"golang.org/x/net/context"
)

var photoType *graphql.Object
var reactionType *graphql.Object
var userType *graphql.Object

var nodeDefinitions *relay.NodeDefinitions
var photoConnection *relay.GraphQLConnectionDefinitions

// Schema is our published GraphQL representation of objects and mutations
var Schema graphql.Schema

func init() {
	/**
	 * We get the node interface and field from the Relay library.
	 *
	 * The first method defines the way we resolve an ID to its object.
	 * The second defines the way we resolve an object to its GraphQL type.
	 */
	nodeDefinitions = relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
		IDFetcher: func(id string, info graphql.ResolveInfo, ct context.Context) (interface{}, error) {
			resolvedID := relay.FromGlobalID(id)
			if resolvedID.Type == "User" {
				return GetUser(resolvedID.ID), nil
			}
			if resolvedID.Type == "Photo" {
				return GetPhoto(resolvedID.ID), nil
			}
			return nil, nil
		},
		TypeResolve: func(p graphql.ResolveTypeParams) *graphql.Object {
			switch p.Value.(type) {
			case *User:
				return userType
			case *Photo:
				return photoType
			}
			return nil
		},
	})

	/**
	 * Define your own types here
	 */
	photoType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Photo",
		Description: "A shiny photo",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Photo", nil),
			"name": &graphql.Field{
				Description: "The name of the photo",
				Type:        graphql.String,
			},
			"reaction": &graphql.Field{
				Description: "The overall reaction",
				Type:        graphql.Float,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					source := p.Source.(*Photo)
					return GetPhotoReaction(source.ID), nil
				},
			},
		},
		Interfaces: []*graphql.Interface{
			nodeDefinitions.NodeInterface,
		},
	})

	photoConnection = relay.ConnectionDefinitions(relay.ConnectionConfig{
		Name:     "PhotoConnection",
		NodeType: photoType,
	})

	reactionType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Reaction",
		Description: "A user's reaction to a photo",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Reaction", nil),
			"reaction": &graphql.Field{
				Description: "The reaction value",
				Type:        graphql.Int,
			},
			"photo": &graphql.Field{
				Description: "The photo related to the reaction",
				Type:        photoType,
			},
			"user": &graphql.Field{
				Description: "The user who generated the reaction",
				Type:        userType,
			},
		},
	})

	userType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "User",
		Description: "A person who uses our app",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("User", nil),
			"name": &graphql.Field{
				Description: "The name of the user",
				Type:        graphql.String,
			},
			"photos": &graphql.Field{
				Type:        photoConnection.ConnectionType,
				Description: "A person's collection of photos",
				Args:        relay.ConnectionArgs,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					args := relay.NewConnectionArguments(p.Args)
					dataSlice := PhotosToInterfaceSlice(GetPhotos()...)
					return relay.ConnectionFromArray(dataSlice, args), nil
				},
			},
			"totalCount": &graphql.Field{
				Type:        graphql.Int,
				Description: "The count of the photos",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return len(GetPhotos()), nil
				},
			},
		},
		Interfaces: []*graphql.Interface{
			nodeDefinitions.NodeInterface,
		},
	})

	/**
	 * This is the type that will be the root of our query,
	 * and the entry point into our schema.
	 */
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"node": nodeDefinitions.NodeField,

			// Add you own root fields here
			"viewer": &graphql.Field{
				Type: userType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetViewer(), nil
				},
			},
		},
	})

	/**
	 * This is the type that will be the root of our mutations,
	 * and the entry point into performing writes in our schema.
	 */
	//	mutationType := graphql.NewObject(graphql.ObjectConfig{
	//		Name: "Mutation",
	//		Fields: graphql.Fields{
	//			// Add you own mutations here
	//		},
	//	})

	/**
	* Finally, we construct our schema (whose starting query type is the query
	* type we defined above) and export it.
	 */
	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
		Types: []graphql.Type{queryType, userType},
	})
	if err != nil {
		panic(err)
	}
}

// PhotosToInterfaceSlice gets an interface slice.
// See https://github.com/golang/go/wiki/InterfaceSlice
func PhotosToInterfaceSlice(photos ...*Photo) []interface{} {
	var interfaceSlice = make([]interface{}, len(photos))
	for i, d := range photos {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}
