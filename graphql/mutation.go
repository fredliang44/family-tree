package graphql

import (
	"github.com/graphql-go/graphql"
)

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootMutation",
	Description: "Root Mutation",
	Fields: graphql.Fields{
		"addGroup":    &addGroup,
		"addCompany":  &addCompany,
		"updateUser":  &updateUser,
		"updateGroup": &updateGroup,
	},
})

var addCompany = graphql.Field{
	Name:        "Company",
	Description: "Add Company",
	Type:        graphql.NewNonNull(groupType),
	Args: graphql.FieldConfigArgument{
		"groupName": &graphql.ArgumentConfig{Type: graphql.String},
		"startYear": &graphql.ArgumentConfig{Type: graphql.Int},
		"endYear":   &graphql.ArgumentConfig{Type: graphql.Int},
		"memberIDs": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.Int),
		},
	},
	Resolve: AddCompany,
}

var addGroup = graphql.Field{
	Name:        "Group",
	Description: "Add group",
	Type:        graphql.NewNonNull(groupType),
	Args: graphql.FieldConfigArgument{
		"groupName":   &graphql.ArgumentConfig{Type: graphql.String},
		"startYear":   &graphql.ArgumentConfig{Type: graphql.Int},
		"endYear":     &graphql.ArgumentConfig{Type: graphql.Int},
		"fromGroupID": &graphql.ArgumentConfig{Type: graphql.Int},
		"leaderIDs": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.Int),
		},
		"toGroupIDs": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.Int),
		},
		"memberIDs": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.Int),
		},
	},
	Resolve: AddGroup,
}

var updateGroup = graphql.Field{
	Name:        "Group",
	Description: "Update group",
	Type:        graphql.NewNonNull(groupType),
	Args: graphql.FieldConfigArgument{
		"id":          &graphql.ArgumentConfig{Type: graphql.Int},
		"groupName":   &graphql.ArgumentConfig{Type: graphql.String},
		"startYear":   &graphql.ArgumentConfig{Type: graphql.Int},
		"endYear":     &graphql.ArgumentConfig{Type: graphql.Int},
		"fromGroupID": &graphql.ArgumentConfig{Type: graphql.Int},
		"toGroupIDs": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.Int),
		},
		"leaderIDs": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.Int),
		},
		"memberIDs": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.Int),
		},
	},
	Resolve: UpdateGroup,
}

var updateUser = graphql.Field{
	Name:        "User",
	Description: "Mutate user",
	Type:        graphql.NewNonNull(userType),
	Args: graphql.FieldConfigArgument{
		"id":               &graphql.ArgumentConfig{Type: graphql.String},
		"password":         &graphql.ArgumentConfig{Type: graphql.String},
		"username":         &graphql.ArgumentConfig{Type: graphql.String},
		"realname":         &graphql.ArgumentConfig{Type: graphql.String},
		"email":            &graphql.ArgumentConfig{Type: graphql.String},
		"phone":            &graphql.ArgumentConfig{Type: graphql.String},
		"avatar":           &graphql.ArgumentConfig{Type: graphql.String},
		"wechat":           &graphql.ArgumentConfig{Type: graphql.String},
		"loaction":         &graphql.ArgumentConfig{Type: graphql.String},
		"inviteCode":       &graphql.ArgumentConfig{Type: graphql.String},
		"createdTime":      &graphql.ArgumentConfig{Type: graphql.String},
		"isGraduate":       &graphql.ArgumentConfig{Type: graphql.Boolean},
		"IsActivated":      &graphql.ArgumentConfig{Type: graphql.Boolean},
		"IsBasicCompleted": &graphql.ArgumentConfig{Type: graphql.Boolean},
		"IsAdmin":          &graphql.ArgumentConfig{Type: graphql.Boolean},
		"mentorIDs": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.Int),
		},
	},
	Resolve: UpdateUser,
}
