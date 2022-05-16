package controller

import (
	"TiktokServer/dfst"
	"TiktokServer/opdb"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	lastTime := c.Query("latest_time")
	if lastTime == "" {
		lastTime = time.Now().Format("2006-01-02 15:04:05")
	} else {
		t, _ := strconv.ParseInt(lastTime, 10, 64)
		lastTime = time.Unix(t, 0).Format("2006-01-02 15:04:05")
	}
	fmt.Println(lastTime)
	videos, err := opdb.FindFeed(lastTime)
	fmt.Println(videos)
	if err != nil {
		c.JSON(http.StatusOK, dfst.FeedResponse{
			Status: dfst.Status{
				StatusCode: 0,
				StatusMsg:  err.Error(),
			},
			NextTime:  0,
			VideoList: []dfst.Video{},
		})
		return
	}
	result := make([]dfst.Video, len(videos))
	for i, v := range videos {
		result[i].Id = v.Id
		result[i].PlayUrl = v.PlayUrl
		result[i].CoverUrl = v.CoverUrl
		result[i].FavoriteCount = v.FavoriteCount
		result[i].CommentCount = v.CommentCount
		user, err := opdb.GetInfoForId(v.FkViUserinfoId)
		if err != nil {
			c.JSON(http.StatusOK, dfst.FeedResponse{
				Status: dfst.Status{
					StatusCode: 0,
					StatusMsg:  err.Error(),
				},
				NextTime:  0,
				VideoList: []dfst.Video{},
			})
			return
		}
		result[i].Author = dfst.User{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      false,
		}
	}
	var nextime int64
	if len(videos) != 0 {
		nextime = videos[len(videos)-1].CreatedAt.Unix()
	}
	c.JSON(http.StatusOK, dfst.FeedResponse{
		Status: dfst.Status{
			StatusCode: 0,
			StatusMsg:  "返回成功",
		},
		NextTime:  nextime,
		VideoList: result,
	})
}
