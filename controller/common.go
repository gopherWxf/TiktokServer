package controller

/*
	BASE
*/
//feed
type Status struct {
	StatusCode int32  `json:"status_code"` //状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  //返回状态描述
}

type User struct {
	Id            int64  `json:"id"`             //用户id
	Name          string `json:"name"`           //用户名称
	FollowCount   int64  `json:"follow_count"`   //关注总数
	FollowerCount int64  `json:"follower_count"` //粉丝总数
	IsFollow      bool   `json:"is_follow"`      //true-已关注，false-未关注
}

type Video struct {
	Id            int64  `json:"id"`             //视频唯一标识
	Author        User   `json:"author"`         //User
	PlayUrl       string `json:"play_url"`       //视频播放地址
	CoverUrl      string `json:"cover_url"`      //视频封面地址
	FavoriteCount int64  `json:"favorite_count"` //视频的点赞总数
	CommentCount  int64  `json:"comment_count"`  //视频的评论总数
	IsFavorite    bool   `json:"is_favorite"`    //true-已点赞，false-未点赞
}
type FeedResponse struct {
	Status            //状态相关
	NextTime  int64   `json:"next_time"`  //可选参数，本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	VideoList []Video `json:"video_list"` //视频列表
}
type FeedRequest struct {
	LatestTime int64 `json:"latest_time"` //可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
}

//id and token
type IdAndToken struct {
	UserId int64  `json:"user_id"` //用户id
	Token  string `json:"token"`   //用户鉴权token
}

//register
type RegisterResponse struct {
	Status     //状态相关
	IdAndToken //id and token
}
type RegisterRequest struct {
	Username string `json:"username"` //注册用户名，最长32个字符
	Password string `json:"password"` //密码，最长32个字符
}

//login
type LoginResponse struct {
	Status     //状态相关
	IdAndToken //id and token
}
type LoginRequest struct {
	Username string `json:"username"` //登录用户名
	Password string `json:"password"` //登录密码
}

//user
type UserResponse struct {
	Status      //状态相关
	User   User `json:"user"` //User
}
type UserRequest struct {
	IdAndToken //id and token
}

//publish action
type PublishActionResponse struct {
	Status //状态相关
}
type PublishActionRequest struct {
	IdAndToken        //id and token
	Data       []byte `json:"data"`
}

//publish list
type PublishListResponse struct {
	Status            //状态相关
	VideoList []Video `json:"video_list"` //视频列表
}
type PublishListRequest struct {
	IdAndToken //id and token
}

/*
	EXTRA-I
*/

//favorite action
type FavoriteActionResponse struct {
	Status //状态相关
}
type FavoriteActionRequest struct {
	IdAndToken       //id and token
	VideoId    int64 `json:"video_id"`    //视频id
	ActionType int32 `json:"action_type"` //1-点赞，2-取消点赞
}

//favorite list
type FavoriteListResponse struct {
	Status            //状态相关
	VideoList []Video `json:"video_list"` //视频列表
}
type FavoriteListRequest struct {
	IdAndToken //id and token
}

//comment action
type CommentActionResponse struct {
	Status
}
type CommentActionRequest struct {
	IdAndToken
	VideoId     int64  `json:"video_id"`     //视频id
	ActionType  int32  `json:"action_type"`  //1-发布评论，2-删除评论
	CommentText string `json:"comment_text"` //用户填写的评论内容，在action_type=1的时候使用
	CommentId   int64  `json:"comment_id"`   //要删除的评论id，在action_type=2的时候使用
}

//comment list
type Comment struct {
	Id         int64  `json:"id"`          //评论id
	User       User   `json:"user"`        //User
	Content    string `json:"content"`     //评论内容
	CreateDate string `json:"create_date"` //评论发布日期，格式 mm-dd
}

type CommentListResponse struct {
	Status
	CommentList []Comment `json:"comment_list"` //评论列表
}
type CommentListRequest struct {
	IdAndToken
	VideoId int64 `json:"video_id"` //视频id
}

/*
	EXTRA-II
*/

//relation action request
type RelationActionResponse struct {
	Status
}
type RelationActionRequest struct {
	IdAndToken
	ToUserId   int64 `json:"to_user_id"`  //对方用户id
	ActionType int32 `json:"action_type"` //1-关注，2-取消关注
}

//relation follow list
type RelationFollowListResponse struct {
	Status
	UserList []User `json:"user_list"` //用户信息列表
}
type RelationFollowListRequest struct {
	IdAndToken
}

//relation follower list
type RelationFollowerListResponse struct {
	Status
	UserList []User `json:"user_list"` //用户信息列表
}
type RelationFollowerListRequest struct {
	IdAndToken
}
