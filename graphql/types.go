package graphql

import (
	t "family-tree/graphql/types"
	"github.com/graphql-go/graphql"
)

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "UserType",
	Description: "User Type",
	Fields: graphql.Fields{
		"id":               &graphql.Field{Type: graphql.Int},
		"username":         &graphql.Field{Type: graphql.String},
		"realname":         &graphql.Field{Type: graphql.String},
		"email":            &graphql.Field{Type: graphql.String},
		"phone":            &graphql.Field{Type: graphql.String},
		"avatar":           &graphql.Field{Type: graphql.String},
		"wechat":           &graphql.Field{Type: graphql.String},
		"location":         &graphql.Field{Type: graphql.String},
		"verifyCode":       &graphql.Field{Type: graphql.String},
		"createdTime":      &graphql.Field{Type: graphql.String},
		"isGraduate":       &graphql.Field{Type: graphql.Boolean},
		"IsActivated":      &graphql.Field{Type: graphql.Boolean},
		"IsValidated":      &graphql.Field{Type: graphql.Boolean},
		"IsBasicCompleted": &graphql.Field{Type: graphql.Boolean},
		"IsAdmin":          &graphql.Field{Type: graphql.Boolean},
		"mentorsID": &graphql.Field{
			Name: "mentorsID Type",
			Type: graphql.NewList(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(t.User).MentorIDs, nil
			},
		},
		"groupsID": &graphql.Field{
			Name: "mentorsID Type",
			Type: graphql.NewList(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(t.User).GroupIDs, nil
			},
		},
	},
})

var groupType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "GroupType",
	Description: "Group Type",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.Int},
		"groupName":   &graphql.Field{Type: graphql.String},
		"startYear":   &graphql.Field{Type: graphql.String},
		"endYear":     &graphql.Field{Type: graphql.String},
		"createdTime": &graphql.Field{Type: graphql.String},
		"memberIDs": &graphql.Field{
			Name: "memberIDs Type",
			Type: graphql.NewList(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(t.Group).MemberIDs, nil
			},
		},
	},
})

var projectType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ProjectType",
	Description: "Project Type",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.Int},
		"titile":      &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.String},
		"year":        &graphql.Field{Type: graphql.Int},
		"startedTime": &graphql.Field{Type: graphql.String},
		"endedTime":   &graphql.Field{Type: graphql.String},
		"adminID":     &graphql.Field{Type: graphql.ID},
		"logo":        &graphql.Field{Type: graphql.String},
		"createdTime": &graphql.Field{Type: graphql.String},
		"images": &graphql.Field{
			Name: "images Type",
			Type: graphql.NewList(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(t.Project).Images, nil
			},
		},
		"memberIDs": &graphql.Field{
			Name: "memberIDs Type",
			Type: graphql.NewList(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(t.Project).MemberIDs, nil
			},
		},
	},
})
