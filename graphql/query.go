package graphql

import (
	"github.com/graphql-go/graphql"
)

var queryCompany = graphql.Field{
	Name:        "Company",
	Description: "Query company",
	Type:        graphql.NewList(userType),

	Args: graphql.FieldConfigArgument{
		"id":   &graphql.ArgumentConfig{Type: graphql.Int},
		"name": &graphql.ArgumentConfig{Type: graphql.String},
	},
	Resolve: GetCompany,
}

var queryUser = graphql.Field{
	Name:        "User",
	Description: "Query user",
	Type:        graphql.NewList(userType),

	Args: graphql.FieldConfigArgument{
		"id":       &graphql.ArgumentConfig{Type: graphql.Int},
		"username": &graphql.ArgumentConfig{Type: graphql.String},
		"phone":    &graphql.ArgumentConfig{Type: graphql.String},
		"email":    &graphql.ArgumentConfig{Type: graphql.String},
	},
	Resolve: GetUser,
}

var queryGroup = graphql.Field{
	Name:        "Group",
	Description: "Query Group",
	Type:        graphql.NewList(groupType),

	Args: graphql.FieldConfigArgument{
		"id":          &graphql.ArgumentConfig{Type: graphql.String},
		"groupName":   &graphql.ArgumentConfig{Type: graphql.String},
		"startYear":   &graphql.ArgumentConfig{Type: graphql.Int},
		"endYear":     &graphql.ArgumentConfig{Type: graphql.Int},
		"fromGroupID": &graphql.ArgumentConfig{Type: graphql.Int},
	},
	Resolve: GetGroup,
}

var queryProject = graphql.Field{
	Name:        "Project",
	Description: "Query Project",
	Type:        graphql.NewList(projectType),

	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{Type: graphql.String},
	},
	Resolve: GetProject,
}

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootQuery",
	Description: "Root Query",
	Fields: graphql.Fields{
		"user":    &queryUser,
		"group":   &queryGroup,
		"project": &queryProject,
		"company": &queryCompany,
	},
})
