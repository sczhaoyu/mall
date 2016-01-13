package model

import (
	"errors"
	"fmt"
	"time"
)

type MSellerIndexFloor struct {
	Id             int64     `json:"id"`
	SellerId       int       `json:"sellerId"` //商家ID
	Name           string    `json:"name"`     //楼层名称
	OrderNo        int       `json:"orderNo"`  //楼层排序号
	Status         int       `json:"status"`   //状态1显示0不显示
	Remark         string    `json:"remark"`   //楼层备注
	CreateUserId   int       `json:"createUserId"`
	CreateUserName string    `json:"createUserName"`
	CreateTime     time.Time `json:"createTime"`
	UpdateUserId   int       `json:"updateUserId"`
	UpdateUserName string    `json:"updateUserName"`
	UpdateTime     time.Time `json:"updateTime"` //
}

func (m *MSellerIndexFloor) Save() error {
	_, err := MDB.Insert(m)
	return err
}

func (m *MSellerIndexFloor) Delete() error {
	_, err := MDB.Where("id=?", m.Id).Delete(MSellerIndexFloor{})
	return err
}
func GetMSellerIndexFloorById(id int64) (*MSellerIndexFloor, error) {
	var ret MSellerIndexFloor
	b, err := MDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("MSellerIndexFloor Not Found Value: %v", id))
	}
	return &ret, nil
}

func (m *MSellerIndexFloor) Update() error {
	_, err := MDB.Where("id=?", m.Id).Update(m)
	return err
}
