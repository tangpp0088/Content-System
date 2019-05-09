package controllers

import (
	_ "github.com/gomodule/redigo/redis"
	"github.com/gomodule/redigo/redis"
	"github.com/astaxie/beego"
)
func init()  {

	//链接函数
	conn,err := redis.Dial("tcp","127.0.0.1:6379")

	if err != nil {

		beego.Error("数据失败")

		return
	}

	//关闭
	defer conn.Close()


	//操作函数
	/*conn.Send("set","redis","test")
	conn.Flush()
	conn.Receive()*/

	//如果t1对应的是整形，t2对应的是字符串
	resp,err := conn.Do("mget","s1","s2","s3")

	//回复助手函数
	result,_ := redis.Values(resp,err)

	//把对应的函数扫描到相应变量里面
	var s1,s3 int
	var s2 string

	redis.Scan(result,&s1,&s2,&s3)

	beego.Info("redis完成数据传输",s1,s2,s3)
}
