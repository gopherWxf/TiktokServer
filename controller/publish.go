package controller

import (
	"TiktokServer/dfst"
	"TiktokServer/middleware"
	"TiktokServer/opdb"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"time"
)

// Publish check token then save upload file to public directory
func PublishAction(c *gin.Context) {
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	userInfo, err := opdb.GetInfo(claims.Id, claims.Name)
	if err != nil {
		c.JSON(http.StatusOK, dfst.PublishActionResponse{Status: dfst.Status{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		}})
		return
	}
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, dfst.PublishActionResponse{Status: dfst.Status{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		}})
		return
	}
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%d_%s", userInfo.Id, time.Now().Unix(), filename)
	saveFile := filepath.Join("./public/", finalName)
	err = c.SaveUploadedFile(data, saveFile)
	if err != nil {
		c.JSON(http.StatusOK, dfst.PublishActionResponse{Status: dfst.Status{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		}})
		return
	}
	//存入数据库视频的相关信息
	err = opdb.InsertVideo(userInfo.Id, finalName)
	if err != nil {
		c.JSON(http.StatusOK, dfst.PublishActionResponse{Status: dfst.Status{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		}})
		return
	}
	c.JSON(http.StatusOK, dfst.PublishActionResponse{Status: dfst.Status{
		StatusCode: 0,
		StatusMsg:  filename + "上传成功",
	}})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	//从数据库中获取数据
	videos, err := opdb.FindAllVideos(claims.Id)
	if err != nil {
		c.JSON(http.StatusOK, dfst.PublishListResponse{
			Status: dfst.Status{
				StatusCode: -1,
				StatusMsg:  err.Error()},
			VideoList: []dfst.Video{},
		})
		return
	}
	user, err := opdb.GetInfo(claims.Id, claims.Name)
	if err != nil {
		c.JSON(http.StatusOK, dfst.PublishListResponse{
			Status: dfst.Status{
				StatusCode: -1,
				StatusMsg:  err.Error()},
			VideoList: []dfst.Video{},
		})
		return
	}

	Author := dfst.User{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      false,
	}
	result := make([]dfst.Video, len(videos))
	for i, v := range videos {
		result[i].Id = v.Id
		result[i].Author = Author
		result[i].PlayUrl = v.PlayUrl
		result[i].CoverUrl = v.CoverUrl
		result[i].FavoriteCount = v.FavoriteCount
		result[i].CommentCount = v.CommentCount
	}
	c.JSON(http.StatusOK, dfst.PublishListResponse{
		Status: dfst.Status{
			StatusCode: 0,
			StatusMsg:  "返回成功",
		},
		VideoList: result,
	})
}
