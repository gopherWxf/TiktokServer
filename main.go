package main

import (
	"TiktokServer/opdb"
	"github.com/gin-gonic/gin"
)

func main() {
	//从配置文件中读取数据库的配置信息并连接数据库
	if dbErr := opdb.InitMySqlConn(); dbErr != nil {
		panic(dbErr)
	}
	defer opdb.DB.Close()
	//初始化表结构
	opdb.InitModel()
	r := gin.Default()
	//后续看情况改成https,暂时不急
	initRouter(r)

	r.Run()
}
