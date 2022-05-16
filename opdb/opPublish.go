package opdb

import "errors"

func InsertVideo(id int64, filename string) error {
	v := Video{
		FkViUserinfoId: id,
		PlayUrl:        "http://" + Svr.Public + ":" + Svr.Port + "/static/" + filename,
		CoverUrl:       "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
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
