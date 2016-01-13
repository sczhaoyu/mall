package model

import (
	"errors"
	"fmt"
	"time"
)

type SellerType struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`       //类型名称
	Money      int       `json:"money"`      //保证金数额
	State      int       `json:"state"`      //状态1、显示；2、不显示
	Sort       int       `json:"sort"`       //排序
	CreateId   int       `json:"createId"`   //创建人
	CreateTime time.Time `json:"createTime"` //创建时间
	UpdateTime time.Time `json:"updateTime"` //更新时间
}

func (s *SellerType) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SellerType) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SellerType{})
	return err
}
func GetSellerTypeById(id int64) (*SellerType, error) {
	var ret SellerType
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SellerType Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SellerType) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
