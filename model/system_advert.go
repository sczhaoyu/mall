package model

import (
	"errors"
	"fmt"
	"time"
)

type SystemAdvert struct {
	Id          int64     `json:"id"`
	AdBeginTime time.Time `json:"adBeginTime"` //开始时间
	AdEndTime   time.Time `json:"adEndTime"`   //结束时间
	AdClickNum  int       `json:"adClickNum"`  //点击数
	AdSlideNum  int       `json:"adSlideNum"`  //幻灯片序号
	AdStatus    int       `json:"adStatus"`    //状态 1-待审核 2-冻结 3-投放中
	AdText      string    `json:"adText"`      //广告文字
	AdTitle     string    `json:"adTitle"`     //广告名称
	AdUrl       string    `json:"adUrl"`       //广告链接
	AdPic       string    `json:"adPic"`       //广告图片
	AdApId      int       `json:"adApId"`      //广告位id
	AdUserId    int       `json:"adUserId"`    //用户id
	Createtime  time.Time `json:"createtime"`  //创建时间
}

func (s *SystemAdvert) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SystemAdvert) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SystemAdvert{})
	return err
}
func GetSystemAdvertById(id int64) (*SystemAdvert, error) {
	var ret SystemAdvert
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SystemAdvert Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SystemAdvert) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
