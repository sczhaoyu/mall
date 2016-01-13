package model

import (
	"errors"
	"fmt"
	"time"
)

type SystemAdvPos struct {
	Id         int64     `json:"id"`
	ApWidth    int       `json:"apWidth"`    //广告位宽度
	ApHeight   int       `json:"apHeight"`   //广告位高度
	ApAccUrl   string    `json:"apAccUrl"`   //默认链接
	ApContent  string    `json:"apContent"`  //广告位简介
	ApPrice    float64   `json:"apPrice"`    //广告位价格
	ApShowType int       `json:"apShowType"` //展示方式 1-可以发布多条广告并随机展示 2-只允许发布并展示一条广告
	ApStatus   int       `json:"apStatus"`   //是否启用 1-启用 0-关闭
	ApSysType  int       `json:"apSysType"`  //是否系统广告 1-是 0-否
	ApText     string    `json:"apText"`     //默认文字
	ApTitle    string    `json:"apTitle"`    //广告位名称
	ApType     string    `json:"apType"`     //广告位类别  img-图片 slide-幻灯 scroll-滚动 text-文字
	ApDefPic   string    `json:"apDefPic"`   //广告位默认图片
	Createtime time.Time `json:"createtime"` //创建时间
}

func (s *SystemAdvPos) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SystemAdvPos) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SystemAdvPos{})
	return err
}
func GetSystemAdvPosById(id int64) (*SystemAdvPos, error) {
	var ret SystemAdvPos
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SystemAdvPos Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SystemAdvPos) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
