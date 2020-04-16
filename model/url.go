package model

import "gopkg.in/mgo.v2/bson"

type Url struct {
	Id int `bson:"_id"`
	SourceUrl string `bson:"SourceUrl"`
	ShortUrl string `bson:"-"`
}
var (
	UrlCollection="url"
)

func (url *Url)GenId()error  {
	sourceUrl:=url.SourceUrl
	err:=GetDb().C(UrlCollection).Find(bson.M{"SourceUrl":sourceUrl}).One(url)
	if err!=nil{
		url.Id,err=IncrMaxId("url")
		if err!=nil{
			return err
		}
	}
	return nil
}

func (url *Url)Insert() error{
	return GetDb().C(UrlCollection).Insert(url)
}

func (url *Url)Save() error  {
	return url.Upsert()
}

func (url *Url)FindById()error  {
	return GetDb().C(UrlCollection).FindId(url.Id).One(url)
}

func (url *Url)Find() error {
	return GetDb().C(UrlCollection).FindId(url.Id).One(url)
}

func (url *Url)Update() error {
	return GetDb().C(UrlCollection).Update(bson.M{"_id":url.Id},url)
}

func (url *Url)Upsert()error  {
	_,err:= GetDb().C(UrlCollection).Upsert(bson.M{"_id":url.Id},url)
	return err
}

func (url *Url)DeleteById() error {
	return GetDb().C(UrlCollection).Remove(bson.M{"_id":url.Id})
}