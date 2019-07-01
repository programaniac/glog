package database

import (
	"github.com/globalsign/mgo"
)

var (
	mgoSession    *mgo.Session
	MongoHosts    = "mongodb://localhost:27017"
	MongoDatabase = "civ2"
)

func _get_session() (*mgo.Session, error) {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(MongoHosts)
		if err != nil {
			return nil, err // no, not really
		}
		mgoSession.SetPoolLimit(500)

	}
	return mgoSession.Clone(), nil
}

func GetDB() (*mgo.Database, error) {
	sess, err := _get_session()
	if err != nil {
		return nil, err
	}

	return sess.DB(MongoDatabase), nil
}

func GetCollection(collection string) (*mgo.Collection, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	return db.C(collection), nil
}
