package graphql

import (
	"family-tree/db"
	types "family-tree/graphql/types"
	"family-tree/utils"
	"fmt"

	"github.com/getsentry/raven-go"
	"github.com/graphql-go/graphql"
	ast "github.com/graphql-go/graphql/language/ast"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func loadParams(params graphql.ResolveParams, key string) interface{} {
	value, isOK := params.Args[key]
	if isOK {
		return value
	}
	return nil
}

func fetchUsersByIDs(IDs []uint64) (resp interface{}, err error) {
	var res []types.User
	var p = bson.M{"id": bson.M{"$in": IDs}}
	err = db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(p).All(&res)

	if err != nil {
		log.Fatal("fetchDetailByUsername: ", err)
	}
	return res, err
}

func getSelectedFields(params graphql.ResolveParams) ([]string, error) {
	fieldASTs := params.Info.FieldASTs
	if len(fieldASTs) == 0 {
		return nil, fmt.Errorf("getSelectedFields: ResolveParams has no fields")
	}
	return selectedFieldsFromSelections(params, fieldASTs[0].SelectionSet.Selections)
}

func selectedFieldsFromSelections(params graphql.ResolveParams, selections []ast.Selection) ([]string, error) {
	var selected []string
	for _, s := range selections {
		switch t := s.(type) {
		case *ast.Field:
			selected = append(selected, s.(*ast.Field).Name.Value)
		case *ast.FragmentSpread:
			n := s.(*ast.FragmentSpread).Name.Value
			frag, ok := params.Info.Fragments[n]
			if !ok {
				return nil, fmt.Errorf("getSelectedFields: no fragment found with name %v", n)
			}
			sel, err := selectedFieldsFromSelections(params, frag.GetSelectionSet().Selections)
			if err != nil {
				raven.CaptureError(err, nil)
				return nil, err
			}
			selected = append(selected, sel...)
		default:
			return nil, fmt.Errorf("getSelectedFields: found unexpected selection type %v", t)
		}
	}
	return selected, nil
}
