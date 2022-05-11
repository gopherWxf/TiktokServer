package opdb

import (
	"time"
)

//用户表
type UserInfo struct {
	Id            int64  `gorm:"AUTO_INCREMENT"`   //用户id
	Name          string `gorm:"type:varchar(32)"` //用户名称
	Password      string `gorm:"type:varchar(32)"` //用户密码
	FollowCount   int64  //关注总数
	FollowerCount int64  //粉丝总数
}

//视频表
type Video struct {
	Id             int64     `gorm:"AUTO_INCREMENT"` //视频唯一标识
	FkViUserinfoId int64     //外键
	PlayUrl        string    //视频播放地址
	CoverUrl       string    //视频封面地址
	FavoriteCount  int64     //视频的点赞总数
	CommentCount   int64     //视频的评论总数
	CreatedAt      time.Time //视频上传日期
}

//点赞表
type Favorite struct {
	UserInfoID int64 //外键
	VideoId    int64 //外键
}

//评论表
type Comment struct {
	Id         int64     `gorm:"AUTO_INCREMENT"` //评论id
	UserInfoID int64     //外键
	VideoId    int64     //外键
	Content    string    //评论内容
	CreatedAt  time.Time //评论发布日期
}

//关注表
type Relation struct {
	UserInfoID   int64
	UserInfoToID int64
}
