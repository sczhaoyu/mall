package model

import (
	"errors"
	"fmt"
	"time"
)

type SellerUser struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Password   string    `json:"password"`
	Code       string    `json:"code"`       //员工号
	RealName   string    `json:"realName"`   //真实姓名
	Phone      string    `json:"phone"`      //联系电话
	Job        string    `json:"job"`        //职务
	SellerId   int       `json:"sellerId"`   //所属商家
	RoleId     int       `json:"roleId"`     //商家角色ID
	State      int       `json:"state"`      //状态 1-正常 2-冻结 3-删除
	CreateId   int       `json:"createId"`   //创建人
	CreateTime time.Time `json:"createTime"` //创建时间
	UpdateId   int       `json:"updateId"`   //修改人
	UpdateTime time.Time `json:"updateTime"` //更新时间（最后一次登陆时间）
}

func (s *SellerUser) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SellerUser) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SellerUser{})
	return err
}
func GetSellerUserById(id int64) (*SellerUser, error) {
	var ret SellerUser
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SellerUser Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SellerUser) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
