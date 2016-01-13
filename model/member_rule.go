package model

import (
	"errors"
	"fmt"
)

type MemberRule struct {
	Id            int64 `json:"id"`
	Register      int   `json:"register"`      //会员注册
	Login         int   `json:"login"`         //会员每天登陆
	OrderEvaluate int   `json:"orderEvaluate"` //订单商品评论
	OrderBuy      int   `json:"orderBuy"`      //购物积分
	OrderMax      int   `json:"orderMax"`      //购物积分上限
	State         int   `json:"state"`         //1、开始；2、关闭
}

func (m *MemberRule) Save() error {
	_, err := ShopDB.Insert(m)
	return err
}

func (m *MemberRule) Delete() error {
	_, err := ShopDB.Where("id=?", m.Id).Delete(MemberRule{})
	return err
}
func GetMemberRuleById(id int64) (*MemberRule, error) {
	var ret MemberRule
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("MemberRule Not Found Value: %v", id))
	}
	return &ret, nil
}

func (m *MemberRule) Update() error {
	_, err := ShopDB.Where("id=?", m.Id).Update(m)
	return err
}
