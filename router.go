package main

import (
	"TiktokServer/controller"
	"TiktokServer/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initRouter(r *gin.Engine) {
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")
	{
		//test
		apiRouter.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, "666")
		})

		/*
			basic apis
		*/
		//无需登录，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个
		apiRouter.GET("/feed/", controller.Feed)
		//新用户注册时提供用户名，密码，昵称即可，用户名需要保证唯一。创建成功后返回用户 id 和权限token
		apiRouter.POST("/user/register/", controller.Register)
		//通过用户名和密码进行登录，登录成功后返回用户 id 和权限 token
		apiRouter.POST("/user/login/", controller.Login)
	}
	apiRouter.Use(middleware.JWTAuth)
	{
		//获取登录用户的 id、昵称，如果实现社交部分的功能，还会返回关注数和粉丝数
		apiRouter.GET("/user/", controller.UserInfo)
		//登录用户选择视频上传
		apiRouter.POST("/publish/action/", controller.PublishAction)
		//登录用户的视频发布列表，直接列出用户所有投稿过的视频
		apiRouter.GET("/publish/list/", controller.PublishList)

		/*
			extra apis - I
		*/
		////登录用户对视频的点赞和取消点赞操作
		//apiRouter.POST("/favorite/action/", controller.FavoriteAction)
		////登录用户的所有点赞视频
		//apiRouter.GET("/favorite/list/", controller.FavoriteList)
		////登录用户对视频进行评论
		//apiRouter.POST("/comment/action/", controller.CommentAction)
		////查看视频的所有评论，按发布时间倒序
		//apiRouter.GET("/comment/list/", controller.CommentList)
		//
		///*
		//	extra apis - II
		//*/
		////登陆用户对其他用户进行关注或取消关注
		//apiRouter.POST("/relation/action/", controller.RelationAction)
		////登陆用户关注的所有用户列表
		//apiRouter.GET("/relation/follow/list/", controller.FollowList)
		////所有关注登录用户的粉丝列表
		//apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	}
}
