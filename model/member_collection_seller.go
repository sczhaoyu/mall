package model

import (
	"errors"
	"fmt"
	"time"
)

type MemberCollectionSeller struct {
	Id         int64     `json:"id"`
	MemberId   int       `json:"memberId"`   //会员ID
	SellerId   int       `json:"sellerId"`   //商铺ID
	CreateTime time.Time `json:"createTime"` //收藏时间
	UpdateTime time.Time `json:"updateTime"` //删除时间
	State      int       `json:"state"`      //1、显示；2、删除
}

func (m *MemberCollectionSeller) Save() error {
	_, err := ShopDB.Insert(m)
	return err
}

func (m *MemberCollectionSeller) Delete() error {
	_, err := ShopDB.Where("id=?", m.Id).Delete(MemberCollectionSeller{})
	return err
}
func GetMemberCollectionSellerById(id int64) (*MemberCollectionSeller, error) {
	var ret MemberCollectionSeller
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("MemberCollectionSeller Not Found Value: %v", id))
	}
	return &ret, nil
}

func (m *MemberCollectionSeller) Update() error {
	_, err := ShopDB.Where("id=?", m.Id).Update(m)
	return err
}
