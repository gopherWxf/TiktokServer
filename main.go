package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//后续看情况改成https,暂时不急
	//initRouter(r)

	r.Run()
}
