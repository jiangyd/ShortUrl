package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)


var (
	Cfg *Config
)

type Config struct {
	Schema string `yaml:"schema"`
	Bind string `yaml:"bind"`
	Mongodb
}

type Mongodb struct {
	Address string `yaml:"address"`
	DB string `yaml:"db"`
}

func init()  {

	f,err:=os.Open("config/config.yaml")
	defer f.Close()
	if err!=nil{
		panic(err)
	}
	data,err:=ioutil.ReadAll(f)
	if err!=nil{
		panic(err)
	}
	yaml.Unmarshal(data,&Cfg)
	if err!=nil{
		panic(err)
	}

}