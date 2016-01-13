package model

import (
	"errors"
	"fmt"
	"time"
)

type SellerRoles struct {
	Id         int64     `json:"id"`
	SellerId   int       `json:"sellerId"`  //商家ID
	RolesName  string    `json:"rolesName"` //角色名称
	RoleCode   string    `json:"roleCode"`  //角色code,唯一
	Content    string    `json:"content"`   //角色描述
	UserId     int       `json:"userId"`    //创建人
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	Status     int       `json:"status"` //1、未删除2、删除
}

func (s *SellerRoles) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SellerRoles) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SellerRoles{})
	return err
}
func GetSellerRolesById(id int64) (*SellerRoles, error) {
	var ret SellerRoles
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SellerRoles Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SellerRoles) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
