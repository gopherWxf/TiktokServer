package controller

import (
	"TiktokServer/dfst"
	"TiktokServer/opdb"
	"github.com/gin-gonic/gin"
	"net/http"
)

//新用户注册时提供用户名，密码，昵称即可，用户名需要保证唯一。创建成功后返回用户 id 和权限token
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if len(username) > 32 || len(password) > 32 || username == "" || password == "" {
		c.JSON(http.StatusOK, dfst.RegisterResponse{
			Status: dfst.Status{
				StatusCode: -1,
				StatusMsg:  "用户名或密码过长",
			},
		})
		return
	}
	//将用户信息插入数据库中
	rr := dfst.RegisterRequest{Username: username, Password: password}
	userInfo, err := opdb.Register(rr)
	//用户已存在或插入错误
	if err != nil {
		c.JSON(http.StatusOK, dfst.RegisterResponse{
			Status: dfst.Status{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	token, err := generateToken(c, username)
	if err != nil {
		c.JSON(http.StatusOK, dfst.RegisterResponse{
			Status: dfst.Status{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	//注册成功
	c.JSON(http.StatusOK, dfst.RegisterResponse{
		Status: dfst.Status{
			StatusCode: 0,
			StatusMsg:  "注册成功",
		}, IdAndToken: dfst.IdAndToken{
			UserId: userInfo.Id,
			Token:  token,
		},
	})
}

func Login(c *gin.Context) {

}

func UserInfo(c *gin.Context) {

}
