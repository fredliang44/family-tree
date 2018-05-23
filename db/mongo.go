package db

import (
	t "family-tree/graphql/types"
	"family-tree/utils"
	ai "github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

// DBSession locate mongo db session
var DBSession = mongoClient()

func mongoClient() *mgo.Session {

	info := &mgo.DialInfo{
		Addrs:    []string{utils.AppConfig.Mongo.Host + ":" + utils.AppConfig.Mongo.Port},
		Database: utils.AppConfig.Mongo.DB,
		Username: utils.AppConfig.Mongo.Username,
		Password: utils.AppConfig.Mongo.Password,
	}

	session, err := mgo.DialWithInfo(info)
	ai.Connect(session.DB(utils.AppConfig.Mongo.DB).C("user"))
	ai.Connect(session.DB(utils.AppConfig.Mongo.DB).C("project"))
	ai.Connect(session.DB(utils.AppConfig.Mongo.DB).C("group"))
	ai.Connect(session.DB(utils.AppConfig.Mongo.DB).C("company"))

	if err != nil {
		log.Println("mongoClient", err)
	}

	return session
}

// FetchUserFromMongo is a func to FetchUserFromMongo
func FetchUserFromMongo(username string) (user t.User, err error) {
	var p = bson.M{}
	var res = t.User{}

	p["username"] = username
	err = DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(p).One(&res)
	if err != nil || res.Username == "" {
		log.Println("GetUser: ", err)
		return res, err
	}
	return res, nil
}
