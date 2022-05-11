package main

import (
	"TiktokServer/opdb"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
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
	r.Run(opdb.Svr.IP + ":" + opdb.Svr.Port)
}
