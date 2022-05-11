package opdb

import (
	"TiktokServer/dfst"
	"errors"
)

func Register(rr dfst.RegisterRequest) (_ UserInfo, _ error) {
	//查看该用户是否在数据库中
	if UserIsExists(rr.Username) {
		return UserInfo{}, errors.New("用户已经存在,请直接登陆")
	}
	userInfo := UserInfo{
		Name:          rr.Username,
		Password:      rr.Password,
		FollowCount:   0,
		FollowerCount: 0,
	}
	//将用户信息插入数据库中
	err := userInfo.Insert()
	if err != nil {
		return UserInfo{}, err
	}
	userInfo.GetInfo()
	return userInfo, nil
}

//查看该用户是否在数据库中
func UserIsExists(name string) bool {
	result := false
	// 指定库
	var user UserInfo
	dbResult := DB.Where("name = ?", name).Find(&user)
	if dbResult.Error != nil {
		result = false
	} else {
		result = true
	}
	return result
}

//将用户信息插入数据库中
func (user *UserInfo) Insert() error {
	return DB.Model(&UserInfo{}).Create(&user).Error
}
func (user *UserInfo) GetInfo() {
	DB.Find(user, "name=? && password=?", user.Name, user.Password)
}
