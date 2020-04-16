package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestInit(t *testing.T)  {
//	assert.Equal(t,Length,int(62),"长度不匹配")
//}

//func TestIdToString(t *testing.T) {
//	str:=IdToString(72)
//	assert.Equal(t,"1a",str,"转换62进制失败")
//}

func TestStringToId(t *testing.T) {
	id:=StringToId("17")
	assert.Equal(t,17,id,"转换62进制失败")
}