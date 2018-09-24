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

var addProject = graphql.Field{
	Name:        "Project",
	Description: "Add project",
	Type:        graphql.NewNonNull(projectType),
	Args: graphql.FieldConfigArgument{
		"title":       &graphql.ArgumentConfig{Type: graphql.String},
		"description": &graphql.ArgumentConfig{Type: graphql.String},
		"startedTime": &graphql.ArgumentConfig{Type: graphql.String},
		"endedTime":   &graphql.ArgumentConfig{Type: graphql.String},
		"adminID":     &graphql.ArgumentConfig{Type: graphql.Int},
		"github":      &graphql.ArgumentConfig{Type: graphql.String},
		"logo":        &graphql.ArgumentConfig{Type: graphql.String},
		"images": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.String),
		},
		"files": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.String),
		},
		"memberIDs": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.Int),
		},
		"roles": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.Int),
		},
	},
	Resolve: AddProject,
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
		"password":         &graphql.ArgumentConfig{Type: graphql.String},
		"username":         &graphql.ArgumentConfig{Type: graphql.String},
		"realname":         &graphql.ArgumentConfig{Type: graphql.String},
		"email":            &graphql.ArgumentConfig{Type: graphql.String},
		"phone":            &graphql.ArgumentConfig{Type: graphql.String},
		"avatar":           &graphql.ArgumentConfig{Type: graphql.String},
		"gender":           &graphql.ArgumentConfig{Type: graphql.Boolean},
		"wechat":           &graphql.ArgumentConfig{Type: graphql.String},
		"loaction":         &graphql.ArgumentConfig{Type: graphql.String},
		"verifyCode":       &graphql.ArgumentConfig{Type: graphql.String},
		"createdTime":      &graphql.ArgumentConfig{Type: graphql.String},
		"joinedYear":       &graphql.ArgumentConfig{Type: graphql.Int},
		"enrollmentYear":   &graphql.ArgumentConfig{Type: graphql.Int},
		"isGraduated":      &graphql.ArgumentConfig{Type: graphql.Boolean},
		"IsActivated":      &graphql.ArgumentConfig{Type: graphql.Boolean},
		"IsBasicCompleted": &graphql.ArgumentConfig{Type: graphql.Boolean},
		"IsAdmin":          &graphql.ArgumentConfig{Type: graphql.Boolean},
		"mentorIDs": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.Int),
		},
	},
	Resolve: UpdateUser,
}

var updateCompany = graphql.Field{
	Name:        "Company",
	Description: "Mutate company",
	Type:        graphql.NewNonNull(companyType),
	Args: graphql.FieldConfigArgument{
		"name":        &graphql.ArgumentConfig{Type: graphql.String},
		"description": &graphql.ArgumentConfig{Type: graphql.String},
		"logo":        &graphql.ArgumentConfig{Type: graphql.String},
		"adminIDs":    &graphql.ArgumentConfig{Type: graphql.Int},
		"images": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.String),
		},
		"memberIDs": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.Int),
		},
	},
	Resolve: UpdateCompany,
}

var updateProject = graphql.Field{
	Name:        "Project",
	Description: "Mutate project",
	Type:        graphql.NewNonNull(companyType),
	Args: graphql.FieldConfigArgument{
		"title":       &graphql.ArgumentConfig{Type: graphql.String},
		"description": &graphql.ArgumentConfig{Type: graphql.String},
		"startedTime": &graphql.ArgumentConfig{Type: graphql.String},
		"endedTime":   &graphql.ArgumentConfig{Type: graphql.String},
		"adminID":     &graphql.ArgumentConfig{Type: graphql.Int},
		"github":      &graphql.ArgumentConfig{Type: graphql.String},
		"logo":        &graphql.ArgumentConfig{Type: graphql.String},
		"images": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.String),
		},
		"files": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.String),
		},
		"memberIDs": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.Int),
		},
		"roles": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.String),
		},
	},
	Resolve: UpdateProject,
}
