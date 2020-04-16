package control

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"io/ioutil"
	"shortUrl/config"
	"shortUrl/model"
	"shortUrl/utils"
)

type Invalid struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Err string `json:"err"`
}



func Short(c *gin.Context)  {
	data,err:=ioutil.ReadAll(c.Request.Body)
	if err!=nil{
		return
	}
	t:=model.Url{}
	json.Unmarshal(data,&t)

	_,err=url.ParseRequestURI(t.SourceUrl)

	if err!=nil{
		var invalid Invalid
		invalid.Code=http.StatusBadRequest
		invalid.Err="invalid URI"
		invalid.Msg=err.Error()
		c.JSON(400,invalid)
		return
	}
	t.GenId()
	t.Save()

	t.ShortUrl=config.Cfg.Schema+":"+"//"+config.Cfg.Bind+"/"+utils.IdToString(t.Id)
	c.JSON(200,t)
}

func Redirct(c *gin.Context)  {
	urlPath:=c.Request.URL.Path
	urlPath=strings.Trim(urlPath,"/")
	id:=utils.StringToId(urlPath)
	u:=&model.Url{Id: id}
	err:=u.FindById()
	if err!=nil{
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "短链接生成工具",
		})
	}
	c.Redirect(http.StatusMovedPermanently,u.SourceUrl)
}

func Index(c *gin.Context)  {

}