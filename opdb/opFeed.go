package opdb

import (
	"errors"
	"fmt"
)

func FindFeed(datatime string) ([]Video, error) {
	temp := make([]Video, 0)
	//result := DB.Where("created_at < ?", datatime).Order("created_at desc").Limit(30).Find(&temp)
	result := DB.Order("created_at desc").Limit(30).Find(&temp)

	fmt.Println(result.Error)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(temp) == 0 {
		return nil, errors.New("record not found")
	}
	return temp, nil
}
