package model

import (
	"errors"
	"fmt"
	"time"
)

type MemberCollectionProduct struct {
	Id         int64     `json:"id"`
	MemberId   int       `json:"memberId"`   //用户ID
	ProductId  int       `json:"productId"`  //商品ID
	CreateTime time.Time `json:"createTime"` //收藏时间
	UpdateTime time.Time `json:"updateTime"` //删除时间
	State      int       `json:"state"`      //1、正常显示；2、删除
}

func (m *MemberCollectionProduct) Save() error {
	_, err := ShopDB.Insert(m)
	return err
}

func (m *MemberCollectionProduct) Delete() error {
	_, err := ShopDB.Where("id=?", m.Id).Delete(MemberCollectionProduct{})
	return err
}
func GetMemberCollectionProductById(id int64) (*MemberCollectionProduct, error) {
	var ret MemberCollectionProduct
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("MemberCollectionProduct Not Found Value: %v", id))
	}
	return &ret, nil
}

func (m *MemberCollectionProduct) Update() error {
	_, err := ShopDB.Where("id=?", m.Id).Update(m)
	return err
}
