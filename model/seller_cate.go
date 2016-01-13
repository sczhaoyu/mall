package model

import (
	"errors"
	"fmt"
	"time"
)

type SellerCate struct {
	Id         int64     `json:"id"`
	SellerId   int       `json:"sellerId"`   //商家id
	Pid        int       `json:"pid"`        //父类ID
	Name       string    `json:"name"`       //分类名称
	Path       string    `json:"path"`       //分类路径
	Sort       int       `json:"sort"`       //排序
	CreateId   int       `json:"createId"`   //创建人
	CreateTime time.Time `json:"createTime"` //创建时间
	Status     int       `json:"status"`     //1、显示；2、不显示；3、删除
}

func (s *SellerCate) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SellerCate) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SellerCate{})
	return err
}
func GetSellerCateById(id int64) (*SellerCate, error) {
	var ret SellerCate
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SellerCate Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SellerCate) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
