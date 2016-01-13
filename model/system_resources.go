package model

import (
	"errors"
	"fmt"
	"time"
)

type SystemResources struct {
	Id         int64     `json:"id"`
	Pid        int       `json:"pid"`
	Url        string    `json:"url"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"createTime"`
	Type       int       `json:"type"`    //1、菜单；2、按钮
	Status     int       `json:"status"`  //1、未删除2、删除
	Scope      int       `json:"scope"`   //应用范围: 1-商家 2-平台
	ResId      string    `json:"resId"`   //资源id,有些菜单及按钮需要特定id以注册点击事件
	ResIcon    string    `json:"resIcon"` //资源图标,按钮资源可能需要指定图标
}

func (s *SystemResources) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SystemResources) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SystemResources{})
	return err
}
func GetSystemResourcesById(id int64) (*SystemResources, error) {
	var ret SystemResources
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SystemResources Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SystemResources) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
