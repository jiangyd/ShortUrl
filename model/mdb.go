package model

import "gopkg.in/mgo.v2"
import "shortUrl/config"


func GetDb() *mgo.Database {
	session,err:=mgo.Dial(config.Cfg.Mongodb.Address)
	if err!=nil{
		panic(err)
	}
	database:=session.DB(config.Cfg.Mongodb.DB)
	return database
}

