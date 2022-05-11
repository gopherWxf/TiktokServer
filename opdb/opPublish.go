package opdb

import "errors"

func InsertVideo(id int64, filename string) error {
	v := Video{
		FkViUserinfoId: id,
		//PlayUrl:        "http://" + Svr.Public + ":" + Svr.Port + "/static/" + filename,
		//CoverUrl:       "https://profile.csdnimg.cn/1/2/9/1_qq_42956653",
		PlayUrl:  "http://itaem.cn/test.mp4",
		CoverUrl: "http://112.74.73.147/1645008588830-c62e8fb48ebedc30c107f184c8dee6b8_1.jpg",
	}
	return v.Insert()
}

func (v *Video) Insert() error {
	return DB.Model(&Video{}).Create(&v).Error
}
func FindAllVideos(id int64) ([]Video, error) {
	temp := make([]Video, 0)
	result := DB.Where("fk_vi_userinfo_id=?", id).Find(&temp)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(temp) == 0 {
		return nil, errors.New("record not found")
	}
	return temp, nil
}
