package model

import (
	"errors"
	"fmt"
	"time"
)

type MSellerIndexBanner struct {
	Id             int64     `json:"id"`
	SellerId       int       `json:"sellerId"`  //商家ID
	Title          string    `json:"title"`     //标题
	LinkUrl        string    `json:"linkUrl"`   //链接地址
	Image          string    `json:"image"`     //图片
	OrderNo        int       `json:"orderNo"`   //排序
	StartTime      time.Time `json:"startTime"` //展示开始时间
	EndTime        time.Time `json:"endTime"`   //展示结束时间
	Status         int       `json:"status"`    //状态1使用0不使用
	CreateUserId   int       `json:"createUserId"`
	CreateUserName string    `json:"createUserName"`
	CreateTime     time.Time `json:"createTime"`
	UpdateUserId   int       `json:"updateUserId"`
	UpdateUserName string    `json:"updateUserName"`
	UpdateTime     time.Time `json:"updateTime"`
}

func (m *MSellerIndexBanner) Save() error {
	_, err := MDB.Insert(m)
	return err
}

func (m *MSellerIndexBanner) Delete() error {
	_, err := MDB.Where("id=?", m.Id).Delete(MSellerIndexBanner{})
	return err
}
func GetMSellerIndexBannerById(id int64) (*MSellerIndexBanner, error) {
	var ret MSellerIndexBanner
	b, err := MDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("MSellerIndexBanner Not Found Value: %v", id))
	}
	return &ret, nil
}

func (m *MSellerIndexBanner) Update() error {
	_, err := MDB.Where("id=?", m.Id).Update(m)
	return err
}
