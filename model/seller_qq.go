package model

import (
	"errors"
	"fmt"
	"time"
)

type SellerQq struct {
	Id         int64     `json:"id"`
	SellerId   int       `json:"sellerId"`   //商家ID
	Name       string    `json:"name"`       //客服名称
	Qq         string    `json:"qq"`         //客服QQ
	CreateId   int       `json:"createId"`   //创建人
	CreateTime time.Time `json:"createTime"` //创建时间
	State      int       `json:"state"`      //1、显示；2、不显示；3、删除
}

func (s *SellerQq) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SellerQq) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SellerQq{})
	return err
}
func GetSellerQqById(id int64) (*SellerQq, error) {
	var ret SellerQq
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SellerQq Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SellerQq) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
