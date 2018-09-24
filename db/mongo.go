package db

import (
	t "github.com/fredliang44/family-tree/graphql/types"
	"github.com/fredliang44/family-tree/utils"
	"log"

	"github.com/getsentry/raven-go"
	"github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
		raven.CaptureErrorAndWait(err, nil)
		log.Panic("mongoClient", err)
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

func checkAdminFromMongo(username string) (isAdmin bool) {
	var p = bson.M{}
	var res = t.User{}

	p["username"] = username
	err := DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(p).One(&res)

	if err != nil || res.Username == "" {
		log.Println("GetUser: ", err)
		return false
	}

	return res.IsAdmin
}

func fetchUserIDFromMongo(username string) (userID uint64, err error) {
	var p = bson.M{}
	var res = t.User{}

	p["username"] = username
	err = DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(p).One(&res)
	if err != nil || res.Username == "" {
		log.Println("GetUser: ", err)
		return res.ID, err
	}
	return res.ID, nil
}
