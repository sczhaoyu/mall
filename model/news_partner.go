package model

import (
	"errors"
	"fmt"
	"time"
)

type NewsPartner struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Image      string    `json:"image"`  //图片标识
	Url        string    `json:"url"`    //链接
	Sort       int       `json:"sort"`   //数字越小，越靠前
	Status     int       `json:"status"` //0、不可见；1、可见
	CreateId   int       `json:"createId"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

func (n *NewsPartner) Save() error {
	_, err := ShopDB.Insert(n)
	return err
}

func (n *NewsPartner) Delete() error {
	_, err := ShopDB.Where("id=?", n.Id).Delete(NewsPartner{})
	return err
}
func GetNewsPartnerById(id int64) (*NewsPartner, error) {
	var ret NewsPartner
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("NewsPartner Not Found Value: %v", id))
	}
	return &ret, nil
}

func (n *NewsPartner) Update() error {
	_, err := ShopDB.Where("id=?", n.Id).Update(n)
	return err
}
