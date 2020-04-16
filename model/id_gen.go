package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type IdGen struct {
	Id string `bson:"_id"`
	MaxValue int `bson:"MaxValue"`
}

var (
	IdGenCollection="IdGen"
)

func IncrMaxId(id string)(maxId int,err error)  {
	idGen:=&IdGen{}
	change := mgo.Change{
		Update: bson.M{"$inc": bson.M{"MaxValue": 1}},
		ReturnNew: true,
		Upsert: true,
	}
	_, err = GetDb().C(IdGenCollection).Find(bson.M{"_id": id}).Apply(change, &idGen)
	if err!=nil{
		return
	}
	maxId=idGen.MaxValue
	return
}
